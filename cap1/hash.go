package main

import (
	"strings"
	"fmt"

	"github.com/alexellis/hmac"
) 

func Generate(inputVar Input) Secret {

	if len(strings.TrimSpace(inputVar.Secret)) == 0 {


		panic("--secret is required")
	}
	

	// fmt.Printf("Computing hash for: %q\nSecret: %q\n", inputVar, secretVar)
	return Secret{
		Hash:    fmt.Sprintf("%x", hmac.Sign([]byte(inputVar.Message), []byte(inputVar.Secret))),
		Secret:  inputVar.Secret,
		Message: inputVar.Message,
	}
}

func(positions []int32, diff int32) int32 {
    for i, val := range positions {
        postions[i] = diff+val
    }    
	
    return positions
})
