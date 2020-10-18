package koinos

import (
    "math/big"
    "encoding/binary"
)

type Serializeable interface {
    Serialize(vb VariableBlob)
}

// --------------------------------
//  String
// --------------------------------

type String string

func (n String) Serialize(vb VariableBlob) VariableBlob {
    return vb
}

func DeserializeString(vb VariableBlob) (uint64,String) {
    return 0,String("")
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
    return append(vb, b)
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
    return append(vb, byte(n))
}

func DeserializeInt8(vb VariableBlob) (uint64,Int8) {
    return 1, Int8(vb[0])
}

// --------------------------------
//  UInt8
// --------------------------------

type UInt8 uint8

func (n UInt8) Serialize(vb VariableBlob) VariableBlob {
    return append(vb, byte(n))
}

func DeserializeUInt8(vb VariableBlob) (uint64,UInt8) {
    return 1, UInt8(vb[0])
}

// --------------------------------
//  Int16
// --------------------------------

type Int16 int16

func (n Int16) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, uint16(n))
    return append(vb, b...)
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
    return append(vb, b...)
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
    return append(vb, b...)
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
    return append(vb, b...)
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
    return append(vb, b...)
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
    return append(vb, b...)
}

func DeserializeUInt64(vb VariableBlob) (uint64,UInt64) {
    return 8, UInt64(binary.BigEndian.Uint64(vb))
}

// ----------------------------------------
//  Int128
// ----------------------------------------

type Int128 struct {
    Value big.Int
}

