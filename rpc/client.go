package main

import (
    "log"
	"bytes"
    "net/rpc"
	"encoding/gob"
)

func main() {

    // Connection to rpc server
	
    client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatalf("Error in dialing. %s", err)
    }
	
    // Arguments object
    args := Args{
        M: "+38640294020",
    }
	var network bytes.Buffer

	// Encode and send a value.
	
	enc := gob.NewEncoder(&network)
	err = enc.Encode(args)
	if err != nil {
		log.Fatal("encode:", err)
	}
	
    // Store returned result
    var result Result
	
    // Call remote procedure with args
    err = client.Call("MSISDN.Multiply", enc, &result)
    if err != nil {
        log.Fatalf("error in Arith", err)
    }
    // Result is stored in result
}