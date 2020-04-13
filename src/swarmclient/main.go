package main

import "log"

func main() {
  CosmosInitialized()	
  ApiInitialized()
  SwarmClientInitialized()
}

func SwarmClientInitialized(){
  log.Println("Bluzelle SwarmClient initialed!")
}
