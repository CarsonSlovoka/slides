package funcs

import (
	"fmt"
	"math"
	"reflect"
)

func MathMaps() Maps {
	return Maps{
		"add": Add,
		"mul": Mul,
		"mod": Mod,
		"pow": Pow,
		"min": Min,
		"max": Max,
	}
}

func arithmetic(values []any, computeFunc func(a *reflect.Value, b reflect.Value) error) (any, error) {
	if len(values) == 0 {
		return 0, nil
	}
	firstValue := reflect.ValueOf(values[0])
	resultValue := reflect.New(firstValue.Type()).Elem()
	resultValue.Set(firstValue)
	for _, value := range values[1:] {
		v := reflect.ValueOf(value)
		if v.Type() != firstValue.Type() {
			return nil, fmt.Errorf("type not match first: %T", values[0])
		}
		if err := computeFunc(&resultValue, v); err != nil {
			return nil, err
		}
	}
	return resultValue.Interface(), nil
}

func Add(values ...any) (any, error) {
	return arithmetic(values, add)
}

// a為 *reflect.Value 讓他成為Assignable的對象，不然會報錯
func add(a *reflect.Value, b reflect.Value) error {
	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a.SetInt(a.Int() + b.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		a.SetUint(a.Uint() + b.Uint())
	case reflect.Float32, reflect.Float64:
		a.SetFloat(a.Float() + b.Float())
	default:
		return fmt.Errorf("invalid Kind: %d", a.Kind())
	}
	return nil
}

func Mul(values ...any) (any, error) {
	return arithmetic(values, mul)
}

func mul(a *reflect.Value, b reflect.Value) error {
	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a.SetInt(a.Int() * b.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		a.SetUint(a.Uint() * b.Uint())
	case reflect.Float32, reflect.Float64:
		a.SetFloat(a.Float() * b.Float())
	default:
		return fmt.Errorf("invalid Kind: %d", a.Kind())
	}
	return nil
}

func Mod(a, b int) int {
	return a % b
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Min(values ...any) (any, error) {
	return arithmetic(values, _min)
}

func _min(a *reflect.Value, b reflect.Value) error {
	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if a.Int() > b.Int() {
			a.SetInt(b.Int())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if a.Uint() > b.Uint() {
			a.SetUint(b.Uint())
		}
	case reflect.Float32, reflect.Float64:
		if a.Float() > b.Float() {
			a.SetFloat(b.Float())
		}
	default:
		return fmt.Errorf("invalid Kind: %d", a.Kind())
	}
	return nil
}

func Max(values ...any) (any, error) {
	return arithmetic(values, _max)
}

func _max(a *reflect.Value, b reflect.Value) error {
	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if a.Int() < b.Int() {
			a.SetInt(b.Int())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if a.Uint() < b.Uint() {
			a.SetUint(b.Uint())
		}
	case reflect.Float32, reflect.Float64:
		if a.Float() < b.Float() {
			a.SetFloat(b.Float())
		}
	default:
		return fmt.Errorf("invalid Kind: %d", a.Kind())
	}
	return nil
}
