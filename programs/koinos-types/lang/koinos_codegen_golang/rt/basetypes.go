package koinos

import (
    "bytes"
    "math/big"
    "encoding/binary"
    "encoding/json"

    "github.com/btcsuite/btcutil/base58"
)

type Serializeable interface {
    Serialize(vb *VariableBlob)
}

// --------------------------------
//  String
// --------------------------------

type String string

func (n *String) Serialize(vb *VariableBlob) *VariableBlob {
    return vb
}

func DeserializeString(vb *VariableBlob) (uint64,*String) {
    s := String("")
    return 0,&s
}

// --------------------------------
//  Boolean
// --------------------------------

type Boolean bool

func (n *Boolean) Serialize(vb *VariableBlob) *VariableBlob {
    var b byte
    if *n {
        b = 1
    }
    x := append(*vb, b)
    return &x
}

func DeserializeBoolean(vb *VariableBlob) (uint64,*Boolean) {
    var b Boolean
    if (*vb)[0] == 1 {
        b = true
    }
    return 1, &b
}

// --------------------------------
//  Int8
// --------------------------------

type Int8 int8

func (n *Int8) Serialize(vb *VariableBlob) *VariableBlob {
    ov := append(*vb, byte(*n))
    return &ov
}

func DeserializeInt8(vb *VariableBlob) (uint64,*Int8) {
    i := Int8((*vb)[0])
    return 1, &i
}

// --------------------------------
//  UInt8
// --------------------------------

type UInt8 uint8

func (n *UInt8) Serialize(vb *VariableBlob) *VariableBlob {
    ov := append(*vb, byte(*n))
    return &ov
}

func DeserializeUInt8(vb *VariableBlob) (uint64,*UInt8) {
    i := UInt8((*vb)[0])
    return 1, &i
}

// --------------------------------
//  Int16
// --------------------------------

type Int16 int16

func (n *Int16) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, uint16(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeInt16(vb *VariableBlob) (uint64,*Int16) {
    i := Int16(binary.BigEndian.Uint16(*vb))
    return 2, &i
}

// --------------------------------
//  UInt16
// --------------------------------

type UInt16 uint16

func (n *UInt16) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, uint16(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeUInt16(vb *VariableBlob) (uint64,*UInt16) {
    i := UInt16(binary.BigEndian.Uint16(*vb))
    return 2, &i
}


// --------------------------------
//  Int32
// --------------------------------

type Int32 int32

func (n *Int32) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeInt32(vb *VariableBlob) (uint64,*Int32) {
    i := Int32(binary.BigEndian.Uint32(*vb))
    return 4, &i
}

// --------------------------------
//  UInt32
// --------------------------------

type UInt32 uint32

func (n *UInt32) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeUInt32(vb *VariableBlob) (uint64,*UInt32) {
    i := UInt32(binary.BigEndian.Uint32(*vb))
    return 4, &i
}

// --------------------------------
//  Int64
// --------------------------------

type Int64 int64

func (n *Int64) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeInt64(vb *VariableBlob) (uint64,*Int64) {
    i := Int64(binary.BigEndian.Uint64(*vb))
    return 8, &i
}

// --------------------------------
//  UInt64
// --------------------------------

type UInt64 uint64

func (n *UInt64) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(*n))
    ov := append(*vb, b...)
    return &ov
}

func DeserializeUInt64(vb *VariableBlob) (uint64,*UInt64) {
    i := UInt64(binary.BigEndian.Uint64(*vb))
    return 8, &i
}

// ----------------------------------------
//  Int128
// ----------------------------------------

type Int128 struct {
    Value big.Int
}

func NewInt128(value string) *Int128 {
    var result Int128 = Int128{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return &result
}

func (n *Int128) Serialize(vb *VariableBlob) *VariableBlob {
    s := SerializeBigInt(&n.Value, 16, true)
    ov := append(*vb, *s...)
    return &ov
}

func DeserializeInt128(vb *VariableBlob) (uint64,*Int128) {
    bi := Int128{Value:*DeserializeBigInt(vb, 16, true)}
    return 16, &bi
}

func (n *Int128) MarshalJSON() ([]byte, error) {
    s := n.Value.String()
    return json.Marshal(s)
}

func (n *Int128) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    n = NewInt128(s)
    return nil
}

// ----------------------------------------
//  UInt128
// ----------------------------------------

type UInt128 struct {
    Value big.Int
}

