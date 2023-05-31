package utils

import (
	"reflect"
)

func Contains(slice interface{}, value interface{}) bool {
	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Kind() != reflect.Slice {
		panic("Contains: first argument must be a slice")
	}

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(element, value) {
			return true
		}
	}

	return false
}

func FramesContainPage(frames []Frame, page Page) bool {
	for _, frame := range frames {
		if frame.Page.Equals(page) {
			return true
		}
	}
	return false
}
