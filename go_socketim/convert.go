package main

type ByteOrder int

const (
    BigEndian ByteOrder = iota
    LittleEndian
)

func StreamToInt32(stream []byte, byteOrder ByteOrder) int32 {
    if len(stream) != 4 {
        return 0
    }
    var u uint32
    if byteOrder == BigEndian {
        u = uint32(stream[0]) + uint32(stream[1])<<8 + uint32(stream[2])<<16 + uint32(stream[3])<<24
    } else {
        u = uint32(stream[0])<<24 + uint32(stream[1])<<16 + uint32(stream[2])<<8 + uint32(stream[3])
    }
    return int32(u)
}

func Int32ToStream(i int32, byteOrder ByteOrder) []byte {
    u := uint32(i)
    stream := [4]byte{0, 0, 0, 0}
    if byteOrder == BigEndian {
        stream[0] = byte(u)
        stream[1] = byte(u >> 8)
        stream[2] = byte(u >> 16)
        stream[3] = byte(u >> 24)
    } else {
        stream[0] = byte(u >> 24)
        stream[1] = byte(u >> 16)
        stream[2] = byte(u >> 8)
        stream[3] = byte(u)
    }
    return stream[:]
}
