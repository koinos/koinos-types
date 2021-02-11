#pragma once
#include <koinos/pack/rt/binary_fwd.hpp>
#include <koinos/pack/rt/exceptions.hpp>
#include <koinos/pack/rt/reflect.hpp>


namespace koinos {

/////////////////////////////////////////////////////////////////////////////////////
// opaque<T> state transition diagram                                              //
/////////////////////////////////////////////////////////////////////////////////////
//                                                                                 //
// +----------+                     +---------+                        +---------+ //
// | Boxed    | --- unbox() ------> | Unboxed | ---- make_mutable() -> | Mutable | //
// | !_native |                     | _native |                        | _native | //
// |  _blob   | <---- box() ------- |         | <- make_immutable() -- | !_blob  | //
// +----------+                     +---------+                        +---------+ //
//                                                                                 //
/////////////////////////////////////////////////////////////////////////////////////
//
// State transitions are implicitly composible (e.g. calling box() Mutable
// also calls make_immutable() ).
//
// Mutable is a substate of Unboxed (e.g. idempotent transitions on unboxed are
// indempotent on Mutable )
//
// All possible transitions:
//
// Call unbox() on Boxed            -> Unboxed (see diagram)
// Call unbox() on Unboxed          -> No-op (idempotent)
// Call unbox() on Mutable          -> No-op (idempotent)
//
// Call box() on Boxed              -> No-op (idempotent)
// Call box() on Unboxed            -> Boxed (see diagram)
// Call box() on Mutable            -> Boxed (composed)
//
// Call make_mutable() on Boxed     -> Mutable (composed)
// Call make_mutable() on Unboxed   -> Unlocked (see diagram)
// Call make_mutable() on Mutable   -> No-op (idempotent)
//
// Call make_immutable() on Boxed   -> No-op (idempotent)
// Call make_immutable() on Unboxed -> No-op (idempotent)
// Call make_immutable() on Mutable -> Unboxed (see diagram)

template< typename T >
class opaque
{
   private:
      mutable std::optional< T >             _native;
      mutable std::optional< variable_blob > _blob;

   public:
      using type = T;

      opaque() : _native( T() ) {}
      opaque( const T& t ) : _native( t ) {}
      opaque( T&& t ) : _native( std::move( t ) ) {}
      opaque( const variable_blob& v ) : _blob( v ) {}
      opaque( variable_blob&& v ) : _blob( std::move( v ) ) {}

      void unbox() const
      {
         if( !_native && _blob )
            _native.emplace( pack::from_variable_blob< T >( *_blob ) );
      }

      void box() const
      {
         if( _native )
         {
            // Composed state transition Mutable -> Unboxed
            if( !_blob )
               serialize();

            // State transition Unboxed -> Boxed
            _native.reset();
         }
      }

      bool is_unboxed() const
      {
         return _native.has_value();
      }

      void make_immutable()
      {
         if( _native && !_blob )
            serialize();
      }

      void make_mutable()
      {
         // Composed state transtition Boxed -> Unboxed
         if( !_native )
            unbox();

         // State transition Unboxed -> Mutable
         // Must check _native again in case unboxing failed.
         if( _native && _blob )
            _blob.reset();
      }

      bool is_mutable() const
      {
         return _native.has_value() && !_blob.has_value();
      }

      /**
       * get_blob() returns a reference to the underlying blob. This is an
       * optimization so that large blobs can be efficiently passed around.
       * It is the responsibility of the caller to ensure the opaque does
       * not go out of scope and invalidate the reference.
       */
      const variable_blob& get_blob() const
      {
         if( _native && !_blob )
         {
            serialize();
         }
         return *_blob;
      }

      /**
       * get_native() and get_const_native() returns a reference to the
       * underlying value. This is an optimization so that large objects can be
       * efficiently passed around. It is the responsibility of the caller to
       * ensure the opaque does not go out of scope and invalidate the
       * reference.
       */
      T& get_native()
      {
         if( !_native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         if( _blob ) throw pack::opaque_locked( "Opaque type is not mutable." );
         return *_native;
      }

      const T& get_const_native() const
      {
         if( !_native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return *_native;
      }

      T& operator*()
      {
         if( !_native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         if( _blob ) throw pack::opaque_locked( "Opaque type is not mutable." );
         return *_native;
      }

      constexpr const T& operator*() const
      {
         if( !_native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return *_native;
      }

      T* operator->()
      {
         if( !_native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         if( _blob ) throw pack::opaque_locked( "Opaque type is not mutable." );
         return &(*_native);
      }

      constexpr const T* operator->() const
      {
         if( !_native ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return &(*_native);
      }

   private:
      void serialize() const
      {
         _blob.emplace( pack::to_variable_blob( *_native ) );
      }
};

} // koinos

namespace koinos::pack {
template< typename T >
struct get_typename< koinos::opaque<T> >
{
   static const char* name()
   {
      static const std::string n = std::string("opaque<") + get_typename<T>::name() + ">";
      return n.c_str();
   }
};
}
