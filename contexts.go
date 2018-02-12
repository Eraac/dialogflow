package dialogflow

// Find context by name
func (c Contexts) Find(name string) (*DFContext, error) {
	cs := []DFContext(c)

	for _, c := range cs {
		if c.Name == name {
			return &c, nil
		}
	}

	return nil, ErrNotFound
}

// ResetContext remove all current contexts
func (r *Response) ResetContext(req *Request) {
	contexts := Contexts{}

	for _, c := range req.Result.Contexts {
		contexts = append(contexts, DFContext{Name: c.Name, Lifespan: 0})
	}

	r.ContextOut = contexts
}
