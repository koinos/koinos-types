package koinos

import (
    "bytes"
    "math/big"
    "encoding/binary"
    "encoding/json"
    "errors"
    "unicode/utf8"

    "github.com/btcsuite/btcutil/base58"
)

type Serializeable interface {
    Serialize(vb *VariableBlob) *VariableBlob
}

// --------------------------------
//  String
// --------------------------------

type String string

func (n *String) Serialize(vb *VariableBlob) *VariableBlob {
    nvb := VariableBlob(make([]byte, len(*n)))
    copy(nvb, *n)

    return nvb.Serialize(vb)
}

func DeserializeString(vb *VariableBlob) (uint64,*String,error) {
    bytes, vb_ptr, err := DeserializeVariableBlob(vb)
    s := String("")
    if err != nil {
        return 0, &s, err
    }

    if !utf8.Valid(*vb_ptr) {
        return 0, &s, errors.New("String is not UTF-8 encoded")
    }
    s = String(*vb_ptr)

    return bytes, &s, nil
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

func DeserializeBoolean(vb *VariableBlob) (uint64,*Boolean,error) {
    var b Boolean

    if len(*vb) < 1 {
        return 0, &b, errors.New("Unexpected EOF")
    }

    if (*vb)[0] == 1 {
        b = true
    } else if (*vb)[0] != 0 {
        return 0, &b, errors.New("Boolean must be 0 or 1")
    }

    return 1, &b, nil
}

// --------------------------------
//  Int8
// --------------------------------

type Int8 int8

func (n *Int8) Serialize(vb *VariableBlob) *VariableBlob {
    ov := append(*vb, byte(*n))
    return &ov
}

func DeserializeInt8(vb *VariableBlob) (uint64,*Int8,error) {
    var i Int8

    if len(*vb) < 1 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = Int8((*vb)[0])

    return 1, &i, nil
}

// --------------------------------
//  UInt8
// --------------------------------

type UInt8 uint8

func (n *UInt8) Serialize(vb *VariableBlob) *VariableBlob {
    ov := append(*vb, byte(*n))
    return &ov
}

func DeserializeUInt8(vb *VariableBlob) (uint64,*UInt8,error) {
    var i UInt8

    if len(*vb) < 1 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = UInt8((*vb)[0])

    return 1, &i, nil
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

func DeserializeInt16(vb *VariableBlob) (uint64,*Int16,error) {
    var i Int16

    if len(*vb) < 2 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = Int16(binary.BigEndian.Uint16(*vb))

    return 2, &i, nil
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

func DeserializeUInt16(vb *VariableBlob) (uint64,*UInt16,error) {
    var i UInt16

    if len(*vb) < 2 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = UInt16(binary.BigEndian.Uint16(*vb))

    return 2, &i, nil
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

func DeserializeInt32(vb *VariableBlob) (uint64,*Int32,error) {
    var i Int32

    if len(*vb) < 4 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = Int32(binary.BigEndian.Uint32(*vb))

    return 4, &i, nil
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

func DeserializeUInt32(vb *VariableBlob) (uint64,*UInt32,error) {
    var i UInt32

    if len(*vb) < 4 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = UInt32(binary.BigEndian.Uint32(*vb))

    return 4, &i, nil
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

func DeserializeInt64(vb *VariableBlob) (uint64,*Int64,error) {
    var i Int64

    if len(*vb) < 8 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = Int64(binary.BigEndian.Uint64(*vb))

    return 8, &i, nil
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

func DeserializeUInt64(vb *VariableBlob) (uint64,*UInt64,error) {
    var i UInt64

    if len(*vb) < 8 {
        return 0, &i, errors.New("Unexpected EOF")
    }

    i = UInt64(binary.BigEndian.Uint64(*vb))

    return 8, &i, nil
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

func DeserializeInt128(vb *VariableBlob) (uint64,*Int128,error) {
    bi_ptr, err := DeserializeBigInt(vb, 16, true)
    i := Int128{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
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

func DeserializeUInt128(vb *VariableBlob) (uint64,*UInt128,error) {
    bi_ptr, err := DeserializeBigInt(vb, 16, false)
    i := UInt128{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
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

func DeserializeInt160(vb *VariableBlob) (uint64,*Int160,error) {
    bi_ptr, err := DeserializeBigInt(vb, 20, true)
    i := Int160{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
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

func DeserializeUInt160(vb *VariableBlob) (uint64,*UInt160,error) {
    bi_ptr, err := DeserializeBigInt(vb, 20, false)
    i := UInt160{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
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

func DeserializeInt256(vb *VariableBlob) (uint64,*Int256,error) {
    bi_ptr, err := DeserializeBigInt(vb, 32, true)
    i := Int256{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
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

func DeserializeUInt256(vb *VariableBlob) (uint64,*UInt256,error) {
    bi_ptr, err := DeserializeBigInt(vb, 32, false)
    i := UInt256{}

    if err != nil {
        return 0, &i, err
    }

    i.Value = *bi_ptr

    return 16, &i, nil
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

func DeserializeVariableBlob(vb *VariableBlob) (uint64,*VariableBlob,error) {
    size,bytes := binary.Uvarint(*vb)
    var result VariableBlob = VariableBlob(make([]byte, 0, size))
    if bytes <= 0 {
        return 0, &result, errors.New("Could not deserialize variable blob size")
    }

    if len(*vb) < bytes + int(size) {
        return 0, &result, errors.New("Unexpected EOF")
    }

    ovb := append(result, (*vb)[bytes:uint64(bytes)+size]...)
    return uint64(uint64(bytes)+size), &ovb, nil
}

func (n *VariableBlob) MarshalJSON() ([]byte, error) {
    nvb := NewVariableBlob()
    nvb = n.Serialize(nvb)
    s := EncodeBytes(*nvb)
    return json.Marshal(s)
}

func (n *VariableBlob) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return nil
    }

    db,err := DecodeBytes(s)
    if err != nil {
        return err
    }
    pdb := VariableBlob(db)
    _,ovb,err := DeserializeVariableBlob(&pdb)
    if err != nil {
        return err
    }
    *n = *ovb

    return nil
}

// --------------------------------
//  TimestampType
// --------------------------------

type TimestampType UInt64

func (n *TimestampType) Serialize(vb *VariableBlob) *VariableBlob {
    un := UInt64(*n)
    return un.Serialize(vb)
}

func DeserializeTimestampType(vb *VariableBlob) (uint64,*TimestampType,error) {
    i,x,err := DeserializeUInt64(vb)
    ox := TimestampType(*x)
    return i,&ox,err
}

// --------------------------------
//  BlockHeightType
// --------------------------------

type BlockHeightType UInt64

func (n *BlockHeightType) Serialize(vb *VariableBlob) *VariableBlob {
    un := UInt64(*n)
    return un.Serialize(vb)
}

func DeserializeBlockHeightType(vb *VariableBlob) (uint64,*BlockHeightType,error) {
    i,x,err := DeserializeUInt64(vb)
    ox := BlockHeightType(*x)
    return i,&ox,err
}

// --------------------------------
//  Multihash
// --------------------------------

type Multihash struct {
    Id UInt64 `json:"hash"`
    Digest VariableBlob `json:"digest"`
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
    vb = EncodeVarint(vb, uint64(n.Id))
    return n.Digest.Serialize(vb)
}

func DeserializeMultihash(vb *VariableBlob) (uint64,*Multihash,error) {
    omh := Multihash{}
    id,isize := binary.Uvarint(*vb)
    if isize <= 0 {
        return 0, &omh, errors.New("Could not deserialize multihash id")
    }
    rvb := (*vb)[isize:]
    dsize,d,err := DeserializeVariableBlob(&rvb)
    if err != nil {
        return 0, &omh, err
    }
    omh.Id = UInt64(id)
    omh.Digest = *d
    return uint64(isize)+dsize, &omh, nil
}

// --------------------------------
//  MultihashVector
// --------------------------------

type MultihashVector struct {
    Id UInt64
    Digests []VariableBlob
}

func (n *MultihashVector) Serialize(vb *VariableBlob) *VariableBlob {
    vb = EncodeVarint(vb, uint64(n.Id))
    size := uint64(0)
    if len(n.Digests) > 0 {
        size = uint64(len(n.Digests[0]))
    }
    vb = EncodeVarint(vb, size)
    vb = EncodeVarint(vb, uint64(len(n.Digests)))

    for _, item := range n.Digests {
        *vb = append(*vb, item...)
    }

    return vb
}

func DeserializeMultihashVector(vb *VariableBlob) (uint64,*MultihashVector,error) {
    omv := MultihashVector{}
    id,i := binary.Uvarint(*vb)
    if i <= 0 {
        return 0, &omv, errors.New("Could not deserialize multihash vector id")
    }
    omv.Id = UInt64(id)
    size,j := binary.Uvarint((*vb)[i:])
    if j <= 0 {
        return 0, &omv, errors.New("Could not deserialize multihash vector hash size")
    }
    i += j
    entries,j := binary.Uvarint((*vb)[i:])
    if j <= 0 {
        return 0, &omv, errors.New("Could not deserialize multihash vector size")
    }
    i += j

    if len(*vb) < i + int(entries) * int(size) {
        return 0, &omv, errors.New("Unexpected EOF")
    }

    for num := uint64(0); num < entries; num++ {
        omv.Digests = append(omv.Digests, (*vb)[i:i+int(size)])
        i += int(size)
    }

    return uint64(i), &omv, nil
}

func (n *MultihashVector) MarshalJSON() ([]byte, error) {
    mhv := struct {
        Id uint64 `json:"hash"`
        Digests []string `json:"digests"`
    }{Id: uint64(n.Id)}

    for _,item := range n.Digests {
        mhv.Digests = append(mhv.Digests, EncodeBytes(item))
    }

    return json.Marshal(&mhv)
}

func (n *MultihashVector) UnmarshalJSON(b []byte) error {
    mhv := struct {
        Id uint64 `json:"hash"`
        Digests []string `json:"digests"`
    }{}

    err := json.Unmarshal(b, &mhv)
    if err != nil {
        return err
    }

    n.Id = UInt64(mhv.Id)
    size := 0
    if len(mhv.Digests) > 0 {
        size = len(mhv.Digests[0])
    }
    for _, item := range mhv.Digests {
        if len(item) != size {
            return errors.New("Multihash vector size mismatch")
        }
        db,err := DecodeBytes(item)
        if err != nil {
            return err
        }
        n.Digests = append(n.Digests, VariableBlob(db))
    }

    return nil
}

// --------------------------------
//  Utility Functions
// --------------------------------

func EncodeBytes(b []byte) string {
    return "z" + base58.Encode(b)
}

func DecodeBytes(s string) ([]byte,error) {
    if len(s) <= 1 {
        return make([]byte, 0),nil
    }

    switch s[0] {
    case 'z':
        return base58.Decode(s[1:]),nil
    default:
        return nil,errors.New("Unknown encoding: " + string(s[0]))
    }
}

func SerializeBigInt(num *big.Int, byte_size int, signed bool) *VariableBlob {
    v := VariableBlob(make([]byte, byte_size))

    if signed && num.Sign() == -1 {
        x := big.NewInt(1)
        x = x.Add(x, num)
        v = x.FillBytes(v)
        for i := 0; i < byte_size; i++ {
            v[i] = ^v[i]
        }
        return &v
    }

    v = num.FillBytes(v)
    return &v
}

func DeserializeBigInt(vb *VariableBlob, byte_size int, signed bool) (*big.Int,error) {
    num := new(big.Int)

    if len(*vb) < byte_size {
        return num, errors.New("Unexpected EOF")
    }

    if signed && (0x80 & (*vb)[0]) == 0x80 {
        v := VariableBlob(make([]byte, byte_size))
        for i := 0; i < byte_size; i++ {
            v[i] = ^((*vb)[i])
        }
        neg := big.NewInt(-1)
        return num.SetBytes(v).Mul(neg, num).Add(neg, num), nil
    }

    return num.SetBytes((*vb)[:byte_size]), nil
}

func EncodeVarint(vb* VariableBlob, value uint64) *VariableBlob {
    header := make([]byte, binary.MaxVarintLen64)
    bytes := binary.PutUvarint(header, value)
    *vb = append(*vb, header[:bytes]...)
    return vb
}
