package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/CookiesChen/MD5/md5"
)

func main(){

	fmt.Println("CookiesChen's MD5 Algorithm")
	fmt.Println("Please Input message:")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Your Input is " + input)

	output :=  md5.Exec(input)

	fmt.Println()
	fmt.Print("MD5 Hash: ")
	for i := 0; i < 4; i++ {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint32(b, output[i])
		bytesBuffer := bytes.NewBuffer(b)
		binary.Read(bytesBuffer, binary.BigEndian, &output[i])
		fmt.Printf("%x", output[i])
	}
	fmt.Println()

}