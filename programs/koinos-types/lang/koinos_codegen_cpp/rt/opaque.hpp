#pragma once
#include <koinos/pack/rt/binary_fwd.hpp>
#include <koinos/pack/rt/exceptions.hpp>
#include <koinos/pack/rt/reflect.hpp>

namespace koinos { namespace types {

template< typename T > class opaque;

} // types

namespace types {

template< typename T >
class opaque
{
   private:
      std::optional< T >               native;
      std::optional< variable_blob >   blob;
      bool locked = true; // Do not allow modification by default, except for defacult constructor

   public:
      using type = T;

      opaque() : native( T() ), locked( false ) {}
      opaque( const T& t ) : native( t ), locked( true ) {}
      opaque( T&& t ) : native( t ), locked( true ) {}
      opaque( const variable_blob& v ) : blob( v ), locked( true ) {}
      opaque( variable_blob&& v ) : blob( std::move( v ) ), locked( true ) {}

      bool unbox() const
      {
         if( !native && blob )
         {
            try
            {
               const_cast< std::optional< T >& >( native ).emplace( pack::from_variable_blob< T >( *blob ) );
            }
            catch( ... ) { return false; }
         }

         return true;
      }

      void box() const
      {
         if( native && !blob )
         {
            const_cast< std::optional< variable_blob >& >( blob ).emplace( pack::to_variable_blob( *native ) );
         }

         const_cast< std::optional< T >& >( native ).reset();
         const_cast< bool& >( locked ) = true;
      }

      bool is_unboxed() const
      {
         return native.has_value();
      }

      const variable_blob& get_blob() const
      {
         if( native && !blob )
         {
            const_cast< std::optional< variable_blob >& >( blob ) = pack::to_variable_blob( *native );
         }
         return *blob;
      }

      operator const variable_blob&() const
      {
         return get_blob();
      }

      void lock()
      {
         locked = true;
      }

      void unlock()
      {
         locked = false;
      }

      bool is_locked() const
      {
         return locked;
      }

      T& get_native()
      {
         if( !native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         if( locked ) throw pack::opaque_locked( "Opaque type is locked." );
         return *native;
      }

      const T& get_const_native() const
      {
         if( !native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return *native;
      }

      T& operator*()
      {
         if( !native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         if( locked ) throw pack::opaque_locked( "Opaque type is locked." );
         blob.reset();
         return *native;
      }

      constexpr const T& operator*() const
      {
         if( !native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return *native;
      }

      T* operator->()
      {
         if( !native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         if( locked ) throw pack::opaque_locked( "Opaque type is locked." );
         blob.reset();
         return &(*native);
      }

      constexpr const T* operator->() const
      {
         if( !native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return &(*native);
      }

      variable_blob& operator=( const variable_blob& other )
      {
         blob = other;
         native.reset();
         return *blob;
      }

      variable_blob& operator=( variable_blob&& other )
      {
         blob = std::move( other );
         native.reset();
         return *blob;
      }

      T& operator=( const T& other )
      {
         native = other;
         blob.reset();
         return *native;
      }

      T& operator=( T&& other )
      {
         native = std::move( other );
         blob.reset();
         return *native;
      }
};

} } // koinos::types
