package main
import (
	"fmt"
	"time"
	"os"
	)
func get_time_log_param(){
	ants :=make([]time.Duration, 0)
	for i:=1; i<20;i++{
		_ = os.Remove("data.txt")
		gener_file(10)
		env := new_env("data.txt")
		env.p = 1.0/float64(i)
		start:=time.Now()
		shortest := start_ants(env, 100); 
		end:=time.Now()
		ants = append(ants, end.Sub(start))
		
		_ = shortest
	}
	
	fmt.Println("Ants algotithm time")
	for j:=0;j<len(ants);j++{
		fmt.Println(j, ants[j])
	}
}

func get_time_log(){
	ants :=make([]time.Duration, 0)
	brutes:=make([]time.Duration, 0)
	for i:=2; i<11;i++{
		_ = os.Remove("data.txt")
		gener_file(i)
		env := new_env("data.txt")
		start:=time.Now()
		shortest := start_ants(env, 100); //fmt.Print(shortest, " ")
		end:=time.Now()
		ants = append(ants, end.Sub(start))
		
		start = time.Now()
		shortest = brute("data.txt"); //fmt.Print(shortest, "\n")
		end=time.Now()
		brutes = append(brutes, end.Sub(start))
		_ = shortest
	}
	
	fmt.Println("Ants algotithm time")
	for j:=0;j<len(ants);j++{
		fmt.Println(j+2, ants[j])
	}
	fmt.Println("Brute force time")
	for j:=0;j<len(ants);j++{
		fmt.Println(j+2, brutes[j])
	}
	
}

func main(){
	get_time_log()
}