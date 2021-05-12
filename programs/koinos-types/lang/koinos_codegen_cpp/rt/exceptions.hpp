#pragma once

#ifdef EXCEPTIONS_ENABLED
   #include <stdexcept>
#else
   #include <cassert>
#endif

#define KOINOS_PACK_MAX_ARRAY_ALLOC_SIZE (1024*1024*10)
#define KOINOS_PACK_MAX_RECURSION_DEPTH  (20)

namespace koinos::pack {

#ifdef EXCEPTIONS_ENABLED
   #define KOINOS_PACK_DECLARE_EXCEPTION( exception ) \
   struct exception final : std::runtime_error        \
   {                                                  \
      exception( const char* what_arg ) :             \
         std::runtime_error( what_arg ) {};           \
      virtual ~exception() override {}                \
   }

   #define KOINOS_PACK_ASSERT( cond, exc_name, msg )  \
   do {                                               \
      if( !(cond) )                                   \
      {                                               \
         throw exc_name( msg );                       \
      }                                               \
   } while (0)
#else
   #define KOINOS_PACK_DECLARE_EXCEPTION( assertion )

   #define KOINOS_PACK_ASSERT( cond, name, msg )      \
   do {                                               \
      if( !(cond) )                                   \
      {                                               \
         assert( #name ": " msg );                    \
      }                                               \
   } while (0)
#endif

/*
 * Generic serialization error. Any uses should consider being replaced
 * with a more specific/descriptire exception.
 */
KOINOS_PACK_DECLARE_EXCEPTION( parse_error );

/*
 * Parsing input recursed too deep.
 */
KOINOS_PACK_DECLARE_EXCEPTION( depth_violation );

/*
 * Parsing input would require allocation too much memory
 */
KOINOS_PACK_DECLARE_EXCEPTION( allocation_violation );

/*
 * Unexpected end of stream while packing/unpacking binary.
 */
KOINOS_PACK_DECLARE_EXCEPTION( stream_error );

/*
 * JSON parsing is out of bounds for the destination integer type.
 */
KOINOS_PACK_DECLARE_EXCEPTION( json_int_out_of_bounds );

/*
 * There was a problem serializing to the JSON object. This is probably
 * caused by programmer error.
 */
KOINOS_PACK_DECLARE_EXCEPTION( json_serialization_error );

/*
 * Incoming JSON type does not match expected type.
 */
KOINOS_PACK_DECLARE_EXCEPTION( json_type_mismatch );

/*
 * There was a problem decoding an encoded byte string
 */
KOINOS_PACK_DECLARE_EXCEPTION( json_decode_error );


KOINOS_PACK_DECLARE_EXCEPTION( bad_cast_exception );

/*
 * Opqaue type is not unboxed
 */
KOINOS_PACK_DECLARE_EXCEPTION( opaque_not_unboxed );

/*
 * Opaque type is locked
 */
KOINOS_PACK_DECLARE_EXCEPTION( opaque_locked );

} // koinos::pack

#undef KOINOS_PACK_DECLARE_EXCEPTION
