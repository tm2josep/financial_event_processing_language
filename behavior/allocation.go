package behavior

import (
	"errors"
)

type hasValue interface {
	getValue(map[string]interface{}) (float64, bool)
}

type Allocation struct {
	Source Field
	Value  hasValue
	Target Field
}

func max(x ...float64) float64 {
	best := x[0]
	for _, v := range x {
		if v > best {
			best = v
		}
	}
	return best
}

func min(x ...float64) float64 {
	best := x[0]
	for _, v := range x[0:] {
		if v < best {
			best = v
		}
	}
	return best
}

func getFieldValue(loss map[string]interface{}, fieldName string) (float64, error) {
	value, hasValue := loss[fieldName]
	if !hasValue {
		err := errors.New("field value not present")
		return 0, err
	}

	newValue, ok := value.(float64)
	if !ok {
		err := errors.New("field not a number")
		return 0, err
	}

	return newValue, nil
}

func (alloc *Allocation) getPertinentValues(loss map[string]interface{}) (float64, float64, float64, error) {
	source, err := getFieldValue(loss, alloc.Source.Name)
	if err != nil {
		return 0, 0, 0, err
	}

	target, err := getFieldValue(loss, alloc.Target.Name)
	if err != nil {
		return 0, 0, 0, err
	}

	value, hasValue := alloc.Value.getValue(loss)
	if !hasValue {
		err := errors.New("invalid value field")
		return 0, 0, 0, err
	}

	return source, target, value, nil
}

func (alloc *Allocation) Apply(loss map[string]interface{}) (map[string]interface{}, error) {

	source, target, value, err := alloc.getPertinentValues(loss)
	if err != nil {
		return nil, err
	}

	quantity := min(source, value)
	source = source - quantity
	target = target + quantity
	loss[alloc.Source.Name] = source
	loss[alloc.Target.Name] = target
	return loss, nil
}
