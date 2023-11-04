package main

import (
	"errors"
	"fmt"
)
var msg string

func init(){
	msg="Вызовется перед мэин"
}



func main() {
	fmt.Println(msg)
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
	message, _ := checkTheClub(80)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(message)
	fmt.Println(calendar("ПТ"))

	func ()  {
		fmt.Println("анонисная функция")
	}()

	text := "простой текст"
	text_2 := "второй простой текст"
	setText(&text, text_2)
	fmt.Println(text)
	fmt.Println(text_2)
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

func calendar(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "ПН":
		return "Понедельник день тяжелый", nil
	case "ВТ":
		return "Понедельник день тяжелый", nil
	case "СР":
		return "Понедельник день тяжелый", nil
	case "ЧТ":
		return "Маленкьая пятница", nil
	case "ПТ":
		return "Ура пятница", nil
	case "СБ":
		return "Отдыхаем по полной", nil
	case "ВС":
		return "Просто отдыхаем", nil
	default:
		return "Что за день такой? ", errors.New("no day of week!")
	}
}

func setText(text *string, text_2 string){
	*text += " из функции добавлен текст"
	text_2 += " а это без * из функции"
	fmt.Println(*text)
	fmt.Println(text_2)
}