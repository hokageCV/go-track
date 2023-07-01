package utils

import (
	"fmt"
	"log"
)

func CheckNilErr(err error, customMessage string) {
	if err != nil {
		if customMessage != "" {
			fmt.Println(customMessage)
		} else {
			log.Fatal(err)
		}
	}
}
