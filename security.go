package dialogflow

import (
	"net/http"
	"log"
)

func (r *Router) isFromDialogflow(w http.ResponseWriter, req *http.Request) bool {
	if req.Method != http.MethodPost {
		log.Printf("Method should be POST, %s given", req.Method)
		w.WriteHeader(http.StatusForbidden)
		return false
	}

	t := req.URL.Query().Get("token")

	if t != r.config.Token {
		log.Printf("Invalid token given (%s)", t)
		w.WriteHeader(http.StatusForbidden)
		return false
	}

	return true
}
