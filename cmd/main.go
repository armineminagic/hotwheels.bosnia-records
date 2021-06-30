package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type dbHandler struct {
	db *sql.DB
}

type HwModel struct {
	Name         string `json:"name"`
	Set          string `json:"set"`
	Year         string `json:"year"`
	Manufacturer string `json:"manufacturer"`
	ModelNumber  string `json:"modelnumber"`
}

func main() {
	db := dbConn()
	defer db.db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/getmodels", middleware(db.GetAllmodels))
	router.HandleFunc("/addmodel", middleware(db.AddModel))
	router.HandleFunc("/removemodel/{name}/{modelnum}", middleware((db.RemoveModel)))
	log.Println("Server started...")
	http.ListenAndServe(":8070", router)
}

func dbConn() *dbHandler {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "hwrecords"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Println(err)
	}
	var dataBase dbHandler
	dataBase.db = db
	return &dataBase
}

func middleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if req.Method == "OPTIONS" {
			return
		}
		log.Println(req.Method)
		log.Println(req.RequestURI)
		handler.ServeHTTP(w, req)
	})
}

func (d *dbHandler) GetAllmodels(w http.ResponseWriter, req *http.Request) {
	var hwModel HwModel
	var hwModels []HwModel

	rows, err := d.db.Query("SELECT name, `set`, year, manufacturer, model_number FROM hwmodels")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		err := rows.Scan(&hwModel.Name, &hwModel.Set, &hwModel.Year, &hwModel.Manufacturer, &hwModel.ModelNumber)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hwModels = append(hwModels, hwModel)
	}
	response, _ := json.MarshalIndent(hwModels, "", " ")
	fmt.Fprintf(w, string(response))
}

func (d *dbHandler) AddModel(w http.ResponseWriter, req *http.Request) {
	var newHWModel HwModel
	inputForm, _ := ioutil.ReadAll(req.Body)

	defer req.Body.Close()
	err := json.Unmarshal(inputForm, &newHWModel)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q, err := d.db.Prepare("INSERT INTO hwmodels(name, `set`, year, manufacturer, model_number) VALUES(?,?,?,?,?)")
	y, _ := strconv.Atoi(newHWModel.Year)
	mn, _ := strconv.Atoi(newHWModel.ModelNumber)
	_, err = q.Exec(newHWModel.Name, newHWModel.Set, y, newHWModel.Manufacturer, mn)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer q.Close()
}

func (d *dbHandler) RemoveModel(w http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)

	_, err := d.db.Query("DELETE FROM hwmodels WHERE name=? and model_number=?", args["name"], args["modelnum"])
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
