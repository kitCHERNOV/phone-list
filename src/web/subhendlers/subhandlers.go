package subhendlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"telesp/pkg/models"
)

func formJsonStr(data models.PersonData) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("jsonFile: ", jsonData)
	return jsonData
}

func SelectHandler(r *http.Request) []models.PersonData {
	data := models.PersonData{}
	respData := []models.PersonData{}

	//for i := 0; i < countOfParams; i++ {
	data.FirstName = r.FormValue("field" + strconv.Itoa(1))
	data.LastName = r.FormValue("field" + strconv.Itoa(2))
	data.MiddleName = r.FormValue("field" + strconv.Itoa(3))
	data.Street = r.FormValue("field" + strconv.Itoa(4))
	data.House = r.FormValue("field" + strconv.Itoa(5))
	data.Building = r.FormValue("field" + strconv.Itoa(6))
	data.Apartment = r.FormValue("field" + strconv.Itoa(7))
	data.PhoneNumber = r.FormValue("field" + strconv.Itoa(8))

	jsonData := formJsonStr(data)

	resp, err := http.Post("http://127.0.0.1:8080/send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err.Error())
	}
	//fmt.Println("aboba was added")
	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read response error:", err)
	}
	fmt.Println("Response: ", string(body))

	// Unmarshal
	if err := json.Unmarshal(body, &respData); err != nil {
		fmt.Println("Unmarshal response error:", err)
	}
	//print new params
	//fmt.Println(data.FirstName)
	//fmt.Println(data.LastName)

	//for _, v := range respData {
	//	showDataArr = append(showDataArr, fmt.Sprintf("%v %v %v %v %v %v %v %v", v.FirstName, v.LastName, v.MiddleName, v.Street, v.House, v.Building, v.Apartment, v.PhoneNumber))
	//}

	return respData
}

func AddHandler(r *http.Request) {
	data := models.PersonData{}

	data.FirstName = r.FormValue("field" + strconv.Itoa(1))
	data.LastName = r.FormValue("field" + strconv.Itoa(2))
	data.MiddleName = r.FormValue("field" + strconv.Itoa(3))
	data.Street = r.FormValue("field" + strconv.Itoa(4))
	data.House = r.FormValue("field" + strconv.Itoa(5))
	data.Building = r.FormValue("field" + strconv.Itoa(6))
	data.Apartment = r.FormValue("field" + strconv.Itoa(7))
	data.PhoneNumber = r.FormValue("field" + strconv.Itoa(8))

	jsonData := formJsonStr(data)

	_, err := http.Post("http://127.0.0.1:8080/add", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err.Error())
	}
	//fmt.Println("aboba was added")
	//defer resp.Body.Close()
}

func UpdateHandler(r *http.Request) {
	data := models.PersonData{}

	data.FirstName = r.FormValue("field" + strconv.Itoa(1))
	data.LastName = r.FormValue("field" + strconv.Itoa(2))
	data.MiddleName = r.FormValue("field" + strconv.Itoa(3))
	data.Street = r.FormValue("field" + strconv.Itoa(4))
	data.House = r.FormValue("field" + strconv.Itoa(5))
	data.Building = r.FormValue("field" + strconv.Itoa(6))
	data.Apartment = r.FormValue("field" + strconv.Itoa(7))
	data.PhoneNumber = r.FormValue("field" + strconv.Itoa(8))

	jsonData := formJsonStr(data)

	_, err := http.Post("http://127.0.0.1:8080/update", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Println(err.Error())
	}
}

func DeleteHandler(r *http.Request) {
	data := models.PersonData{}
	data.PhoneNumber = r.FormValue("field" + strconv.Itoa(8))

	jsonData := formJsonStr(data)

	_, err := http.Post("http://127.0.0.1:8080/delete", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Println(err.Error())
	}
}

// Part of work with certain elements
func SubCertainDeleteHandler(r *http.Request) {
	valueId, err := strconv.Atoi(r.FormValue("certain-delete"))
	if err != nil {
		log.Println(err.Error())
	}

	jsonData := formJsonStr(models.PersonData{ID: valueId})

	_, err = http.Post("http://127.0.0.1:8080/certain-delete", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err.Error())
	}
}

func SubCertainUpdateHandler(r *http.Request) {
	data := models.PersonData{}
	data.ID, _ = strconv.Atoi(r.FormValue("certain-update"))
	data.FirstName = r.FormValue("FirstName")
	data.LastName = r.FormValue("LastName")
	data.MiddleName = r.FormValue("MiddleName")
	data.Street = r.FormValue("Street")
	data.House = r.FormValue("House")
	data.Building = r.FormValue("Building")
	data.Apartment = r.FormValue("Apartment")
	data.PhoneNumber = r.FormValue("PhoneNumber")

	fmt.Println(data.ID)
	fmt.Println(data.FirstName)
	fmt.Println(data.LastName)
	fmt.Println(data.MiddleName)
	fmt.Println(data.Street)
	fmt.Println(data.House)
	fmt.Println(data.Building)
	fmt.Println(data.Apartment)
	fmt.Println(data.PhoneNumber)
}
