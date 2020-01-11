package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"strconv"
	"time"
	)
	
func init_table(amount int) [][]int{
	table := make([][]int, amount, amount)
	for i:=0;i<len(table);i++{
		table[i] = make([]int,0)
	}
	return table
}


func print_table(table [][]int){
	for i, _ := range table{
		fmt.Printf("\n %d ", i)
		for _, j := range table[i]{
			fmt.Print(j)
			fmt.Print(" ")
		}
	}
}

func read_file(name_file string, table [][]int, amount int){
	file, err := os.Open(name_file)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF{
			break}
		res := strings.Split(str, " ")
		article_id, err := strconv.Atoi(res[1]); person_id, err:= strconv.Atoi(res[0])
		fl := true
		for _, j:=range table[article_id]{
			if j == person_id{
				fl = false
				break}
		}
		if fl{
			table[article_id] = append(table[article_id], person_id)
		}
	}
}
		
func estimate_time(name_file string, amount int){
	tmp := make([]time.Duration, 20,20)
	table := init_table(amount)
	for i:=0; i<20; i++{
		start := time.Now()
		read_file(name_file,table, amount)
		end:=time.Now()
		tmp[i] = end.Sub(start)
	}
	var sum time.Duration
	for i:=0; i<len(tmp);i++{
		sum += tmp[i]
	}
	fmt.Print(sum/20)
}
	
func main(){
	amount := 10
	estimate_time("data.txt", amount)
	return
}