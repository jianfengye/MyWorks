package main

import(
    "io"
    "log"
    "encoding/base64"
    "crypto/sha1"
)

func main() {
    secWebsocketKey := "4nxvdy1H63gBcE39fzbddw=="

    guid := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

    h := sha1.New()
    log.Println("accept raw:", secWebsocketKey + guid)

    io.WriteString(h, secWebsocketKey + guid)
    accept := make([]byte, 28)
    base64.StdEncoding.Encode(accept, h.Sum(nil))

    log.Println(string(accept))
}