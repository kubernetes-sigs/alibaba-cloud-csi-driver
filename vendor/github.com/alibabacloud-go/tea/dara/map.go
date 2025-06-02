package dara

import (
	"reflect"
)

type Entry struct {
	Key   string
	Value interface{}
}

// toFloat64 converts a numeric value to float64 for comparison
func Entries(m interface{}) []*Entry {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("Entries: input must be a map")
	}

	entries := make([]*Entry, 0, v.Len())
	for _, key := range v.MapKeys() {
		// 确保 Key 是字符串类型
		if key.Kind() != reflect.String {
			panic("Entries: map keys must be of type string")
		}

		value := v.MapIndex(key)
		entries = append(entries, &Entry{
			Key:   key.String(),
			Value: value.Interface(),
		})
	}
	return entries
}

func KeySet(m interface{}) []string {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("KeySet: input must be a map")
	}

	keys := make([]string, 0, v.Len())
	for _, key := range v.MapKeys() {
		if key.Kind() != reflect.String {
			panic("KeySet: map keys must be of type string")
		}
		keys = append(keys, key.String())
	}
	return keys
}
