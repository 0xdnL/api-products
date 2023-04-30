package main

import (
	_ "github.com/go-sql-driver/mysql" // use package for side-effect
	log "github.com/sirupsen/logrus"
)

/*
type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info(r.RequestURI)
	fmt.Fprintf(w, "hello world")
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Info("[getAllProducts] ", r.RequestURI)
	json.NewEncoder(w).Encode(Products)
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	log.Info("[getProductById] ", r.RequestURI)

	vars := mux.Vars(r)
	key := vars["id"]

	id, err := strconv.Atoi(key)
	if err != nil {
		panic(err)
	}

	for _, product := range Products {
		if product.Id == id {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleReq() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/product/{id}", getProductById)
	router.HandleFunc("/products", getAllProducts)
	router.HandleFunc("/", handler)
	http.ListenAndServe(ServerAddr, router)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type Data struct {
	id   int
	name string
}
*/

func main() {
	/*
		Products = []Product{
			{Id: 1, Name: "chair", Quantity: 100, Price: 100.00},
			{Id: 2, Name: "table", Quantity: 150, Price: 185.00},
			{Id: 3, Name: "lamp", Quantity: 70, Price: 89.90},
		}

		connectionString := fmt.Sprintf("%v:%v@tcp(%v)/%v", DbUser, DbPass, DbHost, DbName)
		db, err := sql.Open("mysql", connectionString)

		checkErr(err)
		defer db.Close()

		results, err := db.Exec("insert into data values (4, 'jkl')")
		checkErr(err)

		lastInsertedId, err := results.LastInsertId()
		checkErr(err)
		log.Info("lastInsertedId: ", lastInsertedId)

		rowsAffected, err := results.RowsAffected()
		checkErr(err)
		log.Info("rowsAffected: ", rowsAffected)

		rows, err := db.Query("select * from data")
		checkErr(err)

		for rows.Next() {
			var data Data
			err := rows.Scan(&data.id, &data.name)
			checkErr(err)
			fmt.Println(data)
		}

		handleReq()
	*/

	app := App{}

	log.Info("init..")
	app.Init()

	log.Info("run on ", ServerAddr, "..")
	app.Run(ServerAddr)
}
