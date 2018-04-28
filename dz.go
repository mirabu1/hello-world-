package main

import (
	"fmt"
)

type client struct {
	LastName, FirstName string
	balans              float64
}

func transfer(cl1, cl2 *client, x float64) {
	if cl1 == nil || cl2 == nil{
		fmt.Println("некорректные данные пользователей")
		return
	}
	if cl1.balans < x {
		fmt.Println("Залезать в долги плохо!")
	} else {
		cl1.balans -= x
		cl2.balans += x
	}
}

func main() {
	fmt.Println("Введите имя, фамилию и баланс отправителя: ")
	var FirstName, LastName string
	var balans float64
	fmt.Scan(&FirstName, &LastName, &balans)
	otpr := client{balans: balans, FirstName: FirstName, LastName: LastName}
	fmt.Println("Введите имя, фамилию и баланс получателя : ")
	fmt.Scan(&FirstName, &LastName, &balans)
	pol := client{balans: balans, FirstName: FirstName, LastName: LastName}
	var suma_perevoda float64
	fmt.Println("введите сумму перевода")
	fmt.Scan(&suma_perevoda)
	transfer(&otpr, &pol, suma_perevoda)
	fmt.Println(otpr.balans, pol.balans)
}
