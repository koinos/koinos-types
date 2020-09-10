package koinos

using (
    "math/big"
)

// TODO: STUBS
//type int128  int64
//type uint128 uint64

type Boolean   bool
type Int8    int8
type UInt8   uint8
type Int16   int16
type UInt16  uint16
type Int32   int32
type UInt32  uint32
type Int64   int64
type UInt64  uint64

type Int128  struct {
    value big.Int
}

type UInt128 struct {
    value big.Int
}

type Int160 struct {
    value big.Int
}

type UInt160 {
    value big.Int
}

type UInt256_t {
    value big.Int
}

type UInt256_t {
    value big.Int
}

type VariableBlob []byte
type TimestampType uint64
type BlockHeightType uint64
type FixedBlob byte

// ----------------
//  Multihash
// ----------------
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

// ----------------
//  Multihash Vector
// ----------------
type MultihashVector struct {
    Id UInt64
    Digests []VariableBlob
}
