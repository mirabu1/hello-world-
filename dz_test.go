package main

import (
	"testing"
)


func TestTransfer(t *testing.T){
	cl1 := client{LastName: "IVAN", FirstName: "IVANOV", balans: 5}
	cl2 := client{LastName: "PETR", FirstName: "PETROV", balans: 5}
	Transfer(&cl1, &cl2, 5)
	if cl1.balans != 0 || cl2.balans != 10{
		t.Errorf("ошибка во время перевода")
	}
	Transfer(&cl1, &cl2, 5)
	if cl1.balans != 0 || cl2.balans != 10{
		t.Errorf("ошибка во время перевода")
	}
}
