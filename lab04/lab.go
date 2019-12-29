package main
import  _ "fmt"

func standart(mtr1, mtr2 *mtr) *mtr{
	if len(mtr1.buff) != len(mtr2.buff[0]){
		panic("Wrong size of matrix")
	}
	row1 := mtr1.rows; col1 := mtr1.cols
	col2 := mtr2.cols
	
	res := new_mtr(row1, col2)
	for i:=0;i<row1;i++{
		for j:=0;j<col1;j++{
			for k:=0;k<col2;k++{
				res.buff[i][j] += mtr1.buff[i][k]*mtr2.buff[k][j]
				}
		}
	}
	return res
}

func winograd(mtr1, mtr2 *mtr) *mtr{
	if mtr2.rows != mtr1.cols{
		panic("Wrong size of matrix")
	}
	row1 := mtr1.rows; col1 := mtr1.cols
	row2:= mtr2.rows; col2 := mtr2.cols
	//d := col1 / 2
	row_factor := make([]int, row1, row1)
	col_factor := make([]int, col2, col2)
	
	for i:=0; i<row1;i++{
		for j:=0;j<col1 / 2;j++{
			row_factor[i] += mtr1.buff[i][2*j] * mtr1.buff[i][2*j+1]
	}}
	for i:=0; i<col2;i++{
		for j:=0;j<row2 / 2;j++{
			col_factor[i] += mtr2.buff[2*j][i]*mtr2.buff[2*j+1][i]
	}}
	
	answer := new_mtr(row1, col2)
	
	for i:=0;i<row1;i++{
		for j:=0;j<col2;j++{
			answer.buff[i][j]+= - row_factor[i] - col_factor[j]
			for k:=0;k<col1 / 2;k++{
				answer.buff[i][j] += ((mtr1.buff[i][2 * k] + mtr2.buff[2 * k + 1][j]) * (mtr1.buff[i][2 * k + 1] + mtr2.buff[2 * k][j]))
	}}}
	
	if (row2 % 2) != 0{
		for i:=0;i<row1;i++{
			for j:=0;j<col2;j++{
				answer.buff[i][j] += mtr1.buff[i][col1 - 1] * mtr2.buff[col1 - 1][j]
	}}}
	return answer
}

func winograd_parallel(mtr1, mtr2 *mtr, threads int) *mtr{
	if mtr2.rows != mtr1.cols{
		panic("Wrong size of matrix")
	}
	row1 := mtr1.rows; col1 := mtr1.cols
	row2:= mtr2.rows; col2 := mtr2.cols
	
	row_factor := make([]int, row1, row1)
	col_factor := make([]int, col2, col2)
	
	for i:=0; i<row1;i++{
		for j:=0;j<col1 / 2;j++{
			row_factor[i] += mtr1.buff[i][2*j] * mtr1.buff[i][2*j+1]
	}}
	for i:=0; i<col2;i++{
		for j:=0;j<row2 / 2;j++{
			col_factor[i] += mtr2.buff[2*j][i]*mtr2.buff[2*j+1][i]
	}}
	answer := new_mtr(row1, col2)
	in := make(chan int); quit := make(chan bool)
	mult := func(){
		for{
			select {
			case i := <- in:
				//sums := make([]int, mtr2.cols)
					for j:=0;j<col2;j++{
					answer.buff[i][j]+= - row_factor[i] - col_factor[j]
							for k:=0;k<col1 / 2;k++{
				answer.buff[i][j] += ((mtr1.buff[i][2 * k] + mtr2.buff[2 * k + 1][j]) * (mtr1.buff[i][2 * k + 1] + mtr2.buff[2 * k][j]))
	}}
			case <- quit:
				return
		}
		}
	}
	for i:=0;i<threads;i++{
		go mult()
	}
	for i:=0;i<mtr1.rows;i++{
		in <- i
	}
	for i:= 0 ; i<threads; i++{
		quit<-true
	}
	return answer
}
func main(){
	/*
	m := new_mtr(4,4); n:= new_mtr(4,4)
	m.fill(); n.fill()
	c:=winograd_parallel(m,n, 2)
	d:=winograd(m,n)
	c.print()
	fmt.Print("winograd\n")
	d.print()
	*/
	for i:=100;i<=1000;i+=100{
		estimate_threads(i)
	}
	}