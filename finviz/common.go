package finviz

import (
	"fmt"
)

func arrayToString(arr []string) string {
	final := ""
	for index, element := range arr {
		if index != 0 {
			final = final + ","
		}
		final = final + string(element)

	}

	return final
}

func printarr(arr []string) {
	for _, element := range arr {
		fmt.Println(element)
	}
	fmt.Println("-----------------------------")
}
