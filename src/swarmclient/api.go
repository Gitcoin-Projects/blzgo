package main

import "log"

const Encode =  ""
const Decode = ""
const Def_mnemonic = ""
const App_service =  "crud"

type API struct {
  mnemonic string
  address string
  uuid int
  chain_id string
  endpoint string
};

//Function list
//Account()
//Count()
//Create()
//Delete()
//DeleteAll()
//Has()//TODO: Ensure encoding is safe
//Init()
//InitData()
//Keys()
//KeyValues()
//Multiupdate()
//Read()//TODO: Ensure encoding is safe
//Rename()
//Status()
//TxCount()//TODO: Convert Hex to String
//TxHas()//TODO: Convert Hex to String
//TxKeys()//TODO: Convert Hex to String
//TxKeyValues()//TODO: Convert Hex to String
//TxRead()//TODO: Convert Hex to String
//Update()
//Version()



func ApiInitialized() {
  log.Println("API initialized!")
}
