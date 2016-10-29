package sorting

import (
	"log"

	"github.com/GoAlgorithms/base"
)

func InsertionSort(slice []interface{}) []interface{} {
	if len(slice) > 0 {
		if _, ok := slice[0].(int); ok {
			//return intSort(base.ConvertSliceToIntSlice(slice))
			return nil
		} else if _, ok := slice[0].(float64); ok {
			return nil
		} else if _, ok := slice[0].(string); ok {
			return nil
		} else {
			log.Println("DATA TYPE NOT SUPPORTED")
			return nil
		}
	}
	log.Println("SLICE MUST HAVE MINIMUM ONE VALUE")
	return nil
}

func IntSort(slice []int) []int {
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := i; j > 0 && base.IntLess(slice[j], slice[j-1]); j-- {
			slice = base.IntegerSliceExchange(slice, j, j-1)
		}
	}
	return slice
}
