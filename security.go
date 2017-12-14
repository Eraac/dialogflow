package dialogflow

import "net/http"

func (r *Router) isFromDialogflow(w http.ResponseWriter, req *http.Request) bool {
	if req.Method != http.MethodPost {
		r.config.Logger.Warn("Method should be POST, %s given", req.Method)
		w.WriteHeader(http.StatusForbidden)
		return false
	}

	t := req.URL.Query().Get("token")

	if t != r.config.Token {
		r.config.Logger.Warn("Invalid token given (%s)", t)
		w.WriteHeader(http.StatusForbidden)
		return false
	}

	return true
}
