package behavior

type TermValue struct {
	CurrentValue float64
}

type Field struct {
	Name string
}

func (field *Field) getValue(loss map[string]interface{}) (interface{}, bool) {
	value, hasKey := loss[field.Name]
	return value, hasKey
}

func (value *TermValue) getValue(loss map[string]interface{}) (float64, bool) {
	return value.CurrentValue, true
}
