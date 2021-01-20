package utils

import "fmt"

//Check - Outputs if passed error isn't nil
func Check(err error) {
	if(err != nil) {
		fmt.Println(err)
	}
} 