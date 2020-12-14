package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var inputs []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines := scanner.Text()
		inputs = append(inputs,lines)
		}

	//for input := range inputs {
	//	fmt.Println(inputs[input])
	//}
	trees := 0
	pos:=0
	for i:=0;i<len(inputs);i++{
		each_line:=inputs[i]
		//fmt.Println(each_line)
		len :=len(each_line)
		//fmt.Println(len)
			// ascii code for # is 35
			if each_line[pos] == 35 {
				trees++
			}
			pos=pos+3
			if pos >=len {
				pos = pos - len
			}
			//fmt.Println(each_line[29])
		}
		//fmt.Println(inputs[i])

	fmt.Println("Tree count-->",trees)
}