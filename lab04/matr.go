package main
import (
	"fmt"
	"math/rand"
	"time"
		)
		
type mtr struct{
	buff [][]int
	rows int
	cols int
}

func new_mtr(row, col int) *mtr{
	if row <= 0 || col <= 0 {
	panic("Size of matrix can't be 0 or less")
	}
	m := new(mtr)
	m.buff = make([][]int, row, row) 
	m.rows, m.cols = row, col
	for i := 0; i < row; i++ {
		m.buff[i] = make([]int, col, col)
	}
	return m
}

func (m *mtr)fill(){
	for i:= 0; i<m.rows; i++{
		for j:=0; j<m.cols; j++{
			m.buff[i][j] = rand.Intn(10)
			}
		}
}
func (m *mtr) print(){
	for i:= 0; i<m.rows;i++{
		for j:=0; j<m.cols;j++{
			fmt.Print(m.buff[i][j], " ")
		}
		fmt.Print("\n")
}}

func estimate_time(size int){
	m := new_mtr(size,size); n := new_mtr(size,size)
	m.fill();n.fill();
	fmt.Print(size, " ")
	
	start := time.Now()
	res := standart(n,m); _ = res;
	t := time.Now()
	fmt.Print("std = ", t.Sub(start), " ")
	
	start = time.Now()
	res = winograd(n,m)
	t = time.Now()
	fmt.Print("win = ", t.Sub(start), " ")
	
	start = time.Now()
	res = winograd_parallel(n,m,2)
	t = time.Now()
	fmt.Print("win_parallel = ", t.Sub(start), "\n")
	
}

func estimate_threads(size int){
	fmt.Println("For size of ", size)
	m := new_mtr(size,size); n := new_mtr(size,size)
	m.fill();n.fill();
	
	start := time.Now()
	res := winograd(n,m); _ = res;
	t := time.Now()
	fmt.Printf("For winograd time is ")
	fmt.Println(t.Sub(start))
	
	start = time.Now()
	res = winograd_parallel(n,m,1); _ = res;
	t = time.Now()
	fmt.Printf("With 1 thread running time is ")
	fmt.Println(t.Sub(start))
	
	for i:=2;i<10;i+=2{
		start := time.Now()
		res := winograd_parallel(n,m,i); _ = res;
		t := time.Now()
		fmt.Printf("With %d threads running time is ",i)
		fmt.Println(t.Sub(start))
	}
}
/*
func main() {
    m := new_mtr(5, 5)
	m.fill()
	m.print()
}
*/