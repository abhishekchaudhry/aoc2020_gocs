package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file,err := os.Open("input/Day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		temp,err := strconv.Atoi(line)
		if err !=nil{
			panic(err)
		}
		input = append(input,temp)
	}

	var sum = 2020
	for i:=0;i<len(input);i++{
		for j:=i+1;j<len(input);j++{
			for k:=j+1;k<len(input);k++{
				if sum == input[i] + input[j] + input[k]{
					x:=input[i]
					y:=input[j]
					z:=input[k]
					fmt.Printf("the first number is %d\n", x)
					fmt.Printf("the second number is %d\n", y)
					fmt.Printf("the third number is %d\n", z)
					result := x * y * z
					fmt.Printf("the output is %d\n", result)
				} else {continue}

			}

		}
	}
}

