package main 
import (
	_"math/rand"
	"time"
	)

type queue struct{
	waiting [](*cake)
	last int
}
func new_queue(amount int) *queue{
	new_queue := new(queue)
	new_queue.waiting = make([](*cake), amount, amount)
	new_queue.last = -1
	return new_queue
}
func (this *queue) push(elem *cake){
	if this.last !=  len(this.waiting)-1{
		this.waiting[this.last+1] = elem
		this.last+=1
	}
}
func (this *queue) pop() *cake{
	elem := this.waiting[0]
	this.waiting = this.waiting[1:]
	this.last -= 1
	return elem
}
	
func first(a *cake, queue_for_topping *queue){
	a.dough = true
			
	a.started_dough = time.Now()
	took_dough := 200
	time.Sleep(time.Duration(took_dough) * time.Millisecond)
	a.finished_dough = time.Now()
	queue_for_topping.push(a)
	}

func second(a *cake, queue_for_decor *queue){
	a.topping = true
				
	a.started_topping = time.Now()
	took_topping := 200
	time.Sleep(time.Duration(took_topping) * time.Millisecond)
				
	a.finished_topping = time.Now()
	queue_for_decor.push(a)
}

func third (a *cake, finished *queue){
	a.decor = true
			
	a.started_decor = time.Now()
	took_decor := 200
	time.Sleep(time.Duration(took_decor) * time.Millisecond)
			
	a.finished_decor = time.Now()
	finished.push(a)
	}

func linear (amount int)*queue{
	queue_for_topping := new_queue(amount)
	queue_for_decor := new_queue(amount)
	finished := new_queue(amount)
	i:= 0
	for ;i!=-1;{
		a := new(cake)
		a.num = i
		first(a, queue_for_topping)
		if queue_for_topping.last >= 0{
			second(queue_for_topping.pop(), queue_for_decor)
		}
		if queue_for_decor.last >= 0{
			third(queue_for_decor.pop(), finished)
		}
		if finished.waiting[len(finished.waiting)-1] != nil{
				return finished}
		i+=1
	}
	return finished
}