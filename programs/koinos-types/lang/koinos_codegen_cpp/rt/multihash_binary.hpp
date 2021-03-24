#pragma once

#include <koinos/pack/rt/basetypes.hpp>
#include <koinos/pack/rt/binary_fwd.hpp>

namespace koinos::pack {

/* Multihash:
 *
 * A varint hash function (key), followed by a varint digest size in bytes,
 * followed by Network Byte Order hash for the specified length.
 */

template< typename Stream >
inline void to_binary( Stream& s, const multihash& v )
{
   to_binary( s, unsigned_int( v.id ) );
   to_binary( s, v.digest );
}

template< typename Stream >
inline void from_binary( Stream& s, multihash& v, uint32_t depth )
{
   unsigned_int id;
   from_binary( s, id );
   v.id = id.value;
   from_binary( s, v.digest );
}

}
