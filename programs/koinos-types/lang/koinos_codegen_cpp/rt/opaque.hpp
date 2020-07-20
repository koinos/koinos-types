#pragma once
#include <koinos/pack/rt/binary_fwd.hpp>
#include <koinos/pack/rt/reflect.hpp>

namespace koinos { namespace types {

template< typename T > class opaque;

} // types

namespace pack {

template< typename Stream, typename T >
inline void to_binary( Stream& s, const types::opaque< T >& v );
template< typename Stream, typename T >
inline void from_binary( Stream& s, types::opaque< T >& v, uint32_t depth = 0 );

} // pack

namespace types {

template< typename T >
class opaque
{
   private:
      std::optional< T > native;
      variable_blob      blob;
      bool               valid, dirty = false;

   public:
      using type = T;

      opaque() : native( T() ), valid( true ), dirty( true ) {}
      opaque( const T& t ) : native( t ) {}
      opaque( T&& t ) : native( t ) {}

      opaque( const variable_blob& v ) : blob( v )
      {
         try
         {
            native = pack::from_variable_blob< T >( blob );
            valid = true;
         }
         catch( ... ) {}
      }

      opaque( variable_blob&& v ) : blob( v )
      {
         try
         {
            native = pack::from_variable_blob< T >( blob );
            valid = true;
         }
         catch( ... ) {}
      }

      operator variable_blob&()
      {
         if( valid && dirty ) pack::to_variable_blob< T >( blob, native );
         return blob;
      }

      operator const variable_blob&() const
      {
         // Assert !dirty
         return blob;
      }

      const variable_blob& get_blob() const
      {
         // Assert !dirty
         return blob;
      }

      operator const T&() const
      {
         return *native;
      }

      T& operator*()
      {
         // TODO: Assert
         dirty = true;
         return *native;
      }

      constexpr const T& operator*() const
      {
         // TODO: Assert
         return *native;
      }

      T* operator->()
      {
         // TODO: Assert
         dirty = true;
         return &(*native);
      }

      constexpr const T* operator->() const
      {
         // TODO: Assert
         return &(*native);
      }

      variable_blob& operator=( const variable_blob& other )
      {
         blob = other;
         try
         {
            native = pack::from_variable_blob< T >( blob );
            valid = true;
         }
         catch( ... ) {}
         return blob;
      }

      variable_blob& operator=( variable_blob&& other )
      {
         blob = other;
         try
         {
            native = pack::from_variable_blob< T >( blob );
            valid = true;
         }
         catch( ... ) {}
         return blob;
      }

      T& operator=( const T& other )
      {
         native = other;
         dirty = true;
         return native;
      }

      T& operator=( T&& other )
      {
         native = other;
         dirty = true;
         return native;
      }
};

} // types

namespace pack
{
/*
template< typename T >
struct reflector< types::opaque< T > >
{
   static auto make_tuple( const types::opaque< T >* t )
   {
      return reflector< T >::make_tuple( t );
   }
};
*/
} } // koinos::types
