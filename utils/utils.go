package utils

import (
	"fmt"
	"log"
)

func CheckNilErr(err error, customMessage ...string) {
	if err != nil {
		if len(customMessage) > 0 {
			fmt.Println(customMessage[0])
		} else {
			log.Fatal(err)
		}
	}
}
