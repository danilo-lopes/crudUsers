package services

import (
	"bufio"
	"os"
	"strings"
)

const (
	passwdFile string = "/etc/passwd"
)

func GetExistentUsers() []string {

	var osUsers []string

	file, err := os.Open(passwdFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		etcLines := reader.Text()
		etcSplited := strings.Split(etcLines, ":")
		osUsers = append(osUsers, etcSplited[0])
	}

	return osUsers
}

func compare(u []string, c string) bool {
	for _, osUser := range u {
		if osUser == c {
			return true
		}
	}

	return false
}
