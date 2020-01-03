package main
import (
	 "fmt"
	)

func main(){
	env := new_env("data.txt")
	
	for i:= 0; i< len(env.weight); i++{
		ant := env.new_ant(i)
		ant.ant_go()
		fmt.Println(i, " ", ant.get_distance())
	}
	
	
	/*
	for _, j := range env.pheromon{
		fmt.Print(j, "\n")
	}
	*/
}