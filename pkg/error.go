package pkg

import "log"

func HandleError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
