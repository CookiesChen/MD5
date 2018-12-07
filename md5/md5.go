package md5

import "fmt"

func Exec(input string){
	bytes := []byte(input)
	message := padding(bytes[:])
	fmt.Println(message)
}