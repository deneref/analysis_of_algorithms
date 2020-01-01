package main
import (
	"fmt"
	"os";
	"io"
	"bufio"
	"regexp"
	)
	
func get_fors(name_file string) []string{
	answer:=make([]string, 0)
	for_pattern:="(for).*\\{+"
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
		if err == io.EOF{
			break}
		if (get_for.MatchString(str)){
			answer = append(answer, get_for.FindString(str))
		}
	}
	return answer
}
func print_answer(answer []string){
	for i:= 0; i< len(answer);i++{
		fmt.Println(answer[i])
	}
}

func main(){
	answer := get_fors("lab.txt")
	print_answer(answer)
}