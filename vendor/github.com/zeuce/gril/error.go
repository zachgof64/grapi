package utils

import (
	"fmt"
)

//Check will log output to console on error
func Check(err error) {
	if(err != nil){
		fmt.Println(err)
		panic(err)
	}
} 