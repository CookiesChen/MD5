package main

import (
	"fmt"
	"github.com/CookiesChen/MD5/md5"
)

func main(){

	fmt.Println("CookiesChen's MD5 Algorithm")
	fmt.Println("Please Input message:")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Your Input is " + input)

	md5.Exec(input)

}