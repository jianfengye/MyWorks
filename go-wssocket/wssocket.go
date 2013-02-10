package main

import(
    "net"
    "log"
    "strings"
    "crypto/sha1"
    "io"
    "encoding/base64"
    "errors"
)

func main() {
    ln, err := net.Listen("tcp", ":8000")
    if err != nil {
        log.Panic(err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println("Accept err:", err)
        }
        for {
            handleConnection(conn)
        }
    }
}

func handleConnection(conn net.Conn) {
    content := make([]byte, 1024)
    _, err := conn.Read(content)
    log.Println(string(content))
    if err != nil {
        log.Println(err)
    }

    isHttp := false
    // 先暂时这么判断
    if string(content[0:3]) == "GET" {
        isHttp = true;
    }
    log.Println("isHttp:", isHttp)
    if isHttp {
        headers := parseHandshake(string(content))
        log.Println("headers", headers)
        secWebsocketKey := headers["Sec-WebSocket-Key"]

        // NOTE：这里省略其他的验证
        guid := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

        // 计算Sec-WebSocket-Accept
        h := sha1.New()
        log.Println("accept raw:", secWebsocketKey + guid)

        io.WriteString(h, secWebsocketKey + guid)
        accept := make([]byte, 28)
        base64.StdEncoding.Encode(accept, h.Sum(nil))
        log.Println(string(accept))

        response := "HTTP/1.1 101 Switching Protocols\r\n"
        response = response + "Sec-WebSocket-Accept: " + string(accept) + "\r\n"
        response = response + "Connection: Upgrade\r\n"
        response = response + "Upgrade: websocket\r\n\r\n" 
        
            
        log.Println("response:", response)
        if lenth, err := conn.Write([]byte(response)); err != nil {
            log.Println(err)
        } else {
            log.Println("send len:", lenth)
        }

        wssocket := NewWsSocket(conn)
        for {
            data, err := wssocket.ReadIframe()
            if err != nil {
                log.Println("readIframe err:" , err)
            }
            log.Println("read data:", string(data))
            err = wssocket.SendIframe([]byte("good"))
            if err != nil {
                log.Println("sendIframe err:" , err)
            }
            log.Println("send data")
        }
        
    } else {
        log.Println(string(content))
        // 直接读取
    }
}

type WsSocket struct {
    MaskingKey []byte
    Conn net.Conn
}

func NewWsSocket(conn net.Conn) *WsSocket {
    return &WsSocket{Conn: conn}
}

func (this *WsSocket)SendIframe(data []byte) error {
    // 这里只处理data长度<125的
    if len(data) >= 125 {
        return errors.New("send iframe data error")
    }

    lenth := len(data)
    maskedData := make([]byte, lenth)
    for i := 0; i < lenth; i++ {
        if this.MaskingKey != nil {
            maskedData[i] = data[i] ^ this.MaskingKey[i % 4]
        } else {
            maskedData[i] = data[i]
        }
    }

    this.Conn.Write([]byte{0x81})

    var payLenByte byte
    if this.MaskingKey != nil && len(this.MaskingKey) != 4 {
        payLenByte = byte(0x80) | byte(lenth)
        this.Conn.Write([]byte{payLenByte})
        this.Conn.Write(this.MaskingKey)
    } else {
        payLenByte = byte(0x00) | byte(lenth)
        this.Conn.Write([]byte{payLenByte})
    }
    this.Conn.Write(data)
    return nil
}

func (this *WsSocket)ReadIframe() (data []byte, err error){
    err = nil

    //第一个字节：FIN + RSV1-3 + OPCODE
    opcodeByte := make([]byte, 1)
    this.Conn.Read(opcodeByte)

    FIN := opcodeByte[0] >> 7
    RSV1 := opcodeByte[0] >> 6 & 1
    RSV2 := opcodeByte[0] >> 5 & 1
    RSV3 := opcodeByte[0] >> 4 & 1
    OPCODE := opcodeByte[0] & 15
    log.Println(RSV1,RSV2,RSV3,OPCODE)

    payloadLenByte := make([]byte, 1)
    this.Conn.Read(payloadLenByte)
    payloadLen := int(payloadLenByte[0] & 0x7F)
    mask := payloadLenByte[0] >> 7

    if payloadLen == 127 {
        extendedByte := make([]byte, 8)
        this.Conn.Read(extendedByte)
    }
    
    maskingByte := make([]byte, 4)
    if mask == 1 {
        this.Conn.Read(maskingByte)
        this.MaskingKey = maskingByte
    }

    payloadDataByte := make([]byte, payloadLen)
    this.Conn.Read(payloadDataByte)
    log.Println("data:", payloadDataByte)

    dataByte := make([]byte, payloadLen)
    for i := 0; i < payloadLen; i++ {
        if mask == 1 {
            dataByte[i] = payloadDataByte[i] ^ maskingByte[i % 4]
        } else {
            dataByte[i] = payloadDataByte[i]
        }
    }

    if FIN == 1 {
        data = dataByte
        return
    }

    nextData, err := this.ReadIframe()
    if err != nil {
        return
    }
    data = append(data, nextData...)
    return
}

func parseHandshake(content string) map[string]string {
    headers := make(map[string]string, 10)
    lines := strings.Split(content, "\r\n")

    for _,line := range lines {
        if len(line) >= 0 {
            words := strings.Split(line, ":")
            if len(words) == 2 {
                headers[strings.Trim(words[0]," ")] = strings.Trim(words[1], " ")
            }
        }
    }
    return headers
}