package SignUp

import (
	"encoding/json"
	"lorryservice/Types"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
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

	if err := createUser(w, user); err != nil {
		return
	}

}
