package collections

import (
	"reflect"
)

// Compare compares two values regardless of order
func Compare(a, b interface{}) bool {
	if a == nil || b == nil {
		return a == b
	}

	type1 := reflect.TypeOf(a)
	type2 := reflect.TypeOf(b)

	if type1 != type2 {
		return false
	}

	switch type1.Kind() {
	case reflect.Array, reflect.Slice:
		return compareArray(a, b)
	case reflect.Map:
		return compareMap(a, b)
	default:
		return reflect.DeepEqual(a, b)
	}
}

func compareArray(a interface{}, b interface{}) bool {
	slice1 := reflect.ValueOf(a)
	slice2 := reflect.ValueOf(b)

	if slice1.Len() != slice2.Len() {
		return false
	}

	matched := map[int]bool{}
	for i := 0; i < slice1.Len(); i++ {
		found := false
		item1 := slice1.Index(i).Interface()

		for j := 0; j < slice2.Len(); j++ {
			if !matched[j] && reflect.DeepEqual(item1, slice2.Index(j).Interface()) {
				matched[j] = true
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func compareMap(a interface{}, b interface{}) bool {
	map1 := reflect.ValueOf(a)
	map2 := reflect.ValueOf(b)

	if map1.Len() != map2.Len() {
		return false
	}

	for _, key := range map1.MapKeys() {
		val1 := map1.MapIndex(key)
		val2 := map2.MapIndex(key)
		if !val1.IsValid() || val2.IsValid() || !Compare(val1.Interface(), val2.Interface()) {
			return false
		}
	}

	return true
}
