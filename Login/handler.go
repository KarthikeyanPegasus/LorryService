package Login

import (
	"encoding/json"
	"lorryservice/Types"
	"net/http"
)

// @Title Login
// @Description Login user with username and password and returns the JWT. which can be used for future references.
// @Params request body types.User
// @Success 200 {object} jwt token
// @Router /login [post]

func Login(w http.ResponseWriter, r *http.Request) {
	var user Types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := authenticate(w, user); err != nil {
		return
	}
}
