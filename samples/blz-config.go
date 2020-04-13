package main

import "fmt"

type BlzConfig struct{
	    address string
	        mnemonic string
		    endpoint string
		        chain_id string
			    gas_params string
		    };
		    var (
			    Blzconfig = BlzConfig{"bluzelle1xhz23a58mku7ch3hx8f9hrx6he6gyujq57y3kp","volcano arrest ceiling physical concert...", "http://localhost:1317", "bluzelle", "{'max_gas': '', 'max_fee': '', 'gas_price': '10'}" }
		    )

func TestImport() {
  fmt.Println("The import worked")
}
