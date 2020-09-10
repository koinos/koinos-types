package koinos

import (
    "encoding/binary"
)

type Serializeable interface {
    Serialize(vb *VariableBlob)
}

// --------------------------------
// variable blob
// --------------------------------

func (n VariableBlob) Serialize(vb VariableBlob) VariableBlob {
    header := make([]byte, 8)
    binary.BigEndian.PutUint64(header, uint64(len(n)))
    vb = appendByte(vb, header...)
    return appendByte(vb, n...)
}

func DeserializeVariableBlob(vb VariableBlob) (uint64,VariableBlob) {
    var size uint64 = binary.BigEndian.Uint64(vb)
    var result VariableBlob = VariableBlob(make([]byte, 0, size))
    return 8+size, appendByte(result, vb[8:]...)
}

// --------------------------------
// Boolean
// --------------------------------

func (n boolean) Serialize(vb VariableBlob) VariableBlob {
    return appendByte(vb, byte(n))
}

func DeserializeBoolean(vb VariableBlob) (uint32,Boolean) {
    return 1, boolean(vb[0])
}

// --------------------------------
// Int8
// --------------------------------

func (n Int8) Serialize(vb VariableBlob) VariableBlob {
    return appendByte(vb, byte(n))
}

func DeserializeInt8(vb VariableBlob) (uint32,Int8) {
    return 1, Int8(vb[0])
}

// --------------------------------
// UInt8
// --------------------------------

func (n UInt8) Serialize(vb VariableBlob) VariableBlob {
    return appendByte(vb, byte(n))
}

func DeserializeUint8(vb VariableBlob) (uint32,UInt8) {
    return 1, UInt8(vb[0])
}

// --------------------------------
// Int16
// --------------------------------

func (n Int16) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, uint16(n))
    return appendByte(vb, b...)
}

func DeserializeInt16(vb VariableBlob) (uint32,Int16) {
    return 2, Int16(binary.BigEndian.Uint16(vb))
}

// --------------------------------
// UInt16
// --------------------------------

func (n UInt16) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, n)
    return appendByte(vb, b...)
}

func DeserializeUInt16(vb VariableBlob) (uint32,UInt16) {
    return 2, binary.BigEndian.Uint16(vb)
}

// --------------------------------
// Int32
// --------------------------------

func (n Int32) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(n))
    return appendByte(vb, b...)
}

func DeserializeInt32(vb VariableBlob) (uint32,Int32) {
    return 4, Int32(binary.BigEndian.Uint32(vb))
}

// --------------------------------
// UInt32
// --------------------------------

func (n UInt32) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, n)
    return appendByte(vb, b...)
}

func DeserializeUint32(vb VariableBlob) (uint32,Int32) {
    return 4, binary.BigEndian.Uint32(vb)
}

// --------------------------------
// Int64
// --------------------------------

func (n Int64) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(n))
    return appendByte(vb, b...)
}

func DeserializeInt64(vb VariableBlob) (uint32,Int64) {
    return 8, Int64(binary.BigEndian.Uint64(vb))
}

// --------------------------------
// UInt64
// --------------------------------

func (n UInt64) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, n)
    return appendByte(vb, b...)
}

func DeserializeUInt64(vb VariableBlob) (uint32,UInt64) {
    return 8, binary.BigEndian.Uint64(vb)
}

// --------------------------------
// Utility functions
// --------------------------------

func AppendToVBlob(vblob VariableBlob, data ...byte) []byte {
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
