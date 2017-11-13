package dialogflow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type (
	Router struct {
		handler map[string]Handler
		debug   bool
	}

	Handler func(r *Request) (*Response, error)
)

func NewRouter(debug bool) *Router {
	return &Router{
		handler: map[string]Handler{},
		debug:   debug,
	}
}

func (r *Router) HandleFunc(action string, h Handler) {
	r.handler[action] = h
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.debug {
		bs, _ := httputil.DumpRequest(req, true)

		fmt.Printf("%s\n", string(bs))
	}

	dfReq := &Request{}

	bs, err := ioutil.ReadAll(req.Body)

	if httpError(err, w) {
		return
	}

	err = json.Unmarshal(bs, dfReq)

	if httpError(err, w) {
		return
	}

	h, ok := r.handler[dfReq.Result.Action]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		log.Println(fmt.Sprintf("action not found: %s\n", dfReq.Result.Action))
		return
	}

	res, err := h(dfReq)

	if httpError(err, w) {
		return
	}

	bs, err = json.Marshal(res)

	if httpError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)

	if r.debug {
		fmt.Printf("\n===Response===\n\n%s\n", string(bs))
	}
}

func httpError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	w.WriteHeader(http.StatusInternalServerError)
	log.Println(err.Error())

	return true
}
