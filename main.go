package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	// var message string = "Лалала" 	//строковое число
	// var number int = 45 			//целое число
	// var halfnumber float32 = 45.4 	//долюное число
	// var b bool = true 				//логическое

	// a, b, c := 1, 2, 3 // множественное присвоение

	// fmt.Println(a, b, c)
	c := summ(5, 6)
	mes := fmt.Sprintf("Сумаа 5 и 6 равна %d", c ) // чтобы объединить в строку число и строку
	fmt.Println(mes)
	fmt.Println("Заглушка")
	fmt.Println("   ")
	message, err := checkTheClub(80)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	
}

func summ(num_1 int, num_2 int) int {
	sum := num_1 + num_2
	return sum
}

func checkTheClub(age int)(string, error) {
	if (age >= 18 && age < 45) {
		return "Входи",  nil
	} else if age >= 45  && age < 65{
		return "Входи",  nil
	} else if age >= 65{
		return "Вход запрещен",  errors.New("too old")
	}
	return "Вход запрещен", errors.New("too young")
}