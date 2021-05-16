package main

import "fmt"

type People struct {
	name string
	age int
	live bool
	website string
	motto string
	girlfriend *People
}

func main() {
	TYD := People{"谭尧丹",21,true,"https://tanyaodan.com","You only live once",nil}
	printPeople(TYD)
	TYD.girlfriend = nil
	var LWK People
	printPeople(LWK)
}

func newPeople() People {
	return People{live: true, girlfriend: nil}
}

func printPeople(people People) {
	fmt.Printf("姓名：%s\n", people.name)
	fmt.Printf("年龄：%d\n", people.age)
	fmt.Printf("网址：%s\n", people.website)
	fmt.Printf("格言：%s\n", people.motto)
}