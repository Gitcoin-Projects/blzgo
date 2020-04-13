package main

import "log"
import "fmt"

func main() {
	log.Println("Config: ",Blzconfig)
}

    //To test go benchmark
    func PrintFifty(){
      for i := 0; i <= 50; i++ {
        fmt.Println("For loop", i)
      }
    }
