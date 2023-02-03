// Определяем високосный год по входным данным

// package main

// import "fmt"

// func main() {
// 	year := 2024

// 	if year%400 == 0 && year%100 != 0 || year%4 == 0 {
// 		fmt.Printf("YES")
// 	} else {
// 		fmt.Printf("NO")
// 	}
// }

// Находим знак числа
// package main

// import "fmt"

// func main() {
// 	x := 0
// 	if x > 0 {
// 		fmt.Println(1)
// 	} else if x < 0 {
// 		fmt.Println(-1)
// 	} else {
// 		fmt.Println(0)
// 	}
// }

// максимум из трех

// package main

// import "fmt"

// func main() {
// 	x, y, z := 7, 7, 2
// 	res := x
// 	if y > res {
// 		res = y
// 	}
// 	if z > res {
// 		res = z
// 	}
// 	fmt.Println(res)
// }

//Задача про слона в шахматах

// package main

// import "fmt"

// func main() {
// 	a, b := 3, 1
// 	x, y := 7, 8
// 	if (a - b) == (x - y) {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// Ход короля

// package main

// import "fmt"

// func main() {
// 	a, b := 4, 4
// 	x, y := 1, 3

// 	if (a-x)*(b-y) <= 1 {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// Шоколадка

// package main

// import "fmt"

// func main() {
// 	n, m, k := 3, 2, 1

// 	if k < n*m && ((k%n == 0) || (k%m == 0)) {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

//Уравнения

// package main

// import "fmt"

// func main() {
// 	a, b := 6, -2
// 	x := -b / a
// 	if -b%a == 0 {
// 		fmt.Println(x)
// 	} else if -b%a != 0 {
// 		fmt.Println("")
// 	} else if x < 0 {
// 		fmt.Println("NO")
// 	}
// }
// Не нашел как правильно решить, не понимаю условие

// Сдача

// package main

// import "fmt"

// func main() {
// 	a, b, c, d := 2, 17, 1, 18
// 	// e, f := 0, 0
// 	if c >= a {
// 		e := ((c*100 + d) - (a*100 + b)) / 100
// 		f := ((c*100 + d) - (a*100 + b)) % 100

// 		fmt.Println(e, f)
// 	} else {
// 		fmt.Println("Недостаточно денег")
// 	}
// }

// Котлеты

// package main

// import "fmt"

// func main() {
// 	k, m, n := 2, 5, 5
// 	if n-k == 0 {
// 		fmt.Printf("Котлеты можно пожарить за один раз в течении %v минут", m*2)
// 	} else if n*2%k == 0 {
// 		fmt.Printf("%v котлет можно пожарить за %v минут", n, m*(n*2/k))
// 	} else {
// 		fmt.Printf("%v котлеток можно пожарить за %v минут", n, m*(1+(n*2/k)))
// 	}
// }
// Не совсем понятно правильно или нет

// Треугольник

// package main

// import "fmt"

// func main() {
// 	a, b, c := 3, 4, 1

// 	if (a+b) > c && (a+c) > b && (b+c) > a {
// 		fmt.Printf("YES")
// 	} else {
// 		fmt.Printf("NO")
// 	}

// }

//Квадратное уравнение

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	a, b, c := 1, 0, 0
// 	d := b ^ 2 - 4*a*c
// 	x := (-b + int(math.Sqrt(d))) / 2 * a
// 	y := (-b - int(math.Sqrt(d))) / 2 * a
// 	if d < 0 {
// 		fmt.Printf("Нет решений")
// 	} else if d == 0 {
// 		x := -b / 2 * a
// 		fmt.Println(x)
// 	} else if d%2 == 0 {

// 		x := (-b + int(math.Sqrt(d))) / 2 * a
// 		y := (-b - int(math.Sqrt(d))) / 2 * a
// 		fmt.Println(x, y)
// 	}
// }
// не ясно почему ругается на дискриминант. Возможно есть некоторые нюансы, которые я не знаю

//Коровы

// package main

// import "fmt"

// func main() {
// 	n := 78
// 	if (n-1)%10 == 0 {
// 		fmt.Printf("%v корова", n)
// 	} else if n%10 > 1 && n%10 < 5 {
// 		fmt.Printf("%v коровы", n)
// 	} else {
// 		fmt.Printf("%v коров", n)
// 	}
// }

//Билеты на метро
/*Выведите пять целых чисел, равные необходимому количеству билетов
на 1, на 5, на 10, на 20, на 60 поездок. Если для какого-то данного n существует
несколько способов приобретения билетов одинаковой стоимости,
необходимо вывести ту комбинацию билетов, которая дает большее число поездок.*/

package main

import "fmt"

func main() {
	const (
		cost1  = 15
		cost5  = 70
		cost10 = 125
		cost20 = 230
		cost60 = 440
	)
	k1, k5, k10, k20, k60 := 1, 5, 10, 20, 60
	var n = 12

	k60 = n / 60
	n %= 60
	k20 = n / 20
	n %= 20
	k10 = n / 10
	n %= 10
	k5 = n / 5
	k1 = n % 5

	if k1*cost1 >= cost5 {
		k1 = 0
		k5 = 5
	}
	if k1*cost1+k5*cost5 >= cost10 {
		k1 = 0
		k5 = 0
		k10 = 10
	}
	if k1*cost1+k5*cost5+k10*cost10 >= cost20 {
		k1 = 0
		k5 = 0
		k10 = 0
		k20 = 20
	}
	if k1*cost1+k5*cost5+k10*cost10+k20*cost20 >= cost60 {
		k1 = 0
		k5 = 0
		k10 = 0
		k20 = 0
		k60 = 60
	}
	fmt.Println(k1, k5, k10, k20, k60)
}
