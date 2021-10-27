package sendgrid

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

func(s service) newSendRequestPost() rest.Request {
	request := sendgrid.GetRequest(s.sendgridAPIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	return request
}
