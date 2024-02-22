package utils

import "fmt"

func Filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func ToStrings[T fmt.Stringer](list []T) []string {
	stringArray := make([]string, len(list))
	for i, s := range list {
		stringArray[i] = s.String()
	}
	return stringArray
}

func ConvertList[T, V any](inputList []T, converter func(T) V) []V {
	var outputList []V
	for _, item := range inputList {
		outputList = append(outputList, converter(item))
	}
	return outputList
}

func ToList[K comparable, V any](m map[K]V) []V {
  list := make([]V, 0, len(m))
  for _, v := range m {
    list = append(list, v)
  }
  return list
}