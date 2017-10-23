package console

import (
	"log"
)

type Console struct{}

func New() Console {
	return Console{}
}

// panic
func (c *Console) Panic(errorType string, errorMsg error) {
	panic("!!!!!!!!!! [ERROR] " + errorType + " | " + errorMsg.Error() + " !!!!!!!!!!")
}

// run test
func (c *Console) Inspect(msg string) {
	log.Print("---------- [SUCCESS] " + msg + " ----------")
}

// db console
func (c *Console) DBError(errorMsg error) {
	log.Print("---------- [DB] " + errorMsg.Error() + " ----------")
}

// sinple Log
func (c *Console) Log(s ...string) {
	if len(s) == 0 {

	} else if len(s) == 1 {
		log.Print(s[0])
	} else {
		var _s string
		for i, v := range s {
			if len(s) == (i + 1) {
				_s = _s + v
			} else {
				_s = _s + v + " | "
			}
		}
		log.Print(_s)
	}
}

// sinple Log
func (c *Console) LogWithLabel(s ...string) {
	if len(s) == 0 {

	} else if len(s) == 1 {
		log.Print(s[0])
	} else {
		var _s string
		for i, v := range s {
			if len(s) == (i + 1) {
				_s = _s + v
			} else if i == 0 {
				_s = "[" + v + "] "
			} else {
				_s = _s + v + " | "
			}
		}
		log.Print(_s)
	}
}
