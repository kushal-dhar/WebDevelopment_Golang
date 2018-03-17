package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println("Error in opening file ", err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("Fatal Error")
	}
}
