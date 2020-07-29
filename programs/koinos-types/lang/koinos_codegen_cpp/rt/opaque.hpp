#pragma once
#include <koinos/pack/rt/binary_fwd.hpp>
#include <koinos/pack/rt/exceptions.hpp>
#include <koinos/pack/rt/reflect.hpp>


namespace koinos::types {

////////////////////////////////////////////////////////////////////////////////////
// opaque<T> state transition diagram                                             //
////////////////////////////////////////////////////////////////////////////////////
//                                                                                //
// +----------+                     +---------+                      +----------+ //
// | Boxed    | --- unbox() ------> | Unboxed | --- unlock() ------> | Unlocked | //
// | !_native |                     | _native |                      |  _native | //
// |  _blob   | <---- box() ------- | _blob   | <---- lock() ------- | !_blob   | //
// +----------+                     +---------+                      +----------+ //
//                                                                                //
////////////////////////////////////////////////////////////////////////////////////
//
// All possible transitions:
//
// Call unbox() on Boxed     -> Unboxed (see diagram)
// Call unbox() on Unboxed   -> No-op (idempotent)
// Call unbox() on Unlocked  -> No-op*
//
// Call box() on Boxed       -> No-op (idempotent)
// Call box() on Unboxed     -> Boxed (see diagram)
// Call box() on Unlocked    -> Illegal state*
//
// Call unlock() on Boxed    -> Illegal state*
// Call unlock() on Unboxed  -> Unlocked (see diagram)
// Call unlock() on Unlocked -> No-op (idempotent)
//
// Call lock() on Boxed      -> Illegal optional dereference*
// Call lock() on Unboxed    -> No-op (idempotent)*
// Call lock() on Unlocked   -> Unboxed (see diagram)
//
// *Current behavior of code, questionable correctness.
//

template< typename T >
class opaque
{
   private:
      // Native and Blob are optional containers
      //
      // When _blob exists and _native does not, the opaque is boxed
      // When _blob and _native both exist, the opaque is unboxed but not mutated
      // When _blob does not exist and _native does, the opaque is unlocked
      // _blob and _native should never both not exist.

      mutable std::optional< T >             _native;
      mutable std::optional< variable_blob > _blob;

   public:
      using type = T;

      opaque() : _native( T() ) {}
      opaque( const T& t ) : _native( t ) { serialize(); }
      opaque( T&& t ) : _native( t ) {}
      opaque( const variable_blob& v ) : _blob( v ) {}
      opaque( variable_blob&& v ) : _blob( std::move( v ) ) {}

      void unbox() const
      {
         if( !_native && _blob )
            _native.emplace( pack::from_variable_blob< T >( *_blob ) );
      }

      void box() const
      {
         if( _native && !_blob )
         {
            serialize();
         }

         _native.reset();
      }

      bool is_unboxed() const
      {
         return _native.has_value();
      }

      const variable_blob& get_blob() const
      {
         if( _native && !_blob )
         {
            serialize();
         }
         return *_blob;
      }

      operator const variable_blob&() const
      {
         return get_blob();
      }

      void lock()
      {
         serialize();
      }

      void unlock()
      {
         _blob.reset();
      }

      bool is_locked() const
      {
         return _blob.has_value();
      }

      T& get_native()
      {
         if( is_locked() ) throw pack::opaque_locked( "Opaque type is locked." );
         return *_native;
      }

      const T& get_const_native() const
      {
         if( !is_unboxed() ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return *_native;
      }

      T& operator*()
      {
         if( is_locked() ) throw pack::opaque_locked( "Opaque type is locked." );
         _blob.reset();
         return *_native;
      }

      constexpr const T& operator*() const
      {
         if( !is_unboxed() ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return *_native;
      }

      T* operator->()
      {
         if( is_locked() ) throw pack::opaque_locked( "Opaque type is locked." );
         _blob.reset();
         return &(*_native);
      }

      constexpr const T* operator->() const
      {
         if( !is_unboxed() ) throw pack::opaque_not_unboxed( "Opaque type not unboxed." );
         return &(*_native);
      }

      const variable_blob& operator=( const variable_blob& other )
      {
         _blob = other;
         _native.reset();
         return *_blob;
      }

      const variable_blob& operator=( variable_blob&& other )
      {
         _blob = std::move( other );
         _native.reset();
         return *_blob;
      }

      const T& operator=( const T& other )
      {
         _native = other;
         _blob.reset();
         return *_native;
      }

      T& operator=( T&& other )
      {
         _native = std::move( other );
         _blob.reset();
         return *_native;
      }

   private:
      void serialize() const
      {
         _blob.emplace( pack::to_variable_blob( *_native ) );
      }
};

} // koinos::types
