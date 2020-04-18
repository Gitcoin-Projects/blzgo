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
var app_endpoint := "http://localhost:1317"
const tx_command := "txs";
//const secp256k1 := Use secp256k1-go
const prefix := 'bluzelle'
const path := "m/44'/118'/0'/0/0"
var private_key;
var account_info := "{ account_number : "", sequence : ""}"
var tx_queue := []
const token_name := 'ubnt'
const MAX_RETRIES := 10
const RETRY_INTERVAL := 1000 // 1 second

//func advance_queue()
//func begin_tx()
//func exact_error_from_message()
//func get_address()
//func get_ec_private_key()
//func handle_account_response()
//func init()
//func make_random_string()
//func next_tx()
//func query()
//func sanitize_string()
//func send_transaction()
//func send_tx()
//func sign_transaction()
//func update_account_sequence()

func CosmosInitialized() {
  log.Println("Cosmos initialized!")
}
