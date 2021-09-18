package logger_test

import (
	"fmt"
	"testing"

	"github.com/Sotaneum/go-logger"
)

func TestRunner(t *testing.T) {
	log := logger.New("./", "log.json")

	log.Add("test")
	log.Add("test1")
	log.Add("test2")

	fmt.Println(log.Get())

	log.Remove()
}
