package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func UpperRange(a int,b int) (int,int){
	mid := (float64(a) + float64(b))/2
	first := float64(b) - mid
	first1 := float64(a) + first
	last := b

	return int(math.Round(first1)),last
}

func LowerRange(a int,b int) (int,int){

	mid := (float64(a) + float64(b))/2
	//m:=math.Floor(float64(mid))
	first := a
	last := float64(b) - mid
	last1:= float64(first) + last

	return int(first),int(last1)
}

func main(){
	file,err := os.Open("input/day5.txt")
	if err!=nil{
		panic(err)
	}

	defer file.Close()

	var line []string
	var seats[]string
	var seats1[]int
	scanner :=bufio.NewScanner(file)
	for scanner.Scan(){
		lines :=scanner.Text()
		//fmt.Println(lines)

		line = append(line,lines)
	}

	fmt.Println(line[1])


	for i:=0;i<len(line);i++{
		U:=127
		L:=0
		CU:=7
		CL:=0
		line1 :=line[i]

		for x:=range line1{
			//fmt.Println(string(line1[x]))
			y:=string(line1[x])
			//fmt.Println(string(y))
			if y =="F" {
				r1,r2:= LowerRange(L,U)
				//fmt.Println(r1)
				//fmt.Println(r2)
				L=r1
				U=r2
			}

			if y =="B" {
				r1,r2:= UpperRange(L,U)
				//fmt.Println(r1)
				//fmt.Println(r2)
				L=r1
				U=r2
			}

			if y == "R"{
				r1,r2:= UpperRange(CL,CU)
				CL=r1
				CU=r2
			}

			if y == "L"{
				r1,r2:= LowerRange(CL,CU)
				CL=r1
				CU=r2
			}

		}
		seat_id:= (L * 8) + CL
		seats1=append(seats1,seat_id)
		seats=append(seats,string(seat_id))
				//fmt.Println("row id is -->",L)
		//fmt.Println("column id is-->",CU)
		//fmt.Println("seat id is -->",seat_id)
	}

		for i:=1;i <len(seats1);i++{
			if seats1[0] < seats1[i] {seats1[0] = seats1[i]}
		}
		fmt.Println("the largest seat id is:",seats1[0])

		//solution for part2
	for i := 0; i < 128*8; i++ {
		if contains(seats, string(i+1)) && contains(seats, string(i-1)) && !contains(seats, string(i)) {
			fmt.Println("the seat ID is ",i)
		}
	}

}
