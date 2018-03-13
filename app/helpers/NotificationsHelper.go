package helpers

import (
	"strconv"

	pusher "github.com/pusher/pusher-http-go"
	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// PusherClient will be all the notifications configuration object
var PusherClient = pusher.Client{
	AppId:   "489882",
	Key:     "a5c0b256fc67a9f088ff",
	Secret:  "c807e685002a1d150951",
	Cluster: "us2",
	Secure:  true,
}

// SendReservationNotification will send a push notification to all connected to the
// reservations channel and with the business
// @params reservation Reservation Model
// @returns none
func SendReservationNotification(reservation models.ReservationModel) {
	data := map[string]string{"message": "" + reservation.ClientInfo.FirstName + " " + reservation.ClientInfo.LastName + " just made a reservation."}
	PusherClient.Trigger("camaleon-reservations", "business-"+strconv.Itoa(reservation.Location.BusinessID), data)
}
