package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/muhardiansyah15/go-rest-api/config"
	"github.com/muhardiansyah15/go-rest-api/helper"
	"github.com/muhardiansyah15/go-rest-api/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// Get JSON Input
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// Get User data based on inputted username
	var user models.User
	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username or password is not valid"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// Password Validation
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username or password is not valid"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// Create the JWT token
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Go Rest API",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Set the algorithm to used when doing sign in
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	//add the token in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Successful login!"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Register(w http.ResponseWriter, r *http.Request) {

	//Get inputted JSON
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//Hashing password using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert a new user to database
	if err := config.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Registration is successful"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete token from cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout successful"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
