package dialogflow

import (
	"reflect"
	"strconv"
)

// HasKey return true if key exist and is not empty
func (p Parameters) HasKey(key string) bool {
	m := map[string]interface{}(p)
	v, ok := m[key]

	if !ok {
		return false
	}

	str, ok := v.(string)

	return ok && str != ""
}

// GetInt return parameter `key` in int
func (p Parameters) GetInt(key string) (int, error) {
	str, err := p.GetString(key)

	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0, err
	}

	return int(f), nil
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

// GetSliceString return parameter `key` in slice of string
func (p Parameters) GetSliceString(key string) ([]string, error) {
	item, err := p.GetParameter(key)

	if err != nil {
		return []string{}, err
	}

	var ss []string

	s := reflect.ValueOf(item)

	for i := 0; i < s.Len(); i++ {
		ss = append(ss, s.Index(i).Interface().(string))
	}

	return ss, nil
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
