package main
import (
	"fmt"
	 "strings"
	"bufio"
	"os"
	_ "strconv"
	"unicode/utf8"
)

func std(str, sub string) int{
	if len(sub) > len(str) || len(sub)== 0 || len(str) == 0{
		return -1
	}
	
	flag := true
	for i:=0;i<len(str);i++{
		flag = true
		if str[i] == sub[0]{
			for j:=0;j<len(sub);j++{
				if str[i+j] != sub[j]{
					flag = false
				}
				if flag{
					return i
				}
			}
		}
	}
	return -1
}
//get prethics for Knuth-Morris-Pratt algorithm
func get_preph(str string) []int{
	res := make([]int, len(str))
	for i:= 1; i<len(str);i++{
		j := res[i-1]
		for (j>0 && str[i] != str[j]){
			j = res[j-1]
		}
		if (str[i] == str[j]){
			j++
		}
		res[i] = j
	}
	return res
}

func KMP(str, sub string) int{
	if len(sub) > len(str) || len(sub)== 0 || len(str) == 0{
		return -1
	}
	
	tmp := []string{sub, str}
	str = strings.Join(tmp, "@")
	pr := get_preph(str)
	len_sub := len(sub)
	for i:=len_sub+1; i<len(str);i++{
		if pr[i] == len_sub{
			return i-2*len_sub
		}
	}
	return -1
}
//get table for Boyer-Moore
func get_table(substring string) map[rune]int {
	length := utf8.RuneCountInString(substring)
	runes := []rune(substring)

	table := make(map[rune]int)

	for i := 0; i < length; i++ {
		j := runes[i]
		table[j] = length - i - 1
	}

	return table
}

func BM(str, sub string) int{
	if len(sub) > len(str) || len(sub)== 0 || len(str) == 0{
		return -1
	}
	table := get_table(sub)
	strrunes := []rune(str)
	subrunes := []rune(sub)
	str_l := utf8.RuneCountInString(str)
	sub_l := utf8.RuneCountInString(sub)
	i:=sub_l-1
	j,k:=i,i
	for (j>=0 && i<=str_l-1){
		j = sub_l-1
		k = i
		for(j>=0 && (table[strrunes[k]] == table[subrunes[j]])){
			k--
			j--
		}
		if _,ok := table[strrunes[i]]; !ok{
			i+= sub_l
		}else{
		i+=table[strrunes[i]]
		}
	}
	if k>= str_l - sub_l{
		return -1
	} else {
		return k+1
	}
}

func call_func(f func(str, sub string) int){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку ")
	str, _ := reader.ReadString('\n')
	fmt.Println("Введите подстроку ")
	substr, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\r\n")
	
	substr = strings.TrimSuffix(substr, "\r\n")
	answer:=f(str, substr)
	if answer != -1{
		fmt.Println("Строка начинается с индекса ", answer)
	}else{
		fmt.Println("Подстрока не найдена :с")
	}
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("1. Стандартный алгоритм\n2. Бойер-Мур\n3. Кнут-Моррис-Пратт\n4. Лог времени\n0. Выход")
		answer, _ := reader.ReadString('\n')
		/*
		res, err := strconv.Atoi(answer)
			if err != nil{
				fmt.Println(err)
			}
		*/
		switch answer{
			case "1\r\n":
				call_func(std)
			case "2\r\n":
				call_func(BM)
			case "3\r\n":
				call_func(KMP)
			case "4\r\n":
				get_log()
			case "0\r\n":
				os.Exit(0)
		}
	}
}