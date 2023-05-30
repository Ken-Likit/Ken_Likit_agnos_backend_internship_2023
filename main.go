package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"unicode"
)

// Struct for Request 
type inputRequestBody struct {
	InitPassword string `json:"init_password"`
}

// Program entry
func main() {
    router := gin.Default()
	router.POST("/passwordValidation", postPasswordValidation)
    router.Run("localhost:8081")
}

// Validates password.
// Returns one of the following:
// OK: Password Strength Evaluated
// BadRequest
func postPasswordValidation(c *gin.Context) {
    var pwRequest inputRequestBody
	
    if err := c.BindJSON(&pwRequest); err != nil {
        c.IndentedJSON(http.StatusBadRequest , gin.H{"Message": "Unable to Process"})
		return
    }
	
    result := strongPassword(pwRequest.InitPassword)
    c.IndentedJSON(http.StatusOK, gin.H{"num_of_steps": result})
}

// Evaluates the input password based on provided criteria.
// Returns the number of recommended changes. 
func strongPassword(password string) int {
	lengthChange := 0
	hasLower := 1
	hasUpper := 1
	hasDigit := 1
	characterChange := 0

	// Number of adds or removes character to get ideal character length
	pwLength := len(password)
	if pwLength >= 20 {
		lengthChange = pwLength - 19
	}
	if pwLength < 6 {
		lengthChange = 6 - pwLength
	}

	// Checks to see if password contains at least 1 digit, lower, and uppercase character.
    // Number of characters to add or change.
	for _, x := range password {
		if unicode.IsLower(x) {
			hasLower = 0
		}
		if unicode.IsUpper(x) {
			hasUpper = 0
		}
		if unicode.IsDigit(x) {
			hasDigit = 0
		}
	}

	// Checks for 3 repreating character in a row.
    // Number of characters to change to avoid three in a row. 
    // Treats lowercase and uppercase characters as different characters. 
	count := 1
	if pwLength > 1 {
		for i := 1; i < pwLength; i++ {
			if password[i-1] == password[i] {
				count++
				if count == 3 {
					characterChange++
					count = 0
				}
			} else {
				count = 1
			}
		}
	}

	missingType := hasLower + hasUpper + hasDigit
	formatChange := lengthChange + characterChange
	totalChange := 0

	if missingType > formatChange {
        totalChange = missingType
    } else {
        totalChange = formatChange
    }

	return totalChange
}

