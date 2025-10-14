package main

import (
	"fmt"

)



func main(){
	user1 := Reader{
		ID: 1,
		FirstName: "Agunda",
		LastName: "Kokoit`i",
		IsActive: true,
	}

	user1.Deactivate()
	//user1.DisplayReader()
	fmt.Println(user1)

}