package models

import (
	"fmt"
	"testing"
)

func TestSalt(t *testing.T) {
	var pwd Password = "45127996wind"
	var ss = ""
	for len(ss) != 11 {
		ss = pwd.GenerateSalt()
		fmt.Println(ss)
	}

}

func TestUser(t *testing.T) {
	admin := GenerateAdmin()
	fmt.Println(admin)

	fmt.Println(Password("45127996").doCompare(admin.Pwd))
	fmt.Println(Password("wind").doCompare(admin.Pwd))
	fmt.Println(Password("45127996wind").doCompare(admin.Pwd))
}
