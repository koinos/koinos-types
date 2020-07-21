#pragma once

#include <iostream>

namespace koinos::pack {

/* signed_int and unsigned_int are dumb wrappers around 64 bit integer types
 * for use in varint serialization. They are not intended to be used for
 * anything else.
 */

struct signed_int
{
   signed_int( int64_t v ) :
      value( v )
   {}

   signed_int() = default;

   int64_t value = 0;
};

static bool operator==( const signed_int &v1, const signed_int &v2 )
{
    return v1.value == v2.value;
}

static std::ostream& operator<<( std::ostream& os, const signed_int& v )
{
    os << v.value;
    return os;
}

struct unsigned_int
{
   unsigned_int( uint64_t v ) :
      value( v )
   {}

   unsigned_int() = default;

   uint64_t value = 0;
};

static bool operator==( const unsigned_int &v1, const unsigned_int &v2 )
{
    return v1.value == v2.value;
}

static std::ostream& operator<<( std::ostream& os, const unsigned_int& v )
{
    os << v.value;
    return os;
}

} // koinos::pack
