package dialogflow

// Config for the router
type Config struct {
	Debug bool
	Token string
}

func (r *Router) IsDebug() bool {
	return r.config.Debug
}
