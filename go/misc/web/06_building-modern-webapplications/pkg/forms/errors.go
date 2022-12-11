package forms

type errors map[string][]string

func (e errors) AddOne(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) GetOne(field string) string {
	values := e[field]
	if len(values) == 0 {
		return ""
	}
	return values[0]
}