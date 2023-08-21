package tarot

import (
	"fmt"
	"testing"
	"time"
)

func TestJsonToMap(t *testing.T) {
	JsonToMap()
	fmt.Println(GetTarotCard())

}

func TestAppendLog(t *testing.T) {
	AppendLog(time.Now(), "32")
}

func TestCheckLog(t *testing.T) {
	checkLog, ok := CheckLog()
	fmt.Println(checkLog, ok)
}
