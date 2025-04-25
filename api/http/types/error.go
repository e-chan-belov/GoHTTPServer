package types

import (
	"encoding/json"
	"hw2/domain"
	"log"
	"net/http"
)

func ProcessError(w http.ResponseWriter, code int, loc string, err error) {
	mes := domain.ErrorMessage{Message: err.Error(), Location: loc}
	if err := json.NewEncoder(w).Encode(mes); err != nil {
		log.Println("IF YOU SEE THIS, SOMETHING REALLY WENT WRONG!!!")
		return
	}
	http.Error(w, err.Error(), code)
}
