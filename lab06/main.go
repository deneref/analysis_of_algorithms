package main
import (
	 "fmt"
	)

func main(){
	env := new_env("data.txt")
	
	shortest := start(env, 100)
	fmt.Printf("Ants algorithm\n")
	for i, j := range shortest{
		fmt.Printf("from %d shortest way is %d\n", i, j)
	}
	fmt.Printf("Brute force\n")
	short:=brute("data.txt")
	for i, j:= range short{
		fmt.Printf("from %d shortest way is %d\n", i, j)
	}
}