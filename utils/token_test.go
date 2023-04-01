package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	password, _ := bcrypt.GenerateFromPassword([]byte("123"), 12)
	fmt.Println(string(password))
}
