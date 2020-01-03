package main
import (
	"fmt"
	"os";
	"io"
	"bufio"
	"strings"
	"strconv"
	"math"
	"math/rand"
	"time"
	)
type env struct{
	weight [][]int
	pheromon[][]float64
	alpha float64 //приоретет пути
	betta float64 //приоретет феромона
	q float64 // переносимый муравьем феромона
	p float64 // коэффициент испарения феромона
	
}
type ant struct{
	env *env
	visited [][]int
	been_to [][]bool
	position int
}

func (e *env) new_ant(pos int) (*ant){
	ant := new(ant)
	ant.env = e
	ant.visited = make([][]int, len(e.weight))
	
	for i:=0; i< len(e.weight);i++{
		ant.visited[i] = make([]int, len(e.weight))
		for j:=0; j<len(e.weight[i]);j++{
			ant.visited[i][j] = e.weight[i][j]
		}
	}
	ant.position = pos
	ant.been_to = make([][]bool, len(e.weight))
	
	for i, _:= range ant.been_to{
		ant.been_to[i] = make([]bool, len(e.weight))
	}
	return ant
}

func new_env(file_name string) *env {
	enviroment := new(env)
	enviroment.weight = get_weights(file_name)
	enviroment.pheromon = make([][]float64, len(enviroment.weight), len(enviroment.weight))
	for i:=0;i<len(enviroment.pheromon); i++{
		enviroment.pheromon[i] = make([]float64, len(enviroment.weight[i]))
		for k, _ := range enviroment.pheromon[i]{
			enviroment.pheromon[i][k] = 0.5
		}
	}
	enviroment.alpha = 4.0
	enviroment.betta = 6.0
	enviroment.q = 1.0
	enviroment.p = 0.5
	return enviroment
}
func get_weights(name_file string) [][]int{
	answer:=make([][]int, 0)
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
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
		cur := strings.Split(str, " ")
		new_line := make([]int, 0)
		for _, i:= range cur{
			i, err := strconv.Atoi(i)
			if err != nil{
				fmt.Println(err)
			}
			new_line = append(new_line, i)
		}
		answer = append(answer, new_line)
	}
	return answer
}

func (ant *ant)count_probapility() []float64{
	res := make([]float64, 0);
	var d float64;
	var sum float64;
	for i, lenght := range ant.visited[ant.position]{
		if lenght != 0{
			d = math.Pow((float64(1)/float64(lenght)), ant.env.alpha) * math.Pow(ant.env.pheromon[ant.position][i], ant.env.betta)
			res = append(res, d)
			sum += d
		} else{
			res = append(res, 0)
		}
	}
	for _, lenght := range res{
		lenght = lenght / sum
	}
	return res
}

func choose_path(probab []float64) int{
	var sum float64
	for _, j := range probab{
		sum += j
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random_fl := r.Float64() * sum
	sum = 0
	for i , j := range probab{
		if random_fl > sum && random_fl<sum+j{
			return i
		} else{
			sum+=j
		}
	}
	return -1
}

func (ant *ant)renew_pheromon(){
	var del_t float64
	del_t = 0
	for k:=0;k<len(ant.env.pheromon);k++{
		for i, j := range ant.env.pheromon[k]{
			if ant.env.weight[k][i] != 0{
				if ant.been_to[k][i]{
					del_t = ant.env.q / float64(ant.env.weight[k][i])
				} else {
					del_t = 0
				}
				ant.env.pheromon[k][i] = (1-ant.env.p) * (float64(j) + del_t)
			}
		}
	}
}

func (ant *ant) go_path(path int){
	for i, _ := range ant.visited{
		ant.visited[i][ant.position] = 0
	}
	ant.been_to[ant.position][path] = true
	ant.position = path
}	

func (ant *ant) ant_go(){
	for{
	//ant.print_info()
		prob := ant.count_probapility()
		//fmt.Println("\nprob ", prob)
		choosen_path := choose_path(prob)
		//fmt.Println("chosen_path", choosen_path)
		if choosen_path == -1{
			break}
		ant.go_path(choosen_path)
		ant.renew_pheromon()
		
	}
}

func (ant *ant) print_info(){
	fmt.Println("Ant pos: ", ant.position)
	for i, _ := range ant.visited{
		fmt.Println()
		for _, k := range ant.visited[i]{
			fmt.Printf("%d ", k)
		}
	}
	fmt.Println("\ncurrent enviroment pheromons")
		for i, _ := range ant.env.pheromon{
		fmt.Println()
		for _, k := range ant.env.pheromon[i]{
			fmt.Printf("%f ", k)
		}
	}
}
			
func (ant *ant) get_distance()int{
	var distance int
	//fmt.Println(ant.env.weight)
	//fmt.Println(ant.visited)
	for i, j:= range ant.been_to{
		for k, z := range j{
			if z{
				//fmt.Println(i," ", k," ", ant.env.weight[i][k])
				distance += ant.env.weight[i][k]
			}
		}
	}
	return distance
}
			
			
			