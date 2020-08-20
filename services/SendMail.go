
package services

// Import relevant dependecy
import (
    "fmt"
    "os"

    // Import Sendgrid Go library
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailObject defines email payload data
type EmailObject struct {
    To      string
    Body    string
    Subject string
}


// SendMail method to send email to user
func SendMail(subject string, body string, to string, html string, name string) bool {
    fmt.Println(os.Getenv("SENDGRID_API_KEY"))

    // The first parameter is how your email name will be
    from := mail.NewEmail("Just Open it", os.Getenv("SENDGRID_FROM_MAIL"))
    // The recipient
    _to := mail.NewEmail(name, to)
    // Body in plain text
    plainTextContent := body
    // Body in html form(You can style a html document convert to string and make it look like the morning brew newsletter)
    htmlContent := html
    // Create message
    message := mail.NewSingleEmail(from, subject, _to, plainTextContent, htmlContent)
    // initialize client
    client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
    _, err := client.Send(message)
    if err != nil {
        return false
    } else {
        return true
    }
}