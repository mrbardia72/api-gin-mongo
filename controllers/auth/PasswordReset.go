package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/api-gin-mongo/forms"
	"github.com/mrbardia72/api-gin-mongo/helpers"
	repo_user "github.com/mrbardia72/api-gin-mongo/repository/user"
	"github.com/mrbardia72/api-gin-mongo/services"
)

// PasswordReset handles user password request
func PasswordReset(c *gin.Context) {
	var data forms.PasswordResetCommand

	// Ensure they provide data based on the schema
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	// Ensures that the password provided matches the confirm
	if data.Password != data.Confirm {
		c.JSON(400, gin.H{"message": "Passwords do not match"})
		c.Abort()
		return
	}

	// Get token from link query sent to your email
	resetToken, _ := c.GetQuery("reset_token")

	// Decode the token
	userID, _ := services.DecodeNonAuthToken(resetToken)

	// Fetch the user
	result, err := repo_user.GetUserByEmail(userID)

	if err != nil {
		// Return response when we get an error while fetching user
		c.JSON(500, gin.H{"message": "Something wrong happened, try again later"})
		c.Abort()
		return
	}
	// Check if account exists
	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User accoun was not found"})
		c.Abort()
		return
	}

	// Hash the new password
	newHashedPassword := helpers.GeneratePasswordHash([]byte(data.Password))

	// Update user account
	_err := repo_user.UpdateUserPass(userID, newHashedPassword)

	if _err != nil {
		// Return response if we are not able to update user password
		c.JSON(500, gin.H{"message": "Somehting happened while updating your password try again"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "Password has been updated, log in"})
	c.Abort()
	return
}
