package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	maxLines := 10
	lineCount := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		if lineCount >= maxLines {
			break
		}
		fmt.Println(scanner.Text())
		lineCount++
	}

	if err := scanner.Err(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

//	args := os.Args
//	if len(args) < 2 {
//		fmt.Println("no arguments provided. Usage: myhead <file_name>")
//		return
//	}
//
//	if len(args) > 2{
//		fmt.Println("too many arguments. Usage: myhead <file_name>")
//		return
//	}
//
//	file_name := args[1]
//	file, err := os.Open(file_name)
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//
//	data := make([]byte, 100)
//	count, err := file.Read(data)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Print(string(data[:count]))
}
