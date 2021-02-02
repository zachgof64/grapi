package utils

import (
	"fmt"
)

//Check - Prints error if error not nil
func Check(err error) {
	if(err != nil){
		fmt.Println(err)
		panic(err)
	}
} 