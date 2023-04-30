package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"
)

var a App

const testDbName = "test"

func TestMain(m *testing.M) {
	createDatabase()
	err := a.Init(DbUser, DbPass, DbHost, testDbName)

	if err != nil {
		log.Fatal("Error occured while app.init()")
	}
	createTable()
	m.Run()
	// cleanup()
}

func cleanup() {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v)/", DbUser, DbPass, DbHost)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf("drop database %v", testDbName)
	_, err = db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

func createDatabase() {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v)/", DbUser, DbPass, DbHost)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", testDbName)
	_, err = db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS products(
		id INT NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		quantity int,
		price float(10,7),
		PRIMARY KEY(id)
		);`

	_, err := a.Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.Db.Exec("DELETE FROM products")
	a.Db.Exec("ALTER TABLE products AUTO_INCREMENT=1")
}

func addProduct(name string, quanity int, price float64) {
	query := fmt.Sprintf("INSERT INTO products (name, quantity, price) VALUES('%v', %v, %v)", name, quanity, price)
	_, err := a.Db.Exec(query)
	if err != nil {
		log.Println(err)
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	createTable()
	addProduct("keyboard", 100, 300)
	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received status: %v", expectedStatusCode, actualStatusCode)
	}
}

func TestCreateProductById(t *testing.T) {

	clearTable()

	var product = []byte(`{"name": "glass", "quantity": 1, "price": 100}`)

	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product)) //@TODO: find out how does bytes.NewBuffer implement io.Reader ?
	req.Header.Set("Content-Type", "application/json")
	response := sendRequest(req)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	// log.Println(m)
	// log.Printf("%T", m["Quantity"])
	if m["Name"] != "glass" {
		t.Errorf("Expected Name: %v, Got: %v", "glass", m["Name"])
	}

	if m["Quantity"] != 1.0 {
		t.Errorf("Expected Quantity: %v, Got: %v", 1.0, m["Quantity"])
	}
}

func TestDeleteProduct(t *testing.T) {
	clearTable()
	addProduct("phone", 10, 20)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/product/1", nil)
	response = sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/product/1", nil)
	response = sendRequest(req)
	checkStatusCode(t, http.StatusNotFound, response.Code)
}

func TestUpdateProductName(t *testing.T) {
	clearTable()
	addProduct("phone", 10, 20)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	var oldValue map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &oldValue)

	var product = []byte(`{"name": "banana", "quantity": 10, "price": 20}`)

	req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(product))
	req.Header.Set("Content-Type", "application/json")
	response = sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	var newValue map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &newValue)

	req, _ = http.NewRequest("GET", "/product/1", nil)
	response = sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	log.Println(oldValue)
	log.Println(newValue)

	if oldValue["ID"] != newValue["ID"] {
		t.Errorf("Expected ID: %v, Got: %v", oldValue["ID"], newValue["ID"])
	}

	if oldValue["Name"] == newValue["Name"] { // if value remains same, we fail
		t.Errorf("Expected Name: %v, Got: %v", oldValue["Name"], newValue["Name"])
	}

	if oldValue["Quantity"] != newValue["Quantity"] {
		t.Errorf("Expected quantity: %v, Got: %v", oldValue["Quantity"], newValue["Quantity"])
	}

	if oldValue["Price"] != newValue["Price"] {
		t.Errorf("Expected price: %v, Got: %v", oldValue["Price"], newValue["Price"])
	}
}

// @TODO: check for produkt that does not exist in db
// @TODO: test get all products api call
// @TODO: test posting wrong datatypes
