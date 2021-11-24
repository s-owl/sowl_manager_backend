package email

import (
	"fmt"
	"os"

	"github.com/s-owl/sowl_manager_backend/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(email string, verifyLink string) error {
	newMail := mail.NewV3Mail()
	NewEmail := mail.NewEmail("Sowl-Manager", "seungheon328@gmail.com")
	newMail.SetFrom(NewEmail)

	newMail.SetTemplateID("d-178a883c4e8c41d08d633acf81a5cf9c")

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("USER", email),
	}
	p.AddTos(tos...)
	p.SetDynamicTemplateData("Weblink", verifyLink)
	newMail.AddPersonalizations(p)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(newMail)
	response, err := sendgrid.API(request)
	if err != nil {
		err = utils.SendEmailError(err)
	} else {
		fmt.Println(response.Body)
	}

	return err
}