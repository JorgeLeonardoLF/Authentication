package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JorgeLeonardoLF/Authentication/database"
	"github.com/JorgeLeonardoLF/Authentication/models"
)

/*RegisterNewAccount should take in some json payload -> sanitize it -> send it over to the db -> return some json result*/
func RegisterNewAccount(w http.ResponseWriter, r *http.Request) {
	/*We get our data from the request (r) body, it is encoded there for we need:
	- To put it into a json decoder
	- decode it into our expected structur*/
	decoder := json.NewDecoder(r.Body)
	newAccount := models.NewAccount{}
	onDecodingErr := decoder.Decode(&newAccount)

	if onDecodingErr != nil {
		RespondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", onDecodingErr))
		return
	}

	statusCode, onCreateErr := database.ApiDbCfg.CreateNewAccountRecord(newAccount)
	if onCreateErr != nil {
		RespondWithError(w, statusCode, fmt.Sprintf("error couldn't create user: %v", onCreateErr))
		return
	}

	RespondWithJSON(w, statusCode, struct{}{})
}