func NewUInt128(value string) *UInt128 {
    var result UInt128 = UInt128{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return &result
}

func (n *UInt128) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 16, false)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeUInt128(vb *VariableBlob) (uint64,*UInt128) {
    bi := UInt128{Value:*DeserializeBigInt(vb, 16, false)}
    return 16, &bi
}

func (n *UInt128) MarshalJSON() ([]byte, error) {
    s := n.Value.String()
    return json.Marshal(s)
}

func (n *UInt128) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    n = NewUInt128(s)
    return nil
}

// ----------------------------------------
//  Int160
// ----------------------------------------

type Int160 struct {
    Value big.Int
}

func NewInt160(value string) *Int160 {
    var result Int160 = Int160{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return &result
}

func (n *Int160) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 20, true)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeInt160(vb *VariableBlob) (uint64,*Int160) {
    bi := Int160{Value:*DeserializeBigInt(vb, 20, true)}
    return 20, &bi
}

func (n *Int160) MarshalJSON() ([]byte, error) {
    s := n.Value.String()
    return json.Marshal(s)
}

func (n *Int160) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    n = NewInt160(s)
    return nil
}

// ----------------------------------------
//  UInt160
// ----------------------------------------

type UInt160 struct {
    Value big.Int
}

func NewUInt160(value string) *UInt160 {
    var result UInt160 = UInt160{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return &result
}

func (n *UInt160) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 20, false)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeUInt160(vb *VariableBlob) (uint64,*UInt160) {
    bi := UInt160{Value:*DeserializeBigInt(vb, 20, false)}
    return 20, &bi
}

func (n *UInt160) MarshalJSON() ([]byte, error) {
    s := n.Value.String()
    return json.Marshal(s)
}

func (n *UInt160) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    n = NewUInt160(s)
    return nil
}

// ----------------------------------------
//  Int256
// ----------------------------------------

type Int256 struct {
    Value big.Int
}

func NewInt256(value string) *Int256 {
    var result Int256 = Int256{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return &result
}

func (n *Int256) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 32, true)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeInt256(vb *VariableBlob) (uint64,*Int256) {
    bi := Int256{Value:*DeserializeBigInt(vb, 32, true)}
    return 32, &bi
}

func (n *Int256) MarshalJSON() ([]byte, error) {
    s := n.Value.String()
    return json.Marshal(s)
}

func (n *Int256) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    n = NewInt256(s)
    return nil
}

// ----------------------------------------
//  UInt256
// ----------------------------------------

type UInt256 struct {
    Value big.Int
}

func NewUInt256(value string) *UInt256 {
    var result UInt256 = UInt256{}
    nv,_ := result.Value.SetString(value, 10)
    result.Value = *nv
    return &result
}

func (n *UInt256) Serialize(vb *VariableBlob) *VariableBlob {
    x := SerializeBigInt(&n.Value, 32, false)
    ov := append(*vb, *x...)
    return &ov
}

func DeserializeUInt256(vb *VariableBlob) (uint64,*UInt256) {
    bi := UInt256{Value:*DeserializeBigInt(vb, 32, false)}
    return 32, &bi
}

func (n *UInt256) MarshalJSON() ([]byte, error) {
    s := n.Value.String()
    return json.Marshal(s)
}

func (n *UInt256) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    n = NewUInt256(s)
    return nil
}

// --------------------------------
//  VariableBlob
// --------------------------------

type VariableBlob []byte

// TODO: Make this variadic for size and size_hint
func NewVariableBlob() *VariableBlob {
    vb := VariableBlob(make([]byte, 0))
    return &vb
}

func (n *VariableBlob) Serialize(vb *VariableBlob) *VariableBlob {
    header := make([]byte, binary.MaxVarintLen64)
    bytes := binary.PutUvarint(header, uint64(len(*n)))
    ovb := append(*vb, header[:bytes]...)
    ovb = append(ovb, *n...)
    return &ovb
}

func DeserializeVariableBlob(vb *VariableBlob) (uint64,*VariableBlob) {
    size,bytes := binary.Uvarint(*vb)
    var result VariableBlob = VariableBlob(make([]byte, 0, size))
    ovb := append(result, (*vb)[bytes:uint64(bytes)+size]...)
    return uint64(uint64(bytes)+size), &ovb
}

func (n *VariableBlob) MarshalJSON() ([]byte, error) {
    nvb := NewVariableBlob()
    nvb = n.Serialize(nvb)
    s := base58.Encode(*nvb)
    return json.Marshal("z" + s)
}

