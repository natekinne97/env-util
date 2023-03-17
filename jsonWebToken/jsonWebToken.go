package jsonWebTokenUtil

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/natekinne97/envUtil"
	"github.com/natekinne97/envUtil/types"
	"golang.org/x/crypto/bcrypt"
)

var loader = types.EnvUtilHelper{}

func GenerateJwt(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = email
	claims["exp"] = time.Now().Add(720 * time.Hour).Unix()
	signedString := envUtil.GetEnv("SECRET_STRING", loader)
	tokenString, err := token.SignedString([]byte(signedString))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJwt(w http.ResponseWriter, r *http.Request, mongo types.MongoClient) (string, error) {
	fmt.Println("Validating token...", r.Header["Token"][0])

	if r.Header["Token"] == nil {
		fmt.Fprintf(w, "Missing token")
		return "", errors.New("missing token")
	}

	signedString := envUtil.GetEnv("SECRET_STRING", loader)

	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error parsing token inside parser")
		}
		return []byte(signedString), nil
	})

	if token == nil {
		fmt.Println("invalid token")
		return "", errors.New("invalid token")
	}
	fmt.Println("Validated and parsed token")
	claims, ok := token.Claims.(jwt.MapClaims)

	if err != nil || !ok {
		fmt.Println("Error parsing token: ", err)
		fmt.Println(w, "error parsing token")
		return "", errors.New("error parsing token")
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		fmt.Println("Token epired")
		return "", errors.New("Unauthorized")
	}

	email := claims["username"].(string)
	fmt.Println("email: ", email)
	mongoToken, err := mongo.GetTokenByEmail(email)
	fmt.Println("Mongo token: ", mongoToken)
	if err != nil {
		fmt.Println("Error getting token by email")
		return "", err
	}
	fmt.Println("Returning validator email: ", email)
	return email, nil
}

func CreateJwt(email string, mongo types.MongoClient) (string, error) {
	token, err := GenerateJwt(email)
	if err != nil {
		return "", err
	}
	// save the token
	err = mongo.InsertToken(email)
	if err != nil {
		return "", nil
	}
	return token, nil
}

func PasswordMatches(plainText string, originalPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			log.Println("Error validating passsowrd: %s", err)
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
