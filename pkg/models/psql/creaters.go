package psql

import (
	"fmt"
	"log"
	"telesp/pkg/models"
)

func AllRows() string {
	Query := "SELECT firstname.firstname_val, lastname.lastname_val, middlename.middlename_val, street.street_val," +
		"house.house_val, apartment.apartment_val, building.building_val, phonenumber.phonenumber_val " +
		"FROM main " +
		"JOIN firstname ON firstname.id = main.first_name " +
		"JOIN lastname ON lastname.id = main.last_name " +
		"JOIN middlename ON middlename.id = main.middle_name " +
		"JOIN street ON street.id = main.street " +
		"JOIN house ON house.id = main.house " +
		"JOIN building ON building.id = main.building " +
		"JOIN apartment ON apartment.id = main.apartment " +
		"JOIN phonenumber ON phonenumber.id = main.phone_number;"
	return Query
}

func CreateSqlQuery(storage *models.PersonData) string {
	Query := "SELECT firstname.firstname_val, lastname.lastname_val, middlename.middlename_val, street.street_val," +
		"house.house_val, apartment.apartment_val, building.building_val, phonenumber.phonenumber_val " +
		"FROM main " +
		"JOIN firstname ON firstname.id = main.first_name " +
		"JOIN lastname ON lastname.id = main.last_name " +
		"JOIN middlename ON middlename.id = main.middle_name " +
		"JOIN street ON street.id = main.street " +
		"JOIN house ON house.id = main.house " +
		"JOIN building ON building.id = main.building " +
		"JOIN apartment ON apartment.id = main.apartment " +
		"JOIN phonenumber ON phonenumber.id = main.phone_number " +
		"WHERE "
	if storage.FirstName != "" {
		Query += "firstname.firstname_val = " + fmt.Sprintf("'%s'", storage.FirstName) + " and "
	}
	if storage.LastName != "" {
		Query += "lastname.lastname_val = " + fmt.Sprintf("'%s'", storage.LastName) + " and "
	}
	if storage.MiddleName != "" {
		Query += "middlename.middlename_val = " + fmt.Sprintf("'%s'", storage.MiddleName) + " and "
	}
	if storage.Street != "" {
		Query += "street.street_val = " + fmt.Sprintf("'%s'", storage.Street) + " and "
	}
	if storage.House != "" {
		Query += "house.house_val = " + fmt.Sprintf("'%s'", storage.House) + " and "
	}
	if storage.Building != "" {
		Query += "building.building_val = " + fmt.Sprintf("'%s'", storage.Building) + " and "
	}
	if storage.Apartment != "" {
		Query += "apartment.apartment_val = " + fmt.Sprintf("'%s'", storage.Apartment) + " and "
	}
	if storage.PhoneNumber != "" {
		Query += "phonenumber.phonenumber_val = " + fmt.Sprintf("'%s'", storage.PhoneNumber)
	}
	if Query[len(Query)-4:] == "and " {
		Query = Query[:len(Query)-4]
	}
	Query += ";"
	//fmt.Println("Query:", Query)
	return Query
}

func CreateSubUpdateQuery(m *TeleSp, storage *models.PersonData) [8]int {
	var params = [8]int{}
	tx, err := m.DB.Begin()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(storage.Street, storage.PhoneNumber)
	defer func() {
		if err != nil {
			// IF find mistake,just rollback tx
			if rollbackerr := tx.Rollback(); rollbackerr != nil {
				log.Fatalf("Transaction have not opened: %v", rollbackerr)
			}
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				log.Fatalf("Transaction have not done: %v", commitErr)
			}
		}
	}()

	// Execute sql queries
	if storage.FirstName != "" {
		err := tx.QueryRow("SELECT insert_or_get_id_of_firstname($1);", storage.FirstName).Scan(&params[0])
		if err != nil {
			log.Fatal(err)
		}
	}
	if storage.LastName != "" {
		err := tx.QueryRow("SELECT insert_or_get_id_of_lastname($1);", storage.LastName).Scan(&params[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	if storage.MiddleName != "" {
		err := tx.QueryRow("SELECT insert_or_get_id_of_middlename($1);", storage.MiddleName).Scan(&params[2])
		if err != nil {
			tx.Rollback()
		}
	}
	if storage.Street != "" {
		err := tx.QueryRow("SELECT insert_or_get_id_of_street($1);", storage.Street).Scan(&params[3])
		if err != nil {
			log.Fatal(err)
		}
	}
	if storage.House != "" {
		err := tx.QueryRow("SELECT insert_or_get_id_of_house($1);", storage.House).Scan(&params[4])
		if err != nil {
			log.Fatal(err)
		}
	}
	if storage.Building != "" {
		err := tx.QueryRow("SELECT insert_or_get_id_of_building($1);", storage.Building).Scan(&params[5])
		if err != nil {
			log.Fatal(err)
		}
	}
	if storage.Apartment != "" {
		err := tx.QueryRow("SELECT insert_or_get_id_of_apartment($1);", storage.Apartment).Scan(&params[6])
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.QueryRow("SELECT insert_or_get_id_of_phonenumber($1);", storage.PhoneNumber).Scan(&params[7])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("params", params)
	return params
}
