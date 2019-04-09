package main

import (
	"fmt"
	"os"
)

const (
	usage = "Usage: [username] [password]"
	errUsage = "Access denied for %q\n"
	errPwd = "Invalid password for %q\n"
	accessOk = "Access granted to %q\n"

	user, user2 = "stkhr", "hoge"
	pass, pass2 = "1111", "2222"
)

func main()  {
	var (
		args = os.Args
		l = len(args) - 1
	)
	if l != 2 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u != user && u != user2 {
		fmt.Printf(errUsage, args[1])
	} else if u == user && p == pass {
		fmt.Printf(accessOk, user)
	} else if u == user2 && p == pass2 {
		fmt.Printf(accessOk, user2)
	} else {
		fmt.Printf(errPwd, user)
	}
}
