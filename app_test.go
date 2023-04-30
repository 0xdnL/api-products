package main

import (
	"database/sql"
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

	//                   "CREATE DATABASE IF NOT EXISTS inventory;"
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
