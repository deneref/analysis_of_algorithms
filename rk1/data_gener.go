package main 
import (
	"fmt"
	"os";  "math/rand"
	)

func main(){
	entity := make([]string, 0)
	entity = append(entity,
    "Person",
    "Location",
    "Organization",
    "Time")
	
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	//defer file.Close()
	//amount:=0 
	for i:=0;i<100;i++{
		m := int64(i)
		rand.Seed(m)
		//str := []byte(fmt.Sprintf("%d %d %s\n", rand.Intn(10), rand.Intn(10), entity[rand.Intn(len(entity))]))
		//fmt.Println(str)
		str:= fmt.Sprintf("%d %d %s\n", rand.Intn(10), rand.Intn(10), entity[rand.Intn(len(entity))])
		fmt.Println(str)
		file.WriteString(str)
	}
	defer file.Close()
	return
}