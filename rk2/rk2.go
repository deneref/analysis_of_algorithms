package main
import (
	"fmt"
	"os";
	"io"
	"bufio"
	"regexp"
	)

func get_fors(name_file string) [][]byte{
	answer:=make([][]byte, 0)
	//for_pattern:="^(for).*\\{+"
	for_pattern:="testGuild=.+&"
	get_for, _ := regexp.Compile(for_pattern)
	file, err := os.Open(name_file)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		//str := []byte(str1)
		//fmt.Println(str)
		if err == io.EOF{
			break}
		answer = append(answer, get_for.Find([]byte(str)))
	}
	//fmt.Println(answer)
	return answer
}
func print_answer(answer [][]byte){
	for i:= 0; i< len(answer);i++{
		fmt.Println(string(answer[i]))
	}
}

func main(){
	answer := get_fors("test.txt")
	print_answer(answer)
}