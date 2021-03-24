#pragma once

#include <koinos/pack/rt/binary_fwd.hpp>
#include <koinos/pack/rt/json_fwd.hpp>
#include <koinos/pack/rt/util/multibase.hpp>

// TODO:  Could the dependency be made more granular?
#include <koinos/pack/rt/binary.hpp>

#include <sstream>

namespace koinos::pack {

// multihash
inline void to_json( json& j, const multihash& v )
{
   std::stringstream ss;
   to_binary( ss, v );
   // Could we avoid this copy?
   std::string serialized_bytes = ss.str();

   std::vector<char> encoded;
   util::encode_multibase( serialized_bytes.c_str(), serialized_bytes.size(), encoded );
   j = std::string(encoded.begin(), encoded.end());
}

inline void from_json( const json& j, multihash& v, uint32_t depth )
{
   if( !(j.is_string()) ) throw json_type_mismatch( "Unexpected JSON type: string exptected" );
   std::string s = j.get<std::string>();
   std::vector<char> serialized_bytes;

   util::decode_multibase( s.c_str(), s.size(), serialized_bytes );

   std::string serialized_bytes_str( serialized_bytes.begin(), serialized_bytes.end() );
   std::stringstream ss(serialized_bytes_str);

   from_binary( ss, v );
   if( ss.tellg() != serialized_bytes_str.size() )
   {
      throw json_type_mismatch( "Multihash JSON had extra bytes" );
   }
}

}
