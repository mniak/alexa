package alexa

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	alexakit "github.com/ericdaugherty/alexa-skills-kit-golang"
)

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	log.Println(err.Error())
	fmt.Fprintf(w, `{"message": "Internal error. Check the logs.")`)
}

func (sf skillFunc) MakeHTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		var requestEnvelope *alexakit.RequestEnvelope
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&requestEnvelope)
		if err != nil {
			handleError(w, err)
			return
		}

		responseEnvelope, err := sf(r.Context(), requestEnvelope)
		if err != nil {
			handleError(w, err)
			return
		}

		enc := json.NewEncoder(w)
		err = enc.Encode(responseEnvelope)
		if err != nil {
			handleError(w, err)
			return
		}
	}
}
