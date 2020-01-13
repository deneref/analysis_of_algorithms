package main
import (
	"fmt"
	"os";
	"io"
	"bufio"
	"regexp"
	)
type avt struct {
	state int
	trans map[byte]int
}	

func new_avt() *avt{
	avt := new(avt)
	avt.state  = 0
	avt.trans = make(map[byte]int)
	avt.trans[byte(' ')] = -1
	avt.trans[byte('f')] = 1
	avt.trans[byte('o')] = 2
	avt.trans[byte('r')] = 3
	avt.trans[byte('}')] = 2
	return avt
}
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
func avt* change_state(char byte){
	
}
func get_fors_avt(name_file string) []string{
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
		str = []byte(str)
		for i, j := range str{
			//todo
		}
	return answer
}

func main(){
	answer := get_fors("lab.txt")
	print_answer(answer)
}