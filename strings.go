package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IdentifyPrefixPostfix(userID, email string) bool {
   return strings.HasSuffix(email, ".io") || strings.HasPrefix(userID, "UID")


}

func ContainsEducative(email string) bool {
    return strings.Contains(email, "educative")
}

func MaskUserName(email string) string {
atIndex := strings.Index(email, "@")
username:= email[:atIndex]
domain := email[atIndex:]
if len(username) > 2 {
	masked := string(username[0]) + strings.Repeat("*", len(username)-2)+ string(username[len(username)-1])
	result := masked + domain
	return result
}
return email
}

func IndexOfAtSymbol(email string) int {
return strings.Index( email, "@")
 }

func TrimAndSplitUserID(userID string) string {
    return  strings.TrimPrefix(userID, "UID-")
}

func ConvertStringToInt(str string) int {
num, _ := strconv.Atoi(str)
return num
}

func userInfo () {

	fmt.Println("strings")
    fmt.Println(IdentifyPrefixPostfix(".io", "evangeline@educative.io"))        // true
    fmt.Println(IdentifyPrefixPostfix("UID", "UID-0123"))                       // true
    fmt.Println(IdentifyPrefixPostfix("UID", "evangeline@educative.io"))        // false
    fmt.Println(ContainsEducative("evangeline@educative.io"))                   // true
    fmt.Println(MaskUserName("evangeline@educative.io"))                        // e******e@educative.io
    fmt.Println(IndexOfAtSymbol("evangeline@educative.io"))                     // 10
    fmt.Println(TrimAndSplitUserID("UID-0123"))                                 // 0123
    fmt.Println(ConvertStringToInt("123"))                                      // 123
}

func getInitial(n string) (string, string){
s := strings.ToUpper(n)
 names:= strings.Split(s, " ")

 var initial []string

 for _, v := range names{
    initial = append(initial, v[:1])
}
if len(initial) > 1{
    return initial[0], initial[1]
} else{
    return initial[0], "_"
}
}

func initialMain(){
   name1, name2 :=  getInitial("Hello World")
   fmt.Println(name1, name2)

}

