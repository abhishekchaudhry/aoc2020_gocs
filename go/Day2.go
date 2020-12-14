package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file,err :=os.Open("input/day2.txt")
	if err != nil{
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
	var validpass =0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines := scanner.Text()
		inputs = append(inputs,lines)
		}

		for input:= range inputs{
			fmt.Println(inputs[input])

			fmt.Sscanf(inputs[input], "%d-%d %c: %s", &min, &max, &key, &pass)
			//fmt.Println("min value is-->",min)
			//fmt.Println("max value is-->",max)
			//fmt.Println("passkey is -->",string(key))
			//fmt.Println("password is-->",string(pass))
			count =0
			for p:= range pass{
				//fmt.Println(string(pass[p]))
				if string(key) == string(pass[p]){
					count++
				}
				//fmt.Println("",count)
			}

			if count<=max && count>=min{
				validpass++
			}


		//line = strings.Split(scanner.Text(),":")
		//fmt.Println(line)
		//split_1 = strings.Split(line[0],"-")
		//split_2 = strings.Split(line[1],"")
		//split_3 = strings.Split(split_1[1],"")
	}

	fmt.Println("valid passwords count",validpass)

	//fmt.Println(line[0])

	//min := split_1[0]
	//max := split_3[0]
	//passkey := split_3[2]
	//password := split_2

	//fmt.Println("first portion",line[0])
	//fmt.Println("second portion",line[1])
	//fmt.Println("first half again break",split_1[0])
	//fmt.Println("first half second break",split_1[1])
	//fmt.Println("second half first break",split_2[0])
	//fmt.Println("second half second break",split_2[4])
	//fmt.Println("spllit3 half second break",split_3[0])
	//fmt.Println("spllit3 half second break",split_3[1])
	//fmt.Println("spllit3 half second break",split_3[2])



}
