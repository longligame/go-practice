package main

import "fmt"




func main() {
	var ret int
	fmt.Println("Helloworld!")

	ret = max(100,32)
	fmt.Println(ret) 
	fmt.Printf( "最大值是 : %d\n", ret)

	a, b:= swap("google","runoob")
	fmt.Println(a,b)
}

func max(num1, num2 int) int {
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

