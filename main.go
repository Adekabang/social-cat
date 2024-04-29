package main

import "github.com/Adekabang/social-cat/app"

func main() {
	var a app.App
	a.CreateConnection()
	a.Routes()
	a.Run()
}
