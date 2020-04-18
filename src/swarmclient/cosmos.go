package main

import "log"

/*
Libs to implement or import in GoLang

const util = require('./util');
const axios = require ('axios');
const ec = require ('elliptic').ec;
const bech32 = require ('bech32');
const bitcoinjs = require('bitcoinjs-lib');
const bip32 = require('bip32');
const bip39 = require('bip39');

*/
var App_endpoint = "http://localhost:1317"
const Tx_command = "txs"
//const secp256k1 := Use secp256k1-go
const Prefix = "bluzelle"
const Path = "m/44'/118'/0'/0/0"
var Private_key string
var Account_info = "{ account_number : '', sequence : ''}"
var Tx_queue []string
const Token_name = "ubnt"
const MAX_RETRIES = 10
const RETRY_INTERVAL = 1000 // 1 second

type Transaction struct {
  Type string
  data string
  deferred string
  gas_price float64
  max_gas float64
  max_fee float64
  retries_left uint8
}

//func Advance_queue()
//func Begin_tx()
//func Exact_error_from_message()
//func Get_address()
//func Get_ec_private_key()
//func Handle_account_response()
//func Init()
//func Make_random_string()
//func Next_tx()
//func Query()
//func Sanitize_string()
//func Send_transaction()
//func Send_tx()
//func Sign_transaction()
//func Update_account_sequence()

func CosmosInitialized() {
  log.Println("Cosmos initialized!")
}
