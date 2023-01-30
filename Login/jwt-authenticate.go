package Login

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	ct "lorryservice/Constants"
	"lorryservice/Database"
	"lorryservice/Types"
	"net/http"
	"time"
)

func authenticate(w http.ResponseWriter, user Types.User) error {

	db := Database.ConnectDB()
	defer db.Close()

	// Get the hashed password from the database
	var hashedPassword string
	err := db.QueryRow(`SELECT password FROM "user".users WHERE username=$1`, user.Username).Scan(&hashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "username doesn't exist", http.StatusUnauthorized)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "invalid username and password", http.StatusUnauthorized)
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(ct.JWTSecret))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Error While signing the JWT", http.StatusInternalServerError)
		return err
	}

	// Return the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})

	return nil
}
