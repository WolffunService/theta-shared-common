package sendgrid

type Service interface {
	SendCodeForMail(email string, code int) error
	SendCustomMail(fromName, fromAddr, email, templateId string) error
}
type service struct {
	sendgridAPIKey      string
	sendgridAddressFrom string
}

// NewService creates a new sendgrid service.
func NewService(sendgridAPIKey, sendgridAddressFrom string) Service {
	return service{
		sendgridAPIKey,
		sendgridAddressFrom,
	}
}