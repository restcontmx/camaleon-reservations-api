package tests

import (
	"testing"

	"github.com/restcontmx/camaleon-reservations-api/app/helpers"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

func TestSendReservationNoticiation(t *testing.T) {
	reservationDummy := models.ReservationModel{
		UID: "asdf-1234-asdf-1234-123",
		ClientInfo: models.ClientInfoModel{
			Email:     "ramiro.gutierrez.alz@gmail.com",
			FirstName: "Ramiro",
			LastName:  "Gutierrez",
			Phone:     "123456789",
		},
		Table: models.TableModel{
			Name: "Table Test",
		},
		Location: models.LocationModel{BusinessID: 1},
	}
	helpers.SendReservationNotification(reservationDummy)
}
