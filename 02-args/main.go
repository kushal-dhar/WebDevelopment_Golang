package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := "Kushal Dhar"

	tpl := `
  <DOCTYPE html>
  <html lang="en"
  <head>
  <meta charset="UTF-8"
  <title>Hello World!</title>
  </head>
  <body>
  <h1>` + name + `</h1>
  </body>
  </html>`

	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println("error in opening file", err)
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(tpl))
}
