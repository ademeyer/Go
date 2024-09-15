package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Request is a bank transactiona
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type`
	Amount float64 `json:"amount`
}

var data = `{
	"user": "Scrooge McDuck",
	"type": "deposit",
	"amount": 1234.4
}
`

func main() {
	rdr := strings.NewReader(data) // Simulate a file/socket reader

	// Decode request
	dec := json.NewDecoder(rdr)

	var req Request
	if err := dec.Decode(&req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	// print the parsed struct
	fmt.Printf("got: %+v\n", req)

	// // Create response
	// prevBalance := 1_000_000.0 // Loaded from database
	// resp := map[string]interface{}{
	// 	"ok": true,
	// 	"balance": prevBalance + req.Amount,
	// }

	// // Encode response
	// enc := json.NewDecoder(os.Stdout)
	// if err := enc.Encode(resp); err != nil {
	// 	log.Fatalf("errpr: can't encode - %s", err)
	// }

	// fmt.Printf("got: %+v\n", resp)

}
