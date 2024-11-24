package psql

import (
	"log"
	"telesp/pkg/models"
)

func (m *TeleSp) CertainDelete(storage *models.PersonData) error {
	query := "DELETE FROM main " +
		"WHERE id = $1;"
	_, err := m.DB.Exec(query, storage.ID)
	if err != nil {
		return err
	}
	return nil
}

// TODO: create a CertainUpdate func which can update selected data row
func (m *TeleSp) CertainUpdate(storage *models.PersonData) error {
	//var query string
	params, id := CreateSubUpdateQuery(m, storage)

	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

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

	_, err = tx.Exec("UPDATE main SET first_name = $1 WHERE id = $2", params[0], id)

	_, err = tx.Exec("UPDATE main SET last_name = $1 WHERE id = $2", params[1], id)

	_, err = tx.Exec("UPDATE main SET middle_name = $1 WHERE id = $2", params[2], id)

	_, err = tx.Exec("UPDATE main SET street = $1 WHERE id = $2", params[3], id)

	_, err = tx.Exec("UPDATE main SET house = $1 WHERE id = $2", params[4], id)

	_, err = tx.Exec("UPDATE main SET building = $1 WHERE id = $2", params[5], id)

	_, err = tx.Exec("UPDATE main SET apartment = $1 WHERE id = $2", params[6], id)

	_, err = tx.Exec("UPDATE main SET phone_number = $1 WHERE id = $2", params[7], id)

	return nil
}
