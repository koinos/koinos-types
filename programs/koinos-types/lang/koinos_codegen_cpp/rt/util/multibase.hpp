
#include <string>
#include <vector>

#include <koinos/pack/rt/util/base58.hpp>
#include <koinos/pack/rt/util/base64.hpp>

namespace koinos::pack::util {

inline void decode_multibase( const char* begin, size_t count, std::vector<char>& dest )
{
   if( count == 0 )
   {
      // Empty string is empty string, no base encoding needed
      dest.clear();
      return;
   }

   switch( begin[0] )
   {
      case 'm':
         decode_base64( begin+1, count-1, dest );
         break;
      case 'M':
         decode_base64pad( begin+1, count-1, dest );
         break;
      case 'u':
         decode_base64url( begin+1, count-1, dest );
         break;
      case 'U':
         decode_base64urlpad( begin+1, count-1, dest );
         break;
      case 'z':
         decode_base58( begin+1, count-1, dest );
         break;
      default:
         throw base_decode_error( "Unsupported multibase specifier" );
   }
}

inline void encode_multibase( const char* begin, size_t count, std::vector<char>& dest, char base = 'z' )
{
   std::vector<char> temp;
   switch(base)
   {
      case 'm':
         encode_base64( begin, count, temp );
         break;
      case 'M':
         encode_base64pad( begin, count, temp );
         break;
      case 'u':
         encode_base64url( begin, count, temp );
         break;
      case 'U':
         encode_base64urlpad( begin, count, temp );
         break;
      case 'z':
         encode_base58( begin, count, temp );
         break;
      default:
         throw base_decode_error( "Unsupported multibase specifier" );
   }
   // Is there an easier way to convert char to a string?
   dest.resize(temp.size()+1);
   dest[0] = base;
   memcpy( dest.data()+1, temp.data(), temp.size() );
}

/*
template< size_t N >
bool decode_base58( const std::string& src, std::array<char,N>& dest )
{
   static_assert( N < MAX_ARRAY_SIZE );
   std::vector<char> v;
   // TODO: This is inefficient. Consider ways to optimize base58 for an array
   if( !decode_base58( src.c_str(), v, N ) || v.size() != N ) return false;
   memcpy( dest.data(), v.data(), N );
   return true;
}

inline void encode_base58( std::string& s, const std::vector<char>& v )
{
   unsigned char* begin = (unsigned char*) v.data();
   s = encode_base58( begin, begin + v.size() );
}

template< size_t N >
inline void encode_base58( std::string& s, const std::array<char,N>& v )
{
   unsigned char* begin = (unsigned char*) v.data();
   s = encode_base58( begin, begin + N );
}

inline void encode_base58( std::string& s, const std::string& v )
{
   unsigned char* begin = (unsigned char*) v.c_str();
   s = encode_base58( begin, begin + v.size() );
}
*/

}
