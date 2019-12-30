package main
import (
	"fmt"
	"os"
	"io"
	"container/list"
	"bufio"
	"strings"
	"strconv"
	)
	
func init_table(amount int) [](*list.List){
	table := make([](*list.List), amount, amount)
	for i:=0;i<len(table);i++{
		table[i] = list.New()
	}
	return table
}


func print_table(table [](*list.List)){
	for i:=0; i<len(table);i++{
		fmt.Printf("\n %d ", i)
		for tmp:=table[i].Front(); tmp != nil; tmp = tmp.Next(){
			fmt.Print(tmp.Value)
			fmt.Print(" ")
		}
	}
}

func read_file(name_file string, table [](*list.List), amount int){
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
		for tmp:=table[article_id].Front(); tmp != nil; tmp = tmp.Next(){
			if tmp.Value == person_id{
				fl = false
				break}
		}
		if fl{
			table[article_id].PushBack(person_id)
		}
	}
}
func perebor(name_file string, amount int){
	file, err := os.Open(name_file)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
	
func main(){
	amount := 10
	table := init_table(amount)
	read_file("data.txt",table, amount)
	print_table(table)
	return
}