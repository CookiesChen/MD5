package md5

import "fmt"

var CV = []uint32{0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476}

func Exec(input string) []uint32{
	bytes := []byte(input)
	/* padding */
	message := padding(bytes[:])

	fmt.Print("After Padding: ")
	for i:=0; i<len(message); i++  {
		fmt.Printf("%x ", message[i])
	}
	fmt.Println()

	/* H */
	for i,times := 0,len(message)/16; i < times; i++ {
		compression(message[i*16:(i+1)*16])
	}

	return CV
}