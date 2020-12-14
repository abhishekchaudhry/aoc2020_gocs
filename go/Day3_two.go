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
	prev_count:=1
	incr:=1
	position:=0
	ranges := [5]int{1,3,5,7,1}
	for x:= range ranges {
		position = ranges[x]
		if x==4{incr=2}
		fmt.Println("Checking with buffer range-->", position)
		pos:=0
		trees := 0
		for i := 0; i < len(inputs); i=i+incr {
			each_line := inputs[i]
			len := len(each_line)

			if each_line[pos] == 35 {trees++}
			pos = pos + position
			if pos >= len {pos = pos - len}

		}

		fmt.Println("Tree count -->",trees)
		if trees==0{trees=1} //to avoid making the multiplication 0 in case of no trees found
		prev_count=prev_count*trees

	}


	fmt.Println("Tree count multiplication -->",prev_count)
}