package main

import "fmt"

func bubleMain() {
	sla := []int{2, 3, 45, 10, -50, -10, 8, 56}
	fmt.Println("before sort :", sla)
	bubbleSort(sla)
	fmt.Println("after sort : ", sla)

}

func bubbleSort (sl []int){
	for pass := 1; pass < len (sl); pass ++{
		for i := 0; i < len(sl) -pass; i++{
			if sl[i] > sl[i + 1]{
				sl[i], sl[i+1] = sl[i+1], sl[i]
			}
		}
	}

}

func reverseS (s string) string{
	sl2 := []byte(s)
	for i, j := 0, len(sl2)-1; i < j; i,j = i+1, j-1{
		sl2[i], sl2[j] = sl2[j], sl2[i]
	}
	return string(sl2)
}

func reverseMain(){
	str := "Google"
	fmt.Printf(reverseS(str))
}