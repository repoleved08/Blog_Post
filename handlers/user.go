package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/repoleved08/blog/config"
	"github.com/repoleved08/blog/models"
	"github.com/repoleved08/blog/models/dto"
	"github.com/repoleved08/blog/validators"
	"gorm.io/gorm"
)

func sendErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, echo.Map{"error": message})
}


// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param registerDTO body dto.RegisterDTO true "Register DTO"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/register [post]
func Register(c echo.Context) error {
	var registerDTO dto.RegisterDTO
	if err := c.Bind(&registerDTO); err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, "invalid input")
	}

	// Validate the input
	if err := validators.ValidateStruct(&registerDTO); err != nil {
		formattedError := validators.FormatValidationError(err)
		return sendErrorResponse(c, http.StatusBadRequest, formattedError)
	}

	// Create the user struct
	user := models.User{
		Username: registerDTO.Username,
		Email:    registerDTO.Email,
		Role:     "user",
	}

	// Hash the user's password
	user.HashPassword(registerDTO.Password)

	// Create the user in the database
	if result := config.DB.Create(&user); result.Error != nil {
		// Log the actual error for debugging purposes
		log.Println(result.Error)
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to create user")
	}

	// Return success response
	return c.JSON(http.StatusCreated, user)
}

// @Summary Login a user
// @Description Login a user
// @Tags users
// @Accept json
// @Produce json
// @Param loginDTO body dto.LoginDTO true "Login DTO"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/login [post]

func Login(c echo.Context) error {
	// Define a LoginDTO to validate input
	var loginDTO dto.LoginDTO
	if err := c.Bind(&loginDTO); err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, "invalid input")
	}

	// Validate the input fields
	if err := validators.ValidateStruct(&loginDTO); err != nil {
		formattedError := validators.FormatValidationError(err)
		return sendErrorResponse(c, http.StatusBadRequest, formattedError)
	}

	// Fetch the user from the database
	var user models.User
	if result := config.DB.Where("username = ?", loginDTO.Username).First(&user); result.Error != nil {
		// Check if the user was found
		if result.Error == gorm.ErrRecordNotFound {
			return sendErrorResponse(c, http.StatusUnauthorized, "invalid credentials")
		}
		// Log any other DB error for debugging
		log.Println(result.Error)
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to login user")
	}

	// Check if the password is correct
	if user.CheckPassword(loginDTO.Password) != nil {
		return sendErrorResponse(c, http.StatusUnauthorized, "invalid credentials")
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("JWT signing error: ", err)
		return sendErrorResponse(c, http.StatusInternalServerError, "failed to generate token")
	}

	// Return the token
	return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}

