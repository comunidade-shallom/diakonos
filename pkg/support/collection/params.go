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

func (p Params) Strings(key string) []string {
	raw, exist := p[key]

	if !exist || raw == nil {
		return []string{}
	}

	val, ok := raw.([]interface{})

	if !ok {
		return []string{}
	}

	list := make([]string, len(val))

	for index, row := range val {
		list[index], _ = row.(string)
	}

	return list
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
