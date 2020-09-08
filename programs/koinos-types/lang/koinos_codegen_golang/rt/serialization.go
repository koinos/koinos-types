package koinos

import (
    "encoding/binary"
)

type Serializeable interface {
    Serialize(vb *variable_blob)
}

// --------------------------------
// variable blob
// --------------------------------

func (n variable_blob) Serialize(vb variable_blob) variable_blob {
    header := make([]byte, 8)
    binary.LittleEndian.PutUint64(header, uint64(len(n)))
    vb = appendByte(vb, header...)
    return appendByte(vb, n...)
}

func DeserializeVariableBlob(vb variable_blob) (size_t,variable_blob) {
    var size size_t = size_t(binary.LittleEndian.Uint64(vb))
    var result variable_blob = variable_blob(make([]byte, 0, size))
    return 8+size, appendByte(result, vb[8:]...)
}

// --------------------------------
// boolean
// --------------------------------

func (n boolean) Serialize(vb variable_blob) variable_blob {
    return appendByte(vb, byte(n))
}

func DeserializeBoolean(vb variable_blob) (uint32,boolean) {
    return 1, boolean(vb[0])
}

// --------------------------------
// int8_t
// --------------------------------

func (n int8_t) Serialize(vb variable_blob) variable_blob {
    return appendByte(vb, byte(n))
}

func DeserializeInt8(vb variable_blob) (uint32,int8_t) {
    return 1, int8_t(vb[0])
}

// --------------------------------
// uint8_t
// --------------------------------

func (n uint8_t) Serialize(vb variable_blob) variable_blob {
    return appendByte(vb, byte(n))
}

func DeserializeUint8(vb variable_blob) (uint32,uint8_t) {
    return 1, uint8_t(vb[0])
}

// --------------------------------
// int16_t
// --------------------------------

func (n int16_t) Serialize(vb variable_blob) variable_blob {
    b := make([]byte, 2)
    binary.LittleEndian.PutUint16(b, uint16(n))
    return appendByte(vb, b...)
}

func DeserializeInt16(vb variable_blob) (uint32,int16_t) {
    return 2, int16_t(binary.LittleEndian.Uint16(vb))
}

// --------------------------------
// uint16_t
// --------------------------------

func (n uint16_t) Serialize(vb variable_blob) variable_blob {
    b := make([]byte, 2)
    binary.LittleEndian.PutUint16(b, n)
    return appendByte(vb, b...)
}

func DeserializeUint16(vb variable_blob) (uint32,uint16_t) {
    return 2, binary.LittleEndian.Uint16(vb)
}

// --------------------------------
// int32_t
// --------------------------------

func (n int32_t) Serialize(vb variable_blob) variable_blob {
    b := make([]byte, 4)
    binary.LittleEndian.PutUint32(b, uint32(n))
    return appendByte(vb, b...)
}

func DeserializeInt32(vb variable_blob) (uint32,int32_t) {
    return 4, int32_t(binary.LittleEndian.Uint32(vb))
}

// --------------------------------
// uint32_t
// --------------------------------

func (n uint32_t) Serialize(vb variable_blob) variable_blob {
    b := make([]byte, 4)
    binary.LittleEndian.PutUint32(b, n)
    return appendByte(vb, b...)
}

func DeserializeUint32(vb variable_blob) (uint32,uint32_t) {
    return 4, binary.LittleEndian.Uint32(vb)
}

// --------------------------------
// int64_t
// --------------------------------

func (n int64_t) Serialize(vb variable_blob) variable_blob {
    b := make([]byte, 8)
    binary.LittleEndian.PutUint64(b, uint64(n))
    return appendByte(vb, b...)
}

func DeserializeInt64(vb variable_blob) (uint32,int64_t) {
    return 8, int64_t(binary.LittleEndian.Uint64(vb))
}

// --------------------------------
// uint64_t
// --------------------------------

func (n uint64_t) Serialize(vb variable_blob) variable_blob {
    b := make([]byte, 8)
    binary.LittleEndian.PutUint64(b, n)
    return appendByte(vb, b...)
}

func DeserializeUint64(vb variable_blob) (uint32,uint64_t) {
    return 8, binary.LittleEndian.Uint64(vb)
}

// --------------------------------
// Utility functions
// --------------------------------

func AppendToVBlob(vblob variable_blob, data ...byte) []byte {
    m := len(vblob)
    n := m + len(data)
    if n > cap(vblob) { // if necessary, reallocate
        // allocate double what's needed
        new_vblob := make([]byte, (n+1)*2)
        copy(new_vblob, vblob)
        vblob = new_vblob
    }
    vblob = vblob[0:n]
    copy(vblob[m:n], data)
    return vblob
}
