package main

import (
	"bytes"
	"fmt"
)

func sliceCheck() {
	var slice [5]int
	 s := slice[1:3]
	 fmt.Println(s)

	for i := 0; i < len(slice); i++ {
		slice[i] = i
		fmt.Println(i)
	}
	 fmt.Println(s)
	 fmt.Printf("the lenght of the original array is %d\n", len(slice))
	 fmt.Printf("the lenght of the slice is %d\n", len(s))
	 fmt.Printf("the capacity of the original array is %d\n", cap(slice))
	 fmt.Printf("the capacity of the slice is %d\n", cap(s))

}

func makeSlice(){
	// array1 := [10] int {1,2,3,}
	// sli1 := make([]int, 10)

	// multidimensional array/slice

values := [][]int{}

row1 := []int {1,2,3}
row2 := []int{5,6,7}

values = append(values, row1)
values = append(values, row2)
fmt.Println("row 1:")
fmt.Println(values[0])
fmt.Println("row2")
fmt.Println(values[1])

fmt.Println(values[1][1])//prints second element of the second row
fmt.Println(values)
fmt.Println(len(values))
fmt.Println(cap(values))
}


func bs(){
var b bytes.Buffer

b.WriteString("Hello")
b.WriteString("world")
fmt.Fprintf(&b, " the edj;ewdl\n")
b.WriteString("i hope i know")
fmt.Println(b.String())
}

func makerange(){
slice1:= make([]int, 4)
slice1[0] = 1
slice1[1] = 9
slice1[2] = 11
slice1[3] = 19

for ix, value := range slice1{
	fmt.Printf("the slice at index [%d] has the value of %d\n", ix, value)
}
}

func rs(){
	value := 0
	screen := [2][2]int{}

	for row :=range screen{
		for column := range screen{
			screen [row] [column] = value
			value = value+1
			fmt.Println(screen[row][column], "")
		}
		fmt.Printf("\n")
	}
}

func enlarge(s []int, factor int) []int{
	ns:= make([]int, len(s)*factor)
	copy(ns, s)
	return ns
}
var s [] int
func sc (){
	s = [] int {1,2,3}
	fmt.Println("the lenght of s before enlarging is:", len(s))
	fmt.Println(s)
	s = enlarge(s,5)
	fmt.Println(s)
		fmt.Println("the lenght of s after enlarging is:", len(s))

}
func mainInsert(){
	s := [] string{"m", "n", "o", "p", "q", "r"}
	in:= []string {"a", "b", "c"}
	res:= insertSlice(s, in, 0)
	fmt.Println(res)
	res= insertSlice(s,in,3)
	fmt.Println(res)
}
func insertSlice (slice, insertion []string, index int) []string{
	result:= make([]string, len(slice) + len(insertion))
	at:= copy(result, slice[:index])
	at+= copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
	}

	func strTest(){
	str := "I am zerubabel"
	substr := str[1:5]
	fmt.Println(substr)
	}

	func  strMut(){
		s:= "Hello"
		c := []byte(s)
		c[0] = 'Y'
		S2 := string(c)
		fmt.Println(S2)

		sl1 := []int {1,2,34,10}
		sl2 := []int {23,13,45,6}
		sl3 := append(sl1, sl2...)
		fmt.Println(sl3)

	}
var arr []byte = []byte{'a','b','a','a','a','c','d','e','f','g'}
func arry (){
	arrl := make([]byte,len(arr))
	fmt.Println(arrl)
}