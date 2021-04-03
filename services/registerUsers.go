package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/danilo-lopes/crudUsers/logger"
	"github.com/danilo-lopes/crudUsers/models"
)

func readUsersFile() []byte {
	file, err := os.Open("users.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	data, _ := ioutil.ReadAll(file)

	return data
}

func addUser(u *models.User) (bool, string) {

	password := "Dan12345678@."

	argUser := []string{"-m", "-d", u.Directory, "-G", u.Group, "-s", u.Shell, u.Name}
	argPass := []string{"-c", fmt.Sprintf("echo %s:%s | chpasswd", u.Name, password)}

	userCmd := exec.Command("useradd", argUser...)
	passCmd := exec.Command("/bin/sh", argPass...)

	_, userCmdErr := userCmd.Output()

	if userCmdErr != nil {
		logger.Log("Error when trying do add user: ", fmt.Sprintf(u.Name), "error message from OS: ", userCmdErr.Error())

		return false, userCmdErr.Error()
	}

	_, UserCmdPassErr := passCmd.Output()

	if UserCmdPassErr != nil {
		logger.Log("Error when trying to set user password to: ", fmt.Sprintf(u.Name), "error message from OS: ", UserCmdPassErr.Error())

		return false, UserCmdPassErr.Error()
	}

	return true, password
}

func RegisterUsers() {
	data := readUsersFile()

	var user models.Users

	json.Unmarshal(data, &user)

	for i := range user.Users {
		info, passwd := addUser(&user.Users[i])

		if info {
			fmt.Println("User was added:", user.Users[i].Name, "Password:", passwd)
		}
	}
}
