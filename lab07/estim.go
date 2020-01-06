package main
import (
	"time"
	"fmt"
	"math/rand"
	)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func get_time(size, size_sub int){
	str:= RandStringRunes(size)
	substr :=RandStringRunes(size_sub)
	fmt.Printf("Для обработки случайной строки размера %d и подстроки размера %d\n", size, size_sub)
	start:=time.Now()
	res := std(str, substr);
	end:=time.Now()
	fmt.Println("Стандартный алгоритм ", end.Sub(start))
	
	start=time.Now()
	res = std(str, substr);
	end=time.Now()
	fmt.Println("Алгорит Бойера-Мура ", end.Sub(start))
	
	start=time.Now()
	res = std(str, substr);
	end=time.Now()
	fmt.Println("Алгоритм КМП ", end.Sub(start))
	_=res
}

func get_log(){
	for i:=10000; i<100000;i+=10000{
		get_time(i, 10)
	}
}