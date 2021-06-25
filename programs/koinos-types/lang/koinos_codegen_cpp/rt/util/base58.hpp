#pragma once

#include <algorithm>
#include <vector>

#include <koinos/pack/rt/exceptions.hpp>
#include <koinos/pack/rt/util/base58impl.hpp>

namespace koinos::pack::util {

inline void decode_base58( const char* begin, size_t count, std::vector<char>& dest )
{
   // TODO:  Avoid temp copy.
   std::vector<char> temp(begin, begin+count);
   temp.push_back('\0');
   if( !impl::decode_base58(temp.data(), dest) )
   {
      throw base_decode_error( "Base58 decoding unsuccessful" );
   }
}

inline void encode_base58( const char* begin, size_t count, std::vector<char>& dest )
{
   // TODO:  Avoid temp copy.
   std::string temp = impl::encode_base58((const unsigned char*) begin, (const unsigned char*) begin+count);
   dest.resize(temp.size());
   std::copy(temp.begin(), temp.end(), dest.begin());
}

} // koinos::pack::util
