package dialogflow

import "net/http"

func (r *Router) isFromDialogflow(w http.ResponseWriter, req *http.Request) bool {
	if req.Method != http.MethodPost {
		r.config.Logger.WithField("method", req.Method).Warn("Method should be POST")
		w.WriteHeader(http.StatusForbidden)
		return false
	}

	t := req.URL.Query().Get("token")

	if t != r.config.Token {
		r.config.Logger.WithField("token", t).Warn("Invalid token given")
		w.WriteHeader(http.StatusForbidden)
		return false
	}

	return true
}
