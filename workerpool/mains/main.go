package main

import "chatapp/workerpool"

func main() {
	/*user := workerpool.NewPoolUser()
	user.Execute()*/

	oddEven := workerpool.NewOddEven(2)
	oddEven.Print()
}
