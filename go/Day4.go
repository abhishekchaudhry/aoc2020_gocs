package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){

	fields :=make(map[string]string)

	file, err := os.Open("input/day4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()


	num_pass :=1
	count:=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines := scanner.Text()
		//fmt.Println(lines)

		if len(lines) != 0{
			split :=strings.Fields(lines)
			for x:=range split{
				line := strings.Split(split[x],":")
				fields[line[0]] = line[1]
			}

			fmt.Println(split)

		}else {
			fmt.Println(len(fields))
			if len(fields)==8{
				count++
			} else if len(fields) ==7 {
				for key,_:=range fields{
					if key!="cid"{
						count++
					}
				}
				count = count -len(fields) + 1 //need to subtract the no. of keys
			}

			for k:=range fields{
				delete(fields,k)
			}
			num_pass++
			continue
			}

	}
	//below part is again added since scanner.scan comes to this point without evaluating the last field
	//better solution would be to store the full input is a map and then check it.
	if len(fields)==8{
		count++
	} else if len(fields) ==7 {
		for key,_:=range fields{
			if key!="cid"{
				count++
			}
		}
		count = count -len(fields) + 1
	}
	fmt.Println("Number of passwords in file --------->",num_pass)
	fmt.Println("Number of passwords which are valid-->",count)
}