package collection

type Params map[string]interface{}

func (p Params) String(key string) string {
	raw, exist := p[key]

	if !exist || raw == nil {
		return ""
	}

	val, ok := raw.(string)

	if ok {
		return val
	}

	return ""
}

func (p Params) Int(key string) int {
	raw, exist := p[key]

	if !exist || raw == nil {
		return 0
	}

	val, ok := raw.(int)

	if ok {
		return val
	}

	return 0
}
