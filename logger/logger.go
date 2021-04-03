package logger

import (
	"os"
	"strings"
	"time"
)

func Log(mgs ...string) {

	file, err := os.OpenFile("_application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic("Cannot Open/Create Log File")
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:5 ") + strings.Join(mgs, "") + "\n")

	file.Close()
}
