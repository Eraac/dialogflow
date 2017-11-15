package dialogflow

// Find context by name
func (c Contexts) Find(name string) (*Context, error) {
	cs := []Context(c)

	for _, c := range cs {
		if c.Name == name {
			return &c, nil
		}
	}

	return nil, ErrNotFound
}

func (r *Response) ResetContext() {
	r.ContextOut = Contexts{
		{Name: "generic", Parameters: Parameters{}, Lifespan: 0},
	}
}
