package main

import (
	"fmt"
	"os";
	"io"
	"bufio";
	"strings"
	)

type avt struct {
	state int
	trans map[key]int
}	

type key struct{
	symb byte
	state int
}

func new_avt() *avt{
	avt := new(avt)
	avt.state  = 0
	avt.trans = make(map[key]int)
	avt.trans[key{symb: byte('f'), state: 0}] = 1
	avt.trans[key{symb: byte('f'), state: 1}] = 0
	avt.trans[key{symb: byte('f'), state: 2}] = 0
	avt.trans[key{symb: byte('f'), state: 3}] = 3
	avt.trans[key{symb: byte('f'), state: 4}] = 0

	avt.trans[key{symb: byte('o'), state: 0}] = 0
	avt.trans[key{symb: byte('o'), state: 1}] = 2
	avt.trans[key{symb: byte('o'), state: 2}] = 0
	avt.trans[key{symb: byte('o'), state: 3}] = 3
	avt.trans[key{symb: byte('o'), state: 4}] = 0
	
	avt.trans[key{symb: byte('r'), state: 0}] = 0
	avt.trans[key{symb: byte('r'), state: 1}] = 0
	avt.trans[key{symb: byte('r'), state: 2}] = 3
	avt.trans[key{symb: byte('r'), state: 3}] = 3
	avt.trans[key{symb: byte('r'), state: 4}] = 0
	
	avt.trans[key{symb: byte('{'), state: 0}] = 0
	avt.trans[key{symb: byte('{'), state: 1}] = 0
	avt.trans[key{symb: byte('{'), state: 2}] = 0
	avt.trans[key{symb: byte('{'), state: 3}] = 4
	avt.trans[key{symb: byte('{'), state: 4}] = 0
	
	avt.trans[key{symb: byte(' '), state: 0}] = 0
	avt.trans[key{symb: byte(' '), state: 1}] = -1
	avt.trans[key{symb: byte(' '), state: 2}] = -1
	avt.trans[key{symb: byte(' '), state: 3}] = 3
	avt.trans[key{symb: byte(' '), state: 4}] = 4
	
	avt.trans[key{symb: byte('	'), state: 0}] = 0
	avt.trans[key{symb: byte('	'), state: 1}] = -1
	avt.trans[key{symb: byte('	'), state: 2}] = -1
	avt.trans[key{symb: byte('	'), state: 3}] = 3
	avt.trans[key{symb: byte('	'), state: 4}] = 4
	
	return avt
}

func (avt *avt) change_state(char byte){
	key := key{char, avt.state}
	if st, ok := avt.trans[key]; ok{
		avt.state = st
	}else{
		if avt.state == 3{
			if char ==  byte('\n'){
				avt.state = 0
			}else{
				avt.state = 3
			}
		}else if avt.state == 0 || avt.state == 1 || avt.state == 2{
			avt.state = -1 //запрещаем все символы кроме пробела
		}else{
			avt.state = 0
		}
		}
}

func get_fors(name_file string) []string{
	answer:=make([]string, 0)
	file, err := os.Open(name_file)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\r\n")
		avt := new_avt()
		if err == io.EOF{
			break}
		//str = []byte(str)
		for _, j := range str{
			avt.change_state(byte(j))
			if avt.state == -1{
				break}
		}
		if avt.state == 4{
				answer = append(answer, str)
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