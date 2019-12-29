package main
import (
	"fmt"
	_"math/rand"
	"time"
	)
	
type cake struct {
	num int
	dough bool
	topping bool
	decor bool
	started_dough time.Time
	finished_dough time.Time
	started_topping time.Time
	finished_topping time.Time
	started_decor time.Time
	finished_decor time.Time
	}
	
func conv(amount int, wait chan int) *queue{
	uno := make(chan *cake, 5)
	dos := make(chan *cake, 5)
	tres := make(chan *cake, 5)
	line := new_queue(amount) 
	first := func(){
		for{
			select{
				case a := <- uno:
				//fmt.Printf("Cake num %d started dough\n", a.num)
				a.dough = true
				
				a.started_dough = time.Now()
				took_dough := 200
				time.Sleep(time.Duration(took_dough) * time.Millisecond)
				
				a.finished_dough = time.Now()
				dos <- a
			}
		}
	}
	
	second:= func(){
	for{
		select{
			case a := <- dos:
				//fmt.Printf("Cake num %d started topping\n", a.num)
				a.topping = true
				
				a.started_topping = time.Now()
				took_topping := 200
				time.Sleep(time.Duration(took_topping) * time.Millisecond)
				
				a.finished_topping = time.Now()
				tres <- a
		}
	}
}
	
	third := func(){
	for{
		select{
			case a := <- tres:
			//fmt.Printf("Cake num %d started decor\n", a.num)
			a.decor = true
			
			a.started_decor = time.Now()
			took_decor := 200
			time.Sleep(time.Duration(took_decor) * time.Millisecond)
			
			a.finished_decor = time.Now()
			line.push(a)
			if (a.num == amount){
			 wait <- 0 }
			
		}
	}
}
	
	go first()
	go second()
	go third()
	for i:=0; i<=amount; i++{
		a := new(cake)
		a.num = i
		uno <- a
	}
	return line
}

func get_log(queue *queue){
	var first_waited time.Duration; var second_waited time.Duration; var third_waited time.Duration
	//first_waited = 0; second_waited = 0; third_waited = 0
	line := queue.waiting
	start := line[0].started_dough
	fmt.Printf("Starting time\n")
	//fmt.Printf(line[0])
	for i:=0;i<len(line);i++{
		if line[i] != nil{
			fmt.Println(i, line[i].started_dough.Sub(start),line[i].started_topping.Sub(start), line[i].started_decor.Sub(start))
		}}
	fmt.Printf("Finishing time\n")
	for i:=0;i<len(line);i++{
		if line[i] != nil{
			fmt.Println(i, line[i].finished_dough.Sub(start),line[i].finished_topping.Sub(start), line[i].finished_decor.Sub(start))
		}}
	fmt.Printf("Линии простаивали\n")
	for i:=0; i<len(line)-1;i++{
		first_waited+=line[i+1].started_dough.Sub(start)-line[i].finished_dough.Sub(start)
		second_waited+=line[i+1].started_topping.Sub(start)-line[i].finished_topping.Sub(start)
		third_waited+=line[i+1].started_decor.Sub(start)-line[i].finished_decor.Sub(start)
	}
	fmt.Println(first_waited, second_waited, third_waited)
}

func main(){
	amount := 20
	
	start := time.Now()
	wait := make(chan int)

	line:=conv(amount, wait)
	<-wait
	get_log(line)
	_ = line
	end := time.Now()
	fmt.Printf("To make %d cakes conv took ", amount)
	fmt.Println(end.Sub(start))
	
	start = time.Now()
	line_1 := linear(amount)
	end = time.Now()
	_ =  line_1;
	get_log(line_1)
	fmt.Printf("To make %d cakes linear conv took ", amount)
	fmt.Println(end.Sub(start))
	
	}
	
	