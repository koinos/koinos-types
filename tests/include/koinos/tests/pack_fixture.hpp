#pragma once
#include <koinos/pack/classes.hpp>

namespace koinos::types {

// Helper struct for using std::visit with std::variants
template< class... Ts > struct overloaded : Ts... { using Ts::operator()...; };
template< class... Ts > overloaded( Ts... ) -> overloaded< Ts... >;

} // koinos::types

struct pack_fixture {};

struct extensions {};

struct test_object
{
   koinos::types::fixed_blob< 8 > id;
   koinos::types::multihash       key;
   std::vector< uint32_t >        vals;
   extensions                     ext;
};

KOINOS_REFLECT( test_object, (id)(key)(vals)(ext) )

KOINOS_REFLECT( extensions, )
