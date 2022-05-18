package betterhandler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	log.Println(err.Error())
	fmt.Fprintf(w, "<pre>Internal error. Check the logs.</pre>")
}

func Make(handlerFunc interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler := lambda.NewHandler(handlerFunc)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handleError(w, err)
			return
		}

		result, err := handler.Invoke(r.Context(), body)
		if err != nil {
			handleError(w, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(result)
	}
}
