#include <boost/test/unit_test.hpp>

#include <koinos/pack/rt/opaque.hpp>
#include <koinos/pack/rt/binary.hpp>

using namespace koinos::pack;

struct opaque_fixture {};

struct opaque_test_object
{
   uint32_t a = 0;
   uint32_t b = 0;
};

KOINOS_REFLECT( opaque_test_object, (a)(b) )

BOOST_FIXTURE_TEST_SUITE( opaque_tests, opaque_fixture )

// Test boxing, unboxing, and failed unboxing

// Test locking, unlocking, and modification

// Test construction and assignment

BOOST_AUTO_TEST_CASE( opaque_boxing )
{
   variable_blob good_bin = {
      0x00, 0x00, 0x00, 0x01,
      0x00, 0x00, 0x00, 0x02
   };

   variable_blob serialized_bin = {
      0x08,
      0x00, 0x00, 0x00, 0x01,
      0x00, 0x00, 0x00, 0x02
   };

   variable_blob mod_bin_a = {
      0x00, 0x00, 0x00, 0x03,
      0x00, 0x00, 0x00, 0x02
   };

   variable_blob mod_bin_b = {
      0x00, 0x00, 0x00, 0x03,
      0x00, 0x00, 0x00, 0x04
   };

   variable_blob bad_bin = {
      0x00, 0x00, 0x00, 0x00
   };

   variable_blob default_bin = {
      0x00, 0x00, 0x00, 0x00,
      0x00, 0x00, 0x00, 0x00
   };

   opaque< opaque_test_object > o( good_bin );

   BOOST_TEST_MESSAGE( "Unbox, access, and box the opaque type" );

   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), good_bin.begin(), good_bin.end() ) );
   BOOST_CHECK_THROW( o->a = 3, opaque_not_unboxed );

   BOOST_CHECK( o.unbox() );
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );
   BOOST_CHECK_THROW( o->a = 3, opaque_locked );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 1 );
      BOOST_CHECK_EQUAL( const_o->b, 2 );
      BOOST_CHECK_EQUAL( (*const_o).a, 1 );
      BOOST_CHECK_EQUAL( (*const_o).b, 2 );
   }

   o.box();
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), good_bin.begin(), good_bin.end() ) );

   auto to_blob = to_variable_blob( o );
   BOOST_CHECK( std::equal( to_blob.begin(), to_blob.end(), serialized_bin.begin(), serialized_bin.end() ) );


   BOOST_TEST_MESSAGE( "Unbox, modidy, and box the opaque type" );

   BOOST_CHECK( o.unbox() );
   o.unlock();
   BOOST_CHECK( !o.is_locked() );
   o->a = 3;
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_a.begin(), mod_bin_a.end() ) );

   o->b = 4;
   o.box();

   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_b.begin(), mod_bin_b.end() ) );


   BOOST_TEST_MESSAGE( "Containment of incompatible binary serialization" );

   o = bad_bin;
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );
   BOOST_CHECK( !o.unbox() );
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), bad_bin.begin(), bad_bin.end() ) );


   BOOST_TEST_MESSAGE( "Assignment and constructors" );

   o = variable_blob{ 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02 };
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );
   BOOST_CHECK( o.unbox() );
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), good_bin.begin(), good_bin.end() ) );

   opaque_test_object ote{ .a = 3, .b = 2 };
   o = ote;
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 3 );
      BOOST_CHECK_EQUAL( const_o->b, 2 );
   }

   o.box();
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_a.begin(), mod_bin_a.end() ) );

   o = opaque_test_object{ .a = 3, .b = 4 };

   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 3 );
      BOOST_CHECK_EQUAL( const_o->b, 4 );
   }

   o.box();
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_b.begin(), mod_bin_b.end() ) );

   o = opaque< opaque_test_object >();

   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( !o.is_locked() );
   BOOST_CHECK_EQUAL( o->a, 0 );
   BOOST_CHECK_EQUAL( o->b, 0 );

   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), default_bin.begin(), default_bin.end() ) );

   o = opaque< opaque_test_object >( ote );
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 3 );
      BOOST_CHECK_EQUAL( const_o->b, 2 );
   }

   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_a.begin(), mod_bin_a.end() ) );

   o = opaque< opaque_test_object >( opaque_test_object{ .a = 3, .b = 4 } );
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 3 );
      BOOST_CHECK_EQUAL( const_o->b, 4 );
   }

   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_b.begin(), mod_bin_b.end() ) );

   o = opaque< opaque_test_object >( mod_bin_a );
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_a.begin(), mod_bin_a.end() ) );

   o = opaque< opaque_test_object >( variable_blob{ 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02 } );
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( o.is_locked() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), good_bin.begin(), good_bin.end() ) );
}

BOOST_AUTO_TEST_SUITE_END()
