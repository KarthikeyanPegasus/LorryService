package SignUp

import (
	"golang.org/x/crypto/bcrypt"
	"lorryservice/Database"
	"lorryservice/Types"
	"net/http"
)

func createUser(w http.ResponseWriter, user Types.User) error {
	db := Database.ConnectDB()
	defer db.Close()

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "unable to generate hash", http.StatusInternalServerError)
		return err
	}

	var id int
	err = db.QueryRow(`SELECT id FROM "user".users WHERE username = $1`, user.Username).Scan(&id)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		http.Error(w, "username already exist", http.StatusConflict)
		return err
	}

	// Insert the new user into the database
	_, err = db.Exec(`INSERT INTO "user".users (name,username, password,roles) VALUES ($1, $2, $3, $4)`, user.Name, user.Username, string(hashedPassword), user.Roles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "issue with connecting db", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}
