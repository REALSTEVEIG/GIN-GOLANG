package services

import (
	"errors"
	"os"
	"project2/database"
	"project2/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func RegisterUser(IUser models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(IUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Firstname: IUser.Firstname, 
		Lastname: IUser.Lastname, 
		Email: IUser.Email, 
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func AuthenticateUser(email, password string) (models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, errors.New("incorrect Password")
	}

	return user, nil
}

func GenerateJWT(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(jwtSecret)
}