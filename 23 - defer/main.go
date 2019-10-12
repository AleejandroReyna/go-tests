package main

import "fmt"

func main() {
	checkName()
}

func checkName() {
	openConnection()
	defer closeConnection()

	showData()
	defer closeData()
}

func showData() {
	fmt.Println("show data")
}

func openConnection() {
	fmt.Println("Open connection")
}

func closeData() {
	fmt.Println("Close data.")
}

func closeConnection() {
	fmt.Println("Connection closed.")
}
