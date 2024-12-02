package web

import (
	//"encoding/json"
	"fmt"
	//"telesp/pkg/models"
	//"telesp/pkg/models/psql"
	sh "telesp/src/web/subhendlers"

	"net/http"
	// "strconv"
)

func SubIndexUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Certain update function called")

	sh.SubCertainUpdateHandler(r)

	//tp, err := template.ParseFiles("./internal/html/main.html")
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//err = tp.Execute(w, showData)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//}
}
