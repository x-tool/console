package console

import(
	"log"
)

type Console struct {}

func New() (Console){
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