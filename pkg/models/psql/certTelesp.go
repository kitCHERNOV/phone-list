package psql

import "telesp/pkg/models"

func (m *TeleSp) CertainDelete(storage *models.PersonData) error {
	query := "DELETE FROM main " +
		"WHERE id = $1;"
	_, err := m.DB.Exec(query, storage.ID)
	if err != nil {
		return err
	}
	return nil
}
