// package main

// import "fmt"

// func main() {
// 	a := 2
// 	b := 9
// 	for i := a; i < b; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	a, b := 1, 16

// 	for i := a; i <= b; i++ {
// 		j := math.Sqrt(float64(i))
// 		if float64(i) == j*j {

// 			fmt.Println(i)
// 		}
// 	}
// }
//непонятки с типом данных, математический пакет не извлекает корни из только целых чисел, из-за этого
// получаем некоторые лишние цисла при вычисления. Нужно подробнее разобраться с типами данных

/*Входные данные
Вводится натуральное число x.

Выходные данные
Выведите сумму цифр числа x.*/

// package main

// import "fmt"

// func main() {
// 	a := 9
// 	j := 0
// 	for i := 0; i <= a; i++ {
// 		j += i
// 	}
// 	fmt.Println(j)

// }

package main

import "fmt"

func main() {
	x := 6
	res := 0
	for i := 0; i < x; i++ {
		res = i
		if i%2 == 0 &&  {
			fmt.Println(res)
		}
	}

}
