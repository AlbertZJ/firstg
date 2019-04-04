package main
import "fmt"

func numbers()(int , int , string ){
   a, b, c := 1, 2, "string"
   return a, b, c
}

func main(){
   _, num2, str := numbers()
   fmt.Println(num2, str) 

}


