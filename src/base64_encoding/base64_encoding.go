package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	sEncode := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEncode)

	sDecode, _ := b64.StdEncoding.DecodeString(sEncode)
	fmt.Println(string(sDecode))

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
