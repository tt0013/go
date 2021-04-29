package main

import "fmt"

func main() {
	// fmt.Println("-------------------")

	// str := "hello沙河小王子"
	// count := 0
	// for _, s := range str {
	// 	fmt.Printf("%c\n", s)
	// 	if len(string(s)) >= 3 {
	// 		count++
	// 	}
	// }
	// fmt.Printf("\"%s\"中的汉子数量是:%d", str, count)
	s := "hello沙河小王子"
	// for i := 0; i < len(s); i++{
	// 	fmt.Printf("%v(%c)\n", s[i], s[i])
	// } 
	// fmt.Println()
	count := 0
	for _, r := range s { //rune
		fmt.Printf("%v\n", r)
		if len(string(r)) >= 3 {
			count++
		}
	}
	// fmt.Println()
	fmt.Printf("%v", count)
}
