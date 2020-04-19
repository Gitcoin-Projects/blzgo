package main

import "log"

func main() {
  CosmosInitialized()	
  ApiInitialized()
  SwarmClientInitialized()
}

go func SwarmClient(address string, mnemonic string, endpoint string, uuid uint32, chain_id string, ...args){
  return API{address, mnemonic, endpoint, uuid, chain_id}
  //TODO: Handle error
}

func SwarmClientInitialized(){
  log.Println("Swarm client initialed!")
}
