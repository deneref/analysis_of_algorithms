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
func read_file(name_file string) ([]int, []int) {
	file, err := os.Open(name_file)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	ids_person := make([]int, 0)
	ids_entity:= make([]int, 0)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF{
			break}
		res := strings.Split(str, " ")
		article_id, err := strconv.Atoi(res[1]); person_id, err:= strconv.Atoi(res[0])
		ids_person = append(ids_person, person_id)
		ids_entity = append(ids_entity, article_id)
	}
	return ids_person, ids_entity
}
func get_articles_hash(ids_person, ids_entity []int, amount int) [][]int{
	articles := make([][]int, amount, amount)
	for i:=0; i<amount; i++{
		articles[i] = make([]int, 0)
	}
	for i:=0; i<amount; i++{
		for j:=0; j<len(ids_entity); j++{
			if ids_entity[j] == i{
				fl := true
				for k:=0; k < len(articles[i]); k++{
					if ids_person[j] == articles[i][k]{
						fl=false}
					}
				if fl{
					articles[i] = append(articles[i], ids_person[j])
				}
			}		
		}
	}
	return articles
} 

func get_articles_brute(ids_person, ids_entity []int, amount int) [][]int{
	articles := make([][]int, 0)
	for i:=0; i<len(ids_person); i++{
		tmp := make([]int, 0)
		tmp = append(tmp, 0)
		for j:=0; j<len(ids_person); j++{
			if ids_entity[j] == ids_entity[j]{
				tmp = append(tmp, j)
			}		
		}
		articles = append(articles, tmp)
	}
	return articles
}
func print_articles(articles [][]int){
	for i:=0; i<len(articles);i++{
		fmt.Print(i, " ")
		for j:=0; j<len(articles[i]);j++{
			fmt.Printf("%d ", articles[i][j])
		}
		fmt.Println()
	}
}
func estimate_time(ids_person, ids_entity []int, amount int){
	tmp := make([]time.Duration, 20,20)
	for i:=0; i<20; i++{
		start := time.Now()
		articles:=get_articles_hash(ids_person, ids_entity, amount)
		end:=time.Now()
		tmp[i] = end.Sub(start)
		
		_ = articles
	}
	start := time.Now()
	articles:=get_articles_brute(ids_person, ids_entity, amount)
	end:=time.Now()
	sum1 := end.Sub(start)
	_ = articles
	var sum time.Duration
	for i:=0; i<len(tmp);i++{
		sum += tmp[i]
	}
	fmt.Println("hash ", sum/20)
	fmt.Println("brute ", sum1)
}
func main(){
	amount := 10
	ids_person, ids_entity:=read_file("data.txt")
	estimate_time(ids_person, ids_entity, amount)
	return
}