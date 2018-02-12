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

	Handler func(ctx *Context) (*Response, error)
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

		fmt.Printf("\n===Request===\n\n%s\n", string(bs))
	}

	if !r.isFromDialogflow(w, req) {
		r.config.Logger.WithField("status", "unauthorized").Warn("unauthorized")
		return
	}

	ctx := &Context{}

	bs, err := ioutil.ReadAll(req.Body)

	if httpError(err, w) {
		r.config.Logger.WithField("status", "error").Error(err)
		return
	}

	if err = json.Unmarshal(bs, &ctx.Request); httpError(err, w) {
		r.config.Logger.WithField("status", "error").Error(err)
		return
	}

	ctx.Logger = r.config.Logger.WithFields(logrus.Fields{
		"action":     "bot_interaction",
		"intent":     ctx.Request.Result.Action,
		"source":     ctx.Source(),
		"session_id": ctx.Request.SessionID,
		"user_id":    ctx.GetUserID(),
		"user_ask":   ctx.Request.Result.ResolvedQuery,
	})

	h, ok := r.handler[ctx.Request.Result.Action]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.WithField("status", "not found").Warn("action not found")
		return
	}

	res, err := h(ctx)

	if httpError(err, w) {
		ctx.Logger.WithField("status", "error").Error(err)
		return
	}

	bs, err = json.Marshal(res)

	if httpError(err, w) {
		ctx.Logger.WithField("status", "error").Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)

	ctx.Logger.WithFields(logrus.Fields{
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
