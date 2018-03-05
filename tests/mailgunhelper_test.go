package tests

import (
	"testing"

	"github.com/restcontmx/camaleon-reservations-api/app/helpers"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

func TestSendConfirmationEmail(t *testing.T) {
	reservationDummy := models.ReservationModel{
		UID: "asdf-1234-asdf-1234-123",
		ClientInfo: models.ClientInfoModel{
			Email:     "ramiro.gutierrez.alz@gmail.com",
			FirstName: "FirstName",
			LastName:  "LastName",
			Phone:     "123456789",
		},
		Table: models.TableModel{
			Name: "Table Test",
		},
	}
	_, err := helpers.SendConfirmationEmail(reservationDummy)

	if err != nil {
		t.Errorf("Sending Mail was incorrect; Error : %s", err)
	}
}