func NewInt128(value string) Int128 {
    var result Int128 = Int128{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return result
}

func (n Int128) Serialize(vb VariableBlob) VariableBlob {
    return append(vb, SerializeBigInt(&n.Value, 16, true)...)
}

func DeserializeInt128(vb VariableBlob) (uint64,Int128) {
    bi := Int128{Value:*DeserializeBigInt(vb, 16, true)}
    return 16,bi
}

// ----------------------------------------
//  UInt128
// ----------------------------------------

type UInt128 struct {
    Value big.Int
}

func NewUInt128(value string) UInt128 {
    var result UInt128 = UInt128{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return result
}

func (n UInt128) Serialize(vb VariableBlob) VariableBlob {
    return append(vb, SerializeBigInt(&n.Value, 20, false)...)
}

func DeserializeUInt128(vb VariableBlob) (uint64,UInt128) {
    bi := UInt128{Value:*DeserializeBigInt(vb, 20, false)}
    return 16,bi
}

// ----------------------------------------
//  Int160
// ----------------------------------------

type Int160 struct {
    Value big.Int
}

func NewInt160(value string) Int160 {
    var result Int160 = Int160{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return result
}

func (n Int160) Serialize(vb VariableBlob) VariableBlob {
    return append(vb, SerializeBigInt(&n.Value, 16, true)...)
}

func DeserializeInt160(vb VariableBlob) (uint64,Int160) {
    bi := Int160{Value:*DeserializeBigInt(vb, 16, true)}
    return 16,bi
}

// ----------------------------------------
//  UInt160
// ----------------------------------------

type UInt160 struct {
    Value big.Int
}

func NewUInt160(value string) UInt160 {
    var result UInt160 = UInt160{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return result
}

func (n UInt160) Serialize(vb VariableBlob) VariableBlob {
    return append(vb, SerializeBigInt(&n.Value, 20, false)...)
}

func DeserializeUInt160(vb VariableBlob) (uint64,UInt160) {
    bi := UInt160{Value:*DeserializeBigInt(vb, 20, false)}
    return 20,bi
}

// ----------------------------------------
//  Int256
// ----------------------------------------

type Int256 struct {
    Value big.Int
}

func NewInt256(value string) Int256 {
    var result Int256 = Int256{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return result
}

func (n Int256) Serialize(vb VariableBlob) VariableBlob {
    return append(vb, SerializeBigInt(&n.Value, 32, true)...)
}

func DeserializeInt256(vb VariableBlob) (uint64,Int256) {
    bi := Int256{Value:*DeserializeBigInt(vb, 32, true)}
    return 32,bi
}

// ----------------------------------------
//  UInt256
// ----------------------------------------

type UInt256 struct {
    Value big.Int
}

func NewUInt256(value string) UInt256 {
    var result UInt256 = UInt256{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return result
}

func (n UInt256) Serialize(vb VariableBlob) VariableBlob {
    return append(vb, SerializeBigInt(&n.Value, 32, false)...)
}

func DeserializeUInt256(vb VariableBlob) (uint64,UInt256) {
    bi := UInt256{Value:*DeserializeBigInt(vb, 32, false)}
    return 32,bi
}

// --------------------------------
//  VariableBlob
// --------------------------------

type VariableBlob []byte

// TODO: Make this variadic for size and size_hint
func NewVariableBlob() VariableBlob {
    return VariableBlob(make([]byte, 0))
}

func (n VariableBlob) Serialize(vb VariableBlob) VariableBlob {
    header := make([]byte, binary.MaxVarintLen64)
    bytes := binary.PutUvarint(header, uint64(len(n)))
    vb = append(vb, header[:bytes]...)
    return append(vb, n...)
}

func DeserializeVariableBlob(vb VariableBlob) (uint64,VariableBlob) {
    size,bytes := binary.Uvarint(vb)
    var result VariableBlob = VariableBlob(make([]byte, 0, size))
    return uint64(uint64(bytes)+size), append(result, vb[bytes:]...)
}

// --------------------------------
//  TimestampType
// --------------------------------

type TimestampType uint64

func (n TimestampType) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(n))
    return append(vb, b...)
}

func DeserializeTimestampType(vb VariableBlob) (uint64,TimestampType) {
    return 8, TimestampType(binary.BigEndian.Uint64(vb))
}

// --------------------------------
//  BlockHeightType
// --------------------------------

type BlockHeightType uint64

func (n BlockHeightType) Serialize(vb VariableBlob) VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(n))
    return append(vb, b...)
}

func DeserializeBlockHeightType(vb VariableBlob) (uint64,BlockHeightType) {
    return 8, BlockHeightType(binary.BigEndian.Uint64(vb))
}

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

func (n Multihash) Serialize(vb VariableBlob) VariableBlob {
    return vb
}

func DeserializeMultihash(vb VariableBlob) (uint64,Multihash) {
    return 0,Multihash{}
}

// --------------------------------
//  Multihash Vector
// --------------------------------

type MultihashVector struct {
    Id UInt64
    Digests []VariableBlob
}

func (n MultihashVector) Serialize(vb VariableBlob) VariableBlob {
    return vb
}

func DeserializeMultihashVector(vb VariableBlob) (uint64,MultihashVector) {
    return 0,MultihashVector{}
}

// --------------------------------
//  Utility Functions
// --------------------------------

func SerializeBigInt(num *big.Int, byte_size int, signed bool) VariableBlob {
    v := VariableBlob(make([]byte, byte_size))

    if signed && num.Sign() == -1 {
        num = num.Add(big.NewInt(1), num)
        v = num.FillBytes(v)
        for i := 0; i < byte_size; i++ {
            v[i] = ^v[i]
        }
        return v
    }

    v = num.FillBytes(v)
    return v
}

func DeserializeBigInt(vb VariableBlob, byte_size int, signed bool) *big.Int {
    num := new(big.Int)
    v := VariableBlob(make([]byte, byte_size))
    _ = copy(v, vb[:byte_size])
    if signed && (0x80 & v[0]) == 0x80 {
        for i := 0; i < byte_size; i++ {
            v[i] = ^v[i]
        }
        neg := big.NewInt(-1)
        return num.SetBytes(v).Mul(neg, num).Add(neg, num)
    }

    return num.SetBytes(v[:byte_size])
}
