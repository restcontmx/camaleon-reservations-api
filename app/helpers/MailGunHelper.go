package helpers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/restcontmx/camaleon-reservations-api/app/models"
)

// MailGunAPIKey is the mail gun authentication key
const MailGunAPIKey = "key-93786709ada1dde6b5857ed681388ac8"

// MailGunURL is the mail gun url
const MailGunURL = "api.mailgun.net/v3/email.balneariolaspalmas.co/messages"

// CamaleonEmail is the camaleon email for reservations
const CamaleonEmail = "reservations@camaleonpos.com"

// SendConfirmationEmail will send a confirmation reservation email
func SendConfirmationEmail(reservation models.ReservationModel) (bool, error) {
	var message = `
		<h2>Thanks for making a reservation on Camaleon Rervations</h2>
		<p>Reservation ID : ` + reservation.UID + `</p>` +
		`<p>Client Name : ` + reservation.ClientInfo.FirstName + ` ` + reservation.ClientInfo.LastName + `</p>` +
		`<p>Phone Number : ` + reservation.ClientInfo.Phone + `</p>` +
		`<p>Table : ` + reservation.Table.Name + `</p>` +
		`<p>Date : ` + reservation.Date.Format("01/02/2006 15:04:05") + `</p>`

	_, err := http.PostForm("https://api:"+MailGunAPIKey+"@"+MailGunURL, url.Values{
		"html":    {message},
		"to":      {reservation.ClientInfo.Email},
		"subject": {"Camaleon Reservations"},
		"from":    {CamaleonEmail},
	})

	if err != nil {
		return false, fmt.Errorf("%s", err)
	}

	return true, nil
}
