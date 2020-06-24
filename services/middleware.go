package services

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mauriliommachado/go-commerce/product-service/models"
)

var app *models.App

//InitMiddleware initate the middleware with app configs
func InitMiddleware(config *models.App) {
	app = config
}

func protectMiddleware(next http.HandlerFunc) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		authorizationHeader := strings.Split(req.Header.Get("Authorization"), " ")[1]
		if authorizationHeader != "" {
			if validateToken(authorizationHeader) {
				next(w, req)
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.Exception{Message: "Invalid Authorization token"})
		return
	})
}

func validateToken(token string) bool {
	resp, err := http.Get(app.AuthService + "/" + token)
	if err != nil {
		print(err)
	}
	return resp.StatusCode == 200
}
