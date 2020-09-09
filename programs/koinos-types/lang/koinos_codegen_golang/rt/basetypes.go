package koinos

using (
    "math/big"
)

// TODO: STUBS
//type int128  int64
//type uint128 uint64

type boolean   bool
type int8_t    int8
type uint8_t   uint8
type int16_t   int16
type uint16_t  uint16
type int32_t   int32
type uint32_t  uint32
type int64_t   int64
type uint64_t  uint64

type int128_t  struct {
    value big.Int
}

type uint128_t struct {
    value big.Int
}

type int160_t struct {
    value big.Int
}

type uint160_t {
    value big.Int
}

type int256_t {
    value big.Int
}

type uint256_t {
    value big.Int
}

type variable_blob []byte
type timestamp_type uint64
type block_height_type uint64
type fixed_blob byte

// ----------------
//  Multihash
// ----------------
type multihash struct {
    id uint64
    digest variable_blob
}

func (m0 *multihash) eq(m1 *multihash) boolean {
    return false
}

func (m0 *multihash) lt(m1 *multihash) boolean {
    return false
}

func (m0 *multihash) gt(m1 *multihash) boolean {
    return false
}

// ----------------
//  Multihash Vector
// ----------------
type multihash_vector struct {
    id uint64_t
    digests []variable_blob
}
