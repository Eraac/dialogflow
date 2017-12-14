package dialogflow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	Router struct {
		handler map[string]Handler
		config  Config
	}

	Handler func(r *Request) (*Response, error)
)

func NewRouter(c Config) *Router {
	if c.Logger == nil {
		c.Logger = logrus.New()
	}

	return &Router{
		handler: map[string]Handler{},
		config:  c,
	}
}

func (r *Router) HandleFunc(action string, h Handler) {
	r.handler[action] = h
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	n := time.Now()

	w.Header().Set("Content-Type", "application/json")

	if r.IsDebug() {
		bs, _ := httputil.DumpRequest(req, true)

		fmt.Printf("%s\n", string(bs))
	}

	if !r.isFromDialogflow(w, req) {
		r.config.Logger.WithField("status", "unauthorized").Warn("unauthorized")
		return
	}

	dfReq := &Request{}

	bs, err := ioutil.ReadAll(req.Body)

	if httpError(err, w) {
		r.config.Logger.WithField("status", "error").Error(err.Error())
		return
	}

	err = json.Unmarshal(bs, dfReq)

	if httpError(err, w) {
		r.config.Logger.WithField("status", "error").Error(err.Error())
		return
	}

	logger := r.config.Logger.WithFields(logrus.Fields{
		"action":     "bot_interaction",
		"intent":     dfReq.Result.Action,
		"source":     dfReq.Source(),
		"session_id": dfReq.SessionID,
		"user_id":    dfReq.GetUserID(),
		"user_ask":   dfReq.Result.ResolvedQuery,
	})

	h, ok := r.handler[dfReq.Result.Action]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		logger.WithField("status", "not found").Warn("action not found")
		return
	}

	res, err := h(dfReq)

	if httpError(err, w) {
		logger.WithField("status", "error").Error(err.Error())
		return
	}

	bs, err = json.Marshal(res)

	if httpError(err, w) {
		logger.WithField("status", "error").Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)

	logger.WithFields(logrus.Fields{
		"response_time": time.Since(n).Seconds(),
		"status":        "success",
	}).Info("success")

	if r.IsDebug() {
		fmt.Printf("\n===Response===\n\n%s\n\n", string(bs))
	}
}

func httpError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	w.WriteHeader(http.StatusInternalServerError)

	return true
}
