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

package main

import "fmt"

func main() {
	x, y, z := 3, 4, 4
	if x > y && x > z && y != z {
		fmt.Println("X", x)
	} else if y > z && y > x && z != x {
		fmt.Println("Y", y)
	} else if z > x && z > y && x != y {
		fmt.Println("Z", z)
	} else {
		fmt.Println("Все числа равны")
	}
}
