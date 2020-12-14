package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var inputs []string
	//var split_1[]string
	//var split_2[]string
	//var split_3[]string

	var min int
	var max int
	var key rune
	var pass string
	var count = 0
	//var common =0
	//var validpass =0
	var total = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		inputs = append(inputs, lines)
	}

	for input := range inputs {
		fmt.Println(inputs[input])
		min = 0
		max = 0
		fmt.Sscanf(inputs[input], "%d-%d %c: %s", &min, &max, &key, &pass)
		//fmt.Println("min value is-->",min)
		//fmt.Println("max value is-->",max)
		//fmt.Println("passkey is -->",string(key))
		//fmt.Println("password is-->",pass)
		count = 0
		//common=0

		//fmt.Println(string(pass[min-1]))
		//fmt.Println(string(pass[max-1]))
			//fmt.Println(string(pass[p]))
			if string(key) == string(pass[min-1]) || string(key) == string(pass[max-1]) {
				if string(pass[min-1]) != string(pass[max-1]) {
					count++
				}
				//fmt.Println("",count)
			}

			total = total + count

			//line = strings.Split(scanner.Text(),":")
			//fmt.Println(line)
			//split_1 = strings.Split(line[0],"-")
			//split_2 = strings.Split(line[1],"")
			//split_3 = strings.Split(split_1[1],"")

	}

	fmt.Println("valid passwords count", total)

}