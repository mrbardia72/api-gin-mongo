package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/api-gin-mongo/forms"
	repo_user "github.com/mrbardia72/api-gin-mongo/repository/user"
	"github.com/mrbardia72/api-gin-mongo/services"
)

// ResetLink handles resending email to user to reset link
func ResetLink(c *gin.Context) {
	// Defined schema for the request body
	var data forms.ResendCommand

	// Ensure the user provides all values from the request.body
	if (c.BindJSON(&data)) != nil {
		// Return 400 status if they don't provide the email
		c.JSON(400, gin.H{"message": "Provided all fields"})
		c.Abort()
		return
	}

	// Fetch the account from the database based on the email
	// provided
	result, err := repo_user.GetUserByEmail(data.Email)

	// Return 404 status if an account was not found
	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	// Return 500 status if something went wrong while fetching
	// account
	if err != nil {
		c.JSON(500, gin.H{"message": "Something wrong happened, try again later"})
		c.Abort()
		return
	}

	// Generate the token that will be used to reset the password
	resetToken, _ := services.GenerateNonAuthToken(result.Email)

	// The link to be clicked in order to perform a password reset
	link := "http://localhost:5000/api/v1/password-reset?reset_token=" + resetToken
	// Define the body of the email
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	// Initialize email sendout
	email := services.SendMail("Reset Password", body, result.Email, html, result.Name)

	// If email was sent, return 200 status code
	if email == true {
		c.JSON(200, gin.H{"messsage": "Check mail"})
		c.Abort()
		return
		// Return 500 status when something wrong happened
	} else {
		c.JSON(500, gin.H{"message": "An issue occured sending you an email"})
		c.Abort()
		return
	}
}
