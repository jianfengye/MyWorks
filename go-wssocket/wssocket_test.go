package wssocket

import(
    "testing"
    "io"
    "bytes"
    "log"
)

func Test_readIframe(t *testing.T) {
    input := []byte{129, 144, 19, 99, 214, 121, 96, 6, 
        184, 29, 51, 5, 164, 22, 126, 67, 181, 21, 122, 6, 184, 13};
    reader := bytes.NewReader(input)
    data, err := readIframe(reader)
    if err != nil {
        t.Error(err)
    }
    t.Error(data)
}

func readIframe(conn io.Reader) (data string, err error){
    err = nil

    //第一个字节：FIN + RSV1-3 + OPCODE
    opcodeByte := make([]byte, 1)
    conn.Read(opcodeByte)

    FIN := opcodeByte[0] >> 7
    RSV1 := opcodeByte[0] >> 6 & 1
    RSV2 := opcodeByte[0] >> 5 & 1
    RSV3 := opcodeByte[0] >> 4 & 1
    OPCODE := opcodeByte[0] & 15
    log.Println(RSV1,RSV2,RSV3,OPCODE)

    // TODO: 根据协议检测各个字节
    payloadLenByte := make([]byte, 1)
    conn.Read(payloadLenByte)
    payloadLen := int(payloadLenByte[0] & 0x7F)
    mask := payloadLenByte[0] >> 7

    if payloadLen == 127 {
        extendedByte := make([]byte, 8)
        conn.Read(extendedByte)
    }
    
    maskingByte := make([]byte, 4)
    if mask == 1 {
        conn.Read(maskingByte)
    }

    payloadDataByte := make([]byte, payloadLen)
    conn.Read(payloadDataByte)

    dataByte := make([]byte, payloadLen)
    for i := 0; i < payloadLen; i++ {
        if mask == 1 {
            dataByte[i] = payloadDataByte[i] ^ maskingByte[i % 4]
        } else {
            dataByte[i] = payloadDataByte[i]
        }
    }

    data = string(dataByte)
    if FIN == 1 {
        return
    }

    nextData, err := readIframe(conn)
    if err != nil {
        return
    }
    data = data + nextData
    return
}