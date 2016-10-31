package base

import "strings"

type Integer int
type Float float64
type String string

type Calculator interface {
	Less(a, b interface{}) bool
	Exchange(s, i, j interface{}) interface{}
}

func IntLess(a, b int) bool {
	return a < b
}

func FloatLess(a, b float64) bool {
	return a < b
}

func StringLess(a, b string) bool {
	return strings.Compare(a, b) < 0
}

func IntegerSliceExchange(slice []int, i, j int) []int {
	aux := slice[i]
	slice[i] = slice[j]
	slice[j] = aux
	return slice
}

func ConvertSliceToIntSlice(slice []interface{}) []int {
	var intSlice []int = make([]int, len(slice))
	for i, d := range slice {
		if value, ok := d.(int); ok {
			intSlice[i] = value
		}
	}
	return intSlice
}

func ConertIntSliceToInterface(slice []int) []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(slice))
	for i, d := range slice {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
