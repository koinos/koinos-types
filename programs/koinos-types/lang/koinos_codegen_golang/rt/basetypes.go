package koinos

import (
    "math/big"
    "encoding/binary"
)

type Serializeable interface {
    Serialize(vb VariableBlob)
}

// --------------------------------
//  Boolean
// --------------------------------

type Boolean bool

func (n Boolean) Serialize(vb VariableBlob) VariableBlob {
    var b byte
    if n {
        b = 1
    }
    return AppendToVBlob(vb, b)
}

func DeserializeBoolean(vb VariableBlob) (uint64,Boolean) {
    var b Boolean
    if vb[0] == 1 {
        b = true
    }
    return 1, b
}

// --------------------------------
//  Int8
// --------------------------------

type Int8 int8

func (n Int8) Serialize(vb VariableBlob) VariableBlob {
    return AppendToVBlob(vb, byte(n))
}

func DeserializeInt8(vb VariableBlob) (uint64,Int8) {
    return 1, Int8(vb[0])
}

// --------------------------------
//  UInt8
// --------------------------------

type UInt8 uint8

func (n UInt8) Serialize(vb VariableBlob) VariableBlob {
    return AppendToVBlob(vb, byte(n))
}

func DeserializeUint8(vb VariableBlob) (uint64,UInt8) {
    return 1, UInt8(vb[0])
}

// --------------------------------
//  Int16
// --------------------------------

type Int16 int16

func (n Int16) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, uint16(n))
    return AppendToVBlob(vb, b...)
}

func DeserializeInt16(vb VariableBlob) (uint64,Int16) {
    return 2, Int16(binary.BigEndian.Uint16(vb))
}

// --------------------------------
//  UInt16
// --------------------------------

type UInt16 uint16

func (n UInt16) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, uint16(n))
    return AppendToVBlob(vb, b...)
}

func DeserializeUInt16(vb VariableBlob) (uint64,UInt16) {
    return 2, UInt16(binary.BigEndian.Uint16(vb))
}


// --------------------------------
//  Int32
// --------------------------------

type Int32 int32

func (n Int32) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(n))
    return AppendToVBlob(vb, b...)
}

func DeserializeInt32(vb VariableBlob) (uint64,Int32) {
    return 4, Int32(binary.BigEndian.Uint32(vb))
}

// --------------------------------
//  UInt32
// --------------------------------

type UInt32 uint32

func (n UInt32) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(n))
    return AppendToVBlob(vb, b...)
}

func DeserializeUInt32(vb VariableBlob) (uint64,UInt32) {
    return 4, UInt32(binary.BigEndian.Uint32(vb))
}

// --------------------------------
//  Int64
// --------------------------------

type Int64 int64

func (n Int64) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(n))
    return AppendToVBlob(vb, b...)
}

func DeserializeInt64(vb VariableBlob) (uint64,Int64) {
    return 8, Int64(binary.BigEndian.Uint64(vb))
}

// --------------------------------
//  UInt64
// --------------------------------

type UInt64 uint64

func (n UInt64) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(n))
    return AppendToVBlob(vb, b...)
}

func DeserializeUInt64(vb VariableBlob) (uint64,UInt64) {
    return 8, UInt64(binary.BigEndian.Uint64(vb))
}

type Int128  struct {
    value big.Int
}

type UInt128 struct {
    value big.Int
}

type Int160 struct {
    value big.Int
}

type UInt160 struct {
    value big.Int
}

type Int256 struct {
    value big.Int
}

type UInt256 struct {
    value big.Int
}

// --------------------------------
//  VariableBlob
// --------------------------------

type VariableBlob []byte

func (n VariableBlob) Serialize(vb VariableBlob) VariableBlob {
    header := make([]byte, binary.MaxVarintLen64)
    bytes := binary.PutUvarint(header, uint64(len(n)))
    vb = AppendToVBlob(vb, header[:bytes]...)
    return AppendToVBlob(vb, n...)
}

func DeserializeVariableBlob(vb VariableBlob) (uint64,VariableBlob) {
    size,bytes := binary.Uvarint(vb)
    var result VariableBlob = VariableBlob(make([]byte, 0, size))
    return uint64(uint64(bytes)+size), AppendToVBlob(result, vb[bytes:]...)
}

// --------------------------------
//  TimestampType
// --------------------------------

type TimestampType uint64

func (n TimestampType) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(n))
    return AppendToVBlob(vb, b...)
}

func DeserializeTimestampType(vb VariableBlob) (uint32,TimestampType) {
    return 8, TimestampType(binary.BigEndian.Uint64(vb))
}

// --------------------------------
//  BlockHeightType
// --------------------------------

type BlockHeightType uint64

func (n BlockHeightType) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(n))
    return AppendToVBlob(vb, b...)
}

func DeserializeBlockHeightType(vb VariableBlob) (uint32,BlockHeightType) {
    return 8, BlockHeightType(binary.BigEndian.Uint64(vb))
}

type FixedBlob byte

// --------------------------------
//  Multihash
// --------------------------------
type Multihash struct {
    Id UInt64
    Digest VariableBlob
}

func (m0 *Multihash) eq(m1 *Multihash) Boolean {
    return false
}

func (m0 *Multihash) lt(m1 *Multihash) Boolean {
    return false
}

func (m0 *Multihash) gt(m1 *Multihash) Boolean {
    return false
}

// --------------------------------
//  Multihash Vector
// --------------------------------
type MultihashVector struct {
    Id UInt64
    Digests []VariableBlob
}

// --------------------------------
//  Utility functions
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
