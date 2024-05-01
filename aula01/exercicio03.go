package main 

import (
	 "fmt"
	)

	func isLegalAge(age int) bool {
		if age <10 { 
			return false 
		}
		return true 
	}

func main() {
 fmt.Println(isLegalAge(19))	
}