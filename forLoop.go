package main

import "fmt"

// counter controlled
func stringLen() {
	str:= "for loop test"
	fmt.Printf("The string has a length of %d\n", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("character at position %d is : %c\n", i, str[i])
	}

}

func test2(){
	str2:= "I got the first one"
	fmt.Printf("this string has a length of %d\n", len(str2))
	 for ix:=0; ix < len(str2); ix++{
		fmt.Printf("The character at %d is : %c\n", ix, str2[ix])
	}
}

// condition controlled

func counterCon(){
	i:= 0
	for i <5 {
		fmt.Printf("this is the %d iteration\n", i)
		i++
	}

}

func rangeCheck(){
	rStr := "checking Range construct"

	for p, c := range rStr{
		fmt.Printf("chenkin the range of index %d at character %c\n", p,c)
	}
}

func continueLoop(){
	for j := 0; j <= 10; j++{
		if j == 5{
			continue
		}

		fmt.Println(j)
	}
}

func fixbuss(){
	for f:=0; f<61; f++{
		if f % 3 ==0{
			 fmt.Println("fiz\n", f)

			 continue
		}
		if f%5 ==0{
			fmt.Println("buzz\n",f)
			continue
		}

		if f%3==0 && f%5 == 0{
			fmt.Print("fizz buzz\n", f)
			continue
		} else{
			fmt.Println(f)
		}

	}
}

func isEven(){
	for i:= 0; i <=5; i++{
		if i %2 == 0{
			fmt.Printf("%d is even\n", i)
		} else{
			fmt.Printf("%d is !even\n", i)
		}
	}
}

func loop(){

	s:= "golang"
	fmt.Println(s[0:3])
	for _, char:= range "Hello"{
		fmt.Printf("%c", char)
	}
}

