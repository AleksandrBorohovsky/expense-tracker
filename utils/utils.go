package utils

import "fmt"

func Check(err error, message string) {
	if err != nil {
		fmt.Println(err, message)
		return
	}

}
