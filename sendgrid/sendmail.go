package sendgrid

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func(s service) SendCodeForMail(email string, code int) error {
	var request = s.newSendRequestPost()
	body := s.mailBody(email, code)
	request.Body = body

	_, err := sendgrid.API(request)
	return err
}

func(s service) mailBody(email string, code int) []byte {
	m := mail.NewV3Mail()
	mailSubject := fmt.Sprintf("Wolffun Id [ %d ]", code)

	//from
	address := s.sendgridAddressFrom
	from := mail.NewEmail("Wolffun Support", address)
	m.SetFrom(from)

	//to
	address = email
	to := mail.NewEmail("", address)

	//personal
	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetDynamicTemplateData("subject", mailSubject)
	p.SetDynamicTemplateData("code", code)

	m.AddPersonalizations(p)

	m.SetTemplateID("d-61421615e40c43679d699e1b2150a657")
	m.Subject = mailSubject

	return mail.GetRequestBody(m)
}

//for mkt
func(s service) SendCustomMail(from, email, templateId string) error {
	var request = s.newSendRequestPost()
	body := s.customMailBody(from, email, templateId)
	request.Body = body

	_, err := sendgrid.API(request)
	return err
}

func(s service) customMailBody(fromAddr, email, templateId string) []byte {
	m := mail.NewV3Mail()

	//from
	if fromAddr == "" {
		fromAddr = s.sendgridAddressFrom
	}
	from := mail.NewEmail("", fromAddr)
	m.SetFrom(from)

	//to
	address := email
	to := mail.NewEmail("", address)

	//personal
	p := mail.NewPersonalization()
	p.AddTos(to)

	m.AddPersonalizations(p)

	m.SetTemplateID(templateId)

	return mail.GetRequestBody(m)
}
