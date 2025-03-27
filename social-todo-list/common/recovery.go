package common

import "log"

func Recovery() {
	if err := recover(); err != nil {
		log.Println("Recovered: ", err)
	}
}
