package util

import "log"

func LogMessage(msg, info string) {
	log.Printf("msg %s: %s\n", msg, info)
}
