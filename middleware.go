package dialogflow

func MiddlewarePermission(next Handler, callback, reason string, permissions ...string) Handler {
	return func(ctx *Context) (*Response, error) {
		res := NewResponse()

		if !ctx.GetGoogleData().IsGranted(permissions...) {
			res.AskForPermission(callback, reason, permissions...)

			return res, nil
		}

		return next(ctx)
	}
}
