#pragma once

#include <cstddef>
#include <cstdint>
#include <vector>

#define B64_ENCODE_TABLE                                           \
"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

#define B64_URLSAFE_ENCODE_TABLE                                   \
"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_="

#define B64_DECODE_TABLE                                           \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\x3e\xff\xff\xff\x3f" \
"\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\xff\xff\xff\x40\xff\xff" \
"\xff\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e" \
"\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\xff\xff\xff\xff\xff" \
"\xff\x1a\x1b\x1c\x1d\x1e\x1f\x20\x21\x22\x23\x24\x25\x26\x27\x28" \
"\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30\x31\x32\x33\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff"

#define B64_URLSAFE_DECODE_TABLE                                   \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\x3e\xff\xff" \
"\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\xff\xff\xff\x40\xff\xff" \
"\xff\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e" \
"\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\xff\xff\xff\xff\x3f" \
"\xff\x1a\x1b\x1c\x1d\x1e\x1f\x20\x21\x22\x23\x24\x25\x26\x27\x28" \
"\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30\x31\x32\x33\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff" \
"\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff"

namespace koinos::pack::util {

inline void encode_base64_table( const char* begin, size_t count, std::vector<char>& dest, const char* table, bool pad )
{
   size_t num_groups = count / 3;
   size_t odd_bytes = count % 3;
   size_t odd_out_bytes[] = {0, 2, 3};

   dest.resize(num_groups * 4 + odd_out_bytes[odd_bytes]);
   size_t i = 0, j = 0;
   uint8_t a, b, c;

   for( i=0,j=0; i<count - odd_bytes; i+=3,j+=4 )
   {
      a = uint8_t(begin[i  ]);
      b = uint8_t(begin[i+1]);
      c = uint8_t(begin[i+2]);

      // 3 bytes turns to 4 bytes
      dest[j  ] = table[                  (a >> 2)];
      dest[j+1] = table[((a&0x03) << 4) | (b >> 4)];
      dest[j+2] = table[((b&0x0f) << 2) | (c >> 6)];
      dest[j+3] = table[((c&0x3f)     )           ];
   }

   switch( odd_bytes )
   {
      case 0:
         break;
      case 1:
         a = uint8_t(begin[i  ]);
         dest[j  ] = table[                  (a >> 2)];
         dest[j+1] = table[((a&0x03) << 4)           ];
         if( pad )
         {
            dest[j+2] = table[64];
            dest[j+3] = table[64];
         }
         break;
      case 2:
         a = uint8_t(begin[i  ]);
         b = uint8_t(begin[i+1]);
         dest[j  ] = table[                  (a >> 2)];
         dest[j+1] = table[((a&0x03) << 4) | (b >> 4)];
         dest[j+2] = table[((b&0x0f) << 2)           ];
         if( pad )
         {
            dest[j+3] = table[64];
         }
         break;
      default:
         throw base_decode_error( "Illegal base64 encoder state (should never happen)" );
   }
}

inline void encode_base64( const char* begin, size_t count, std::vector<char>& dest )
{
   return encode_base64_table( begin, count, dest, B64_ENCODE_TABLE, false );
}

inline void encode_base64pad( const char* begin, size_t count, std::vector<char>& dest )
{
   return encode_base64_table( begin, count, dest, B64_ENCODE_TABLE, true );
}

inline void encode_base64url( const char* begin, size_t count, std::vector<char>& dest )
{
   return encode_base64_table( begin, count, dest, B64_URLSAFE_ENCODE_TABLE, false );
}

inline void encode_base64urlpad( const char* begin, size_t count, std::vector<char>& dest )
{
   return encode_base64_table( begin, count, dest, B64_URLSAFE_ENCODE_TABLE, true );
}

inline void decode_base64_table( const char* begin, size_t count, std::vector<char>& dest, const char* table, bool pad )
{
   if( count == 0 )
   {
      dest.resize(0);
      return;
   }

   size_t odd_bytes = count % 4;

   if( pad )
   {
      // Make sure we have the correct number of padding bytes.
      // First, padding should ensure we're a multiple of 4.
      if( odd_bytes > 0 )
      {
         throw base_decode_error( "Base64 required padding not present" );
      }

      // Count up to 3 padding characters at the end.
      char pad_char = table[0x40];
      size_t pad_bytes = 0;
      for( size_t i=count-1; i>=0; i-- )
      {
         if( begin[i] != pad_char )
            break;
         pad_bytes++;
         if( pad_bytes >= 3 )
            break;
      }

      // Proceed as if we were called with pad == false
      count -= pad_bytes;
      odd_bytes = count % 4;
   }

   size_t num_groups = count / 4;
   size_t odd_out_bytes = 0;
   if( odd_bytes > 0 )
   {
      // If we got 1 extra input byte, we don't have enough to make a full output byte; input is bad.
      if( odd_bytes == 1 )
      {
         throw base_decode_error( "Base64 nonsensical number of odd bytes present" );
      }
      // If we got 2 extra input bytes, we should have 1 extra output byte.
      // If we got 3 extra input bytes, we should have 2 extra output bytes.
      odd_out_bytes = odd_bytes-1;
   }

   dest.resize(num_groups * 3 + odd_out_bytes);
   size_t i = 0, j = 0;
   uint8_t a, b, c, d;

   for( i=0,j=0; i<count - odd_bytes; i+=4,j+=3 )
   {
      a = uint8_t(table[begin[i  ]]);
      b = uint8_t(table[begin[i+1]]);
      c = uint8_t(table[begin[i+2]]);
      d = uint8_t(table[begin[i+3]]);

      if( (a|b|c|d) & 0xc0 )
      {
         throw base_decode_error( "Base64 illegal character" );
      }

      // 4 bytes turns to 3 bytes
      dest[j  ] = char((a << 2) | (b >> 4));
      dest[j+1] = char((b << 4) | (c >> 2));
      dest[j+2] = char((c << 6) |  d      );
   }

   switch( odd_bytes )
   {
      case 0:
         break;
      case 2:
         a = uint8_t(table[begin[i  ]]);
         b = uint8_t(table[begin[i+1]]);

         dest[j  ] = char((a << 2) | (b >> 4));

         if( (a|b) & 0xc0 )
         {
            throw base_decode_error( "Base64 illegal character" );
         }
         if( (b & 0x0f) != 0 )
         {
            throw base_decode_error( "Base64 nonzero padding bits" );
         }
         break;
      case 3:
         a = uint8_t(table[begin[i  ]]);
         b = uint8_t(table[begin[i+1]]);
         c = uint8_t(table[begin[i+2]]);

         dest[j  ] = char((a << 2) | (b >> 4));
         dest[j+1] = char((b << 4) | (c >> 2));

         if( (a|b|c) & 0xc0 )
         {
            throw base_decode_error( "Base64 illegal character" );
         }
         if( (c & 0x03) != 0 )
         {
            throw base_decode_error( "Base64 nonzero padding bits" );
         }
         break;
      default:
         throw base_decode_error( "Illegal base64 decoder state (should never happen)" );
   }
}

inline void decode_base64( const char* begin, size_t count, std::vector<char>& dest )
{
   return decode_base64_table( begin, count, dest, B64_DECODE_TABLE, false );
}

inline void decode_base64pad( const char* begin, size_t count, std::vector<char>& dest )
{
   return decode_base64_table( begin, count, dest, B64_DECODE_TABLE, true );
}

inline void decode_base64url( const char* begin, size_t count, std::vector<char>& dest )
{
   return decode_base64_table( begin, count, dest, B64_URLSAFE_DECODE_TABLE, false );
}

inline void decode_base64urlpad( const char* begin, size_t count, std::vector<char>& dest )
{
   return decode_base64_table( begin, count, dest, B64_URLSAFE_DECODE_TABLE, true );
}

}
