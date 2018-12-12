package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// User defines the user data
type User struct {
	Name     string
	Password string
	Uid      string
	Gid      string
	Gecos    string
	Home     string
	Shell    string
}

// users is a slice of User structs
var users []User

// userStrings creates a byte slice of the command output
func userStrings() ([]string, error) {
	cmd := exec.Command("/usr/bin/getent", "passwd")
	res, err := cmd.Output()
	if err != nil {
		fmt.Println("Error processing the passwd list")
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("Cannot process empty string")
	}
	return strings.Split(string(res), string('\n')), nil
}

// usersFull return a []User slice of users
func usersFull(u []string) []User {
	for i := 0; i < len(u)-1; i++ {
		splitted := strings.Split(u[i], string(':'))
		users = append(users, User{
			Name:     splitted[0],
			Password: splitted[1],
			Uid:      splitted[2],
			Gid:      splitted[3],
			Gecos:    splitted[4],
			Home:     splitted[5],
			Shell:    splitted[6],
		})
	}
	return users
}

func main() {
	// Obtain the []string user slice
	userSlice, err := userStrings()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Load all users in a []User slice
	users := usersFull(userSlice)

	// Create a final map embedding the users struct
	finalMap := map[string]map[string][]User{"passwd": {"content": users}}

	// Marshal results to json object
	b, err := json.Marshal(finalMap)
	if err != nil {
		fmt.Println("Cannot marshal results")
		os.Exit(1)
	}
	fmt.Println(string(b))
}
