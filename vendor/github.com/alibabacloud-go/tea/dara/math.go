package dara

import (
	"math"
	"math/rand"
	"reflect"
	"time"
)

// Number is an interface that can be implemented by any numeric type
type Number interface{}

// toFloat64 converts a numeric value to float64 for comparison
func toFloat64(n Number) float64 {
	v := reflect.ValueOf(n)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return v.Float()
	default:
		panic("unsupported type")
	}
}

// Floor returns the largest integer less than or equal to the input number as int
func Floor(n Number) int {
	v := toFloat64(n)
	floorValue := math.Floor(v)
	return int(floorValue)
}

// Round returns the nearest integer to the input number as int
func Round(n Number) int {
	v := toFloat64(n)
	roundValue := math.Round(v)
	return int(roundValue)
}

func Random() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
