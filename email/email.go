package email

import (
	"context"
	"fmt"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
	"github.com/s-owl/sowl_manager_backend/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func ExtractVerifyLink(context context.Context, email string) (string, error) {
	authClient := firebaseapp.App().Auth

	actionCodeSettings := &auth.ActionCodeSettings{
		URL:             "http://localhost:8080/api/user/signup",
		HandleCodeInApp: false,
	}

	verifyLink, err := authClient.EmailVerificationLinkWithSettings(context, email, actionCodeSettings)
	if err != nil {
		err = fmt.Errorf("VerifyEmailLink: %w", err)
		utils.VerifyLinkError(err)
		return "", err
	}

	return verifyLink, nil
}

func SendEmail(email string, verifyLink string) error {
	template := mail.NewV3Mail()
	newEmail := mail.NewEmail("Sowl-Manager", "seungheon328@gmail.com")
	template.SetFrom(newEmail)

	template.SetTemplateID("d-178a883c4e8c41d08d633acf81a5cf9c")

	personalSettings := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("user", email),
	}
	personalSettings.AddTos(tos...)
	personalSettings.SetDynamicTemplateData("Weblink", verifyLink)

	template.AddPersonalizations(personalSettings)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(template)
	response, err := sendgrid.API(request)
	if err != nil {
		err = utils.SendEmailError(err)
	} else {
		fmt.Println(response.Body)
	}

	return err
}
