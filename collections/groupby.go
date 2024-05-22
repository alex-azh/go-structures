package collections

import "unsafe"

func GroupBy[TValue, TObject comparable](object *TObject, v *TValue, arr []*TObject) map[TValue][]*TObject {
	getField := getFieldByPtr(object, v)

	switch len(arr) {
	case 0:
		return make(map[TValue][]*TObject)
	case 1:
		return map[TValue][]*TObject{getField(arr[0]): {arr[0]}}
	}

	result := make(map[TValue][]*TObject, len(arr)*4/10)
	for i := 0; i < len(arr); i++ {
		v := getField(arr[i])
		result[v] = append(result[v], arr[i])
	}
	return result
}

func getFieldByPtr[TValue, TObject comparable](object *TObject, value *TValue) func(obj *TObject) TValue {
	delta := uintptr(unsafe.Pointer(value)) - uintptr(unsafe.Pointer(object))
	return func(obj *TObject) TValue {
		return *(*TValue)(unsafe.Pointer(uintptr(unsafe.Pointer(obj)) + delta))
	}
}
