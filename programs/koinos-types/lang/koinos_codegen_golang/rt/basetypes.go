
// TODO: STUBS
//type int128  int64
//type uint128 uint64

type boolean bool
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
    id uint64
    digests []variable_blob
}
