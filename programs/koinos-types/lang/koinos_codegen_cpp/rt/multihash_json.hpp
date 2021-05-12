#pragma once

#include <koinos/pack/rt/binary_fwd.hpp>
#include <koinos/pack/rt/json_fwd.hpp>
#include <koinos/pack/rt/util/base58.hpp>

// TODO:  Could the dependency be made more granular?
#include <koinos/pack/rt/binary.hpp>

#include <sstream>

namespace koinos::pack {

// multihash
inline void to_json( json& j, const multihash& v )
{
   std::stringstream ss;
   to_binary( ss, v );
   std::string serialized_b58;
   util::encode_base58( serialized_b58, ss.str() );
   j = std::string("z") + serialized_b58;
}

inline void from_json( const json& j, multihash& v, uint32_t depth )
{
   KOINOS_PACK_ASSERT( j.is_string(), json_type_mismatch, "Unexpected JSON type: string exptected" );
   std::string s = j.get<std::string>();
   KOINOS_PACK_ASSERT( !(s.size() < 2), json_type_mismatch, "Multihash contains too few characters" );
   // TODO:  Accept multicodecs other than 'z'
   KOINOS_PACK_ASSERT( s[0] == 'z', json_type_mismatch, "Only z multicodec is supported for multihash" );

   std::vector<char> serialized_bytes;

   KOINOS_PACK_ASSERT( util::decode_base58( s.c_str()+1, serialized_bytes ), json_type_mismatch, "Multihash could not deserialize base58" );

   std::string serialized_bytes_str( serialized_bytes.begin(), serialized_bytes.end() );
   std::stringstream ss(serialized_bytes_str);

   from_binary( ss, v );
   KOINOS_PACK_ASSERT( ss.tellg() == serialized_bytes_str.size(), json_type_mismatch, "Multihash JSON had extra bytes" );
}

}
