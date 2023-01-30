package Login

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	ct "lorryservice/Constants"
	"net/http"
)

func JWTValidate(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ct.JWTSecret, nil
	})
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
