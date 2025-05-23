package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"telesp/pkg/models"
	"telesp/pkg/models/psql"
)

func CertainDeleteHandler(w http.ResponseWriter, r *http.Request) {
	const op = "CertainDeleteHandler"
	log.Println("CertainDeletHandler was called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
	}

	var data models.PersonData

	body, err := io.ReadAll(r.Body)
	decodErr := json.Unmarshal(body, &data)
	if decodErr != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := psql.OpenConn()
	if err != nil {
		log.Println("err of OpenConn"+op, err)
	}
	defer db.DB.Close()

	// add temporary solution
	err = db.CertainDelete(&data)
	if err != nil {
		log.Println("err of call CertainDelete into database: ", err)
	}
	log.Println("main row was seccessed deleted")
}

func CertainUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CertainUpdateHandler called")
	var data models.PersonData

	body, err := io.ReadAll(r.Body)
	decodErr := json.Unmarshal(body, &data)
	if decodErr != nil {
		http.Error(w, decodErr.Error(), http.StatusBadRequest)
	}
	fmt.Println(data)
	db, err := psql.OpenConn()
	if err != nil {
		log.Println("err of OpenConn"+"", err)
	}
	err = db.CertainUpdate(&data)
	if err != nil {
		log.Println("err of call CertainUpdate into database: ", err)
	}
}
