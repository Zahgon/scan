package scan

import (
	"reflect"
	"sync"
)

var valuesCache cache = &sync.Map{}

// Values scans a struct and returns the values associated with the columns
// provided. Only simple value types are supported (i.e. Bool, Ints, Uints,
// Floats, Interface, String)
func Values(cols []string, v interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadFields(val reflect.Value) map[string][]int { _ = "STUB: not implemented"; return nil }

func writeFieldsCache(val reflect.Value) map[string][]int { _ = "STUB: not implemented"; return nil }

func writeFields(val reflect.Value, m map[string][]int, index []int) {
	_ = "STUB: not implemented"
	return
}
