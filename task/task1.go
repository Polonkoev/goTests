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
