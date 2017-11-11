package dialogflow


// GetInt return parameter `key` in int
func (p Parameters) GetInt(key string) (int, error) {
	item, err := p.GetParameter(key)

	if err != nil {
		return 0, err
	}

	i, ok := item.(int)

	if !ok {
		return 0, ErrCastFail
	}

	return i, nil
}

// GetString return parameter `key` in string
func (p Parameters) GetString(key string) (string, error) {
	item, err := p.GetParameter(key)

	if err != nil {
		return "", err
	}

	s, ok := item.(string)

	if !ok {
		return "", ErrCastFail
	}

	return s, nil
}

// GetParameter return parameter `key`
func (p Parameters) GetParameter(key string) (interface{}, error) {
	m := map[string]interface{}(p)

	item, ok := m[key]

	if !ok {
		return nil, ErrNotFound
	}

	return item, nil
}