func (n *VariableBlob) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return nil
    }

    // Assume base58 encoding for now
    svb := VariableBlob(base58.Decode(s[1:]))
    _,n = DeserializeVariableBlob(&svb)

    return nil
}

// --------------------------------
//  TimestampType
// --------------------------------

type TimestampType uint64

func (n *TimestampType) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(*n))
    ovb := append(*vb, b...)
    return &ovb
}

func DeserializeTimestampType(vb *VariableBlob) (uint64,*TimestampType) {
    ots := TimestampType(binary.BigEndian.Uint64(*vb))
    return 8, &ots
}

// --------------------------------
//  BlockHeightType
// --------------------------------

type BlockHeightType uint64

func (n *BlockHeightType) Serialize(vb *VariableBlob) *VariableBlob {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(*n))
    ovb := append(*vb, b...)
    return &ovb
}

func DeserializeBlockHeightType(vb *VariableBlob) (uint64,*BlockHeightType) {
    obh := BlockHeightType(binary.BigEndian.Uint64(*vb))
    return 8, &obh
}

// --------------------------------
//  Multihash
// --------------------------------

type Multihash struct {
    Id UInt64
    Digest VariableBlob
}

func (m0 *Multihash) Equals(m1 *Multihash) bool {
    return (m0.Id == m1.Id) && bytes.Equal(m0.Digest, m1.Digest)
}

func (m0 *Multihash) LessThan(m1 *Multihash) bool {
    r := m0.Id - m1.Id
    if (r < 0) {
        return true
    }
    if (r > 0) {
        return false
    }
    return (len(m0.Digest) - len(m1.Digest)) < 0
}

func (m0 *Multihash) GreaterThan(m1 *Multihash) bool {
    return !m0.Equals(m1) && !m0.LessThan(m1)
}

func (n *Multihash) Serialize(vb *VariableBlob) *VariableBlob {
    vb = n.Id.Serialize(vb)
    return n.Digest.Serialize(vb)
}

func DeserializeMultihash(vb *VariableBlob) (uint64,*Multihash) {
    omh := Multihash{}
    isize,id := DeserializeUInt64(vb)
    rvb := (*vb)[isize:]
    dsize,d := DeserializeVariableBlob(&rvb)
    omh.Id = *id
    omh.Digest = *d
    return isize+dsize, &omh
}

// --------------------------------
//  Multihash Vector
// --------------------------------

type MultihashVector struct {
    Id UInt64
    Digests []VariableBlob
}

func (n *MultihashVector) Serialize(vb *VariableBlob) *VariableBlob {
    vb = n.Id.Serialize(vb)
    header := make([]byte, binary.MaxVarintLen64)
    bytes := binary.PutUvarint(header, uint64(len(n.Digests)))
    ovb := append(*vb, header[:bytes]...)
    vb = &ovb
    for _, item := range n.Digests {
        vb = item.Serialize(vb)
    }
    return vb
}

func DeserializeMultihashVector(vb *VariableBlob) (uint64,*MultihashVector) {
    omv := MultihashVector{}
    i,id := DeserializeUInt64(vb)
    size,bytes := binary.Uvarint((*vb)[i:])
    i += uint64(bytes)
    d := make([]VariableBlob, 0, size)
    var j uint64
    var item *VariableBlob
    for num := uint64(0); num < size; num++ {
        ovb := (*vb)[i:]
        j,item = DeserializeVariableBlob(&ovb)
        i += j
        d = append(d, *item)
    }
    omv.Id = *id
    omv.Digests = d
    return i, &omv
}

// --------------------------------
//  Utility Functions
// --------------------------------

func SerializeBigInt(num *big.Int, byte_size int, signed bool) *VariableBlob {
    v := VariableBlob(make([]byte, byte_size))

    if signed && num.Sign() == -1 {
        num = num.Add(big.NewInt(1), num)
        v = num.FillBytes(v)
        for i := 0; i < byte_size; i++ {
            v[i] = ^v[i]
        }
        return &v
    }

    v = num.FillBytes(v)
    return &v
}

func DeserializeBigInt(vb *VariableBlob, byte_size int, signed bool) *big.Int {
    num := new(big.Int)
    v := VariableBlob(make([]byte, byte_size))
    _ = copy(v, (*vb)[:byte_size])
    if signed && (0x80 & v[0]) == 0x80 {
        for i := 0; i < byte_size; i++ {
            v[i] = ^v[i]
        }
        neg := big.NewInt(-1)
        return num.SetBytes(v).Mul(neg, num).Add(neg, num)
    }

    return num.SetBytes(v[:byte_size])
}
