#include <boost/test/unit_test.hpp>

#include <koinos/pack/rt/opaque.hpp>
#include <koinos/pack/rt/binary.hpp>
#include <koinos/pack/rt/json.hpp>

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
   BOOST_CHECK_THROW( o.get_const_native(), opaque_not_unboxed );

   o.unbox();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );
   BOOST_CHECK_THROW( o.get_native(), opaque_locked );
   o.get_const_native();

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

   BOOST_TEST_MESSAGE( "Serialize binary and json" );

   auto to_blob = to_variable_blob( o );
   BOOST_CHECK( std::equal( to_blob.begin(), to_blob.end(), serialized_bin.begin(), serialized_bin.end() ) );

   o = from_variable_blob< opaque< opaque_test_object > >( serialized_bin );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), good_bin.begin(), good_bin.end() ) );

   nlohmann::json expected_json;
   expected_json["opaque"]["type"] = "opaque_test_object";
   expected_json["opaque"]["value"] = "z31SRtpx1";
   o = variable_blob{ 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A };

   nlohmann::json actual_json;
   to_json( actual_json, o );
   BOOST_CHECK( actual_json == expected_json );

   expected_json["opaque"]["value"] = "z1111";
   from_json( expected_json, o );
   to_json( actual_json, o );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), bad_bin.begin(), bad_bin.end() ) );

   expected_json = nlohmann::json();
   expected_json[ "a" ] = uint32_t(1);
   expected_json[ "b" ] = uint32_t(2);

   o = good_bin;
   to_json( actual_json, o );
   BOOST_CHECK( actual_json == expected_json );

   from_json( expected_json, o );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), good_bin.begin(), good_bin.end() ) );

   BOOST_TEST_MESSAGE( "Check state transitions between boxed, unboxed, and mutable" );

   // Check no-op boxed->boxed (box and make_immutable)
   o.box();
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );

   o.make_immutable();
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );

   // Check box->unboxed (unbox)
   o.unbox();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );

   // Check no-op unboxed->unboxed (unbox and make_immutable)
   o.unbox();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );

   o.make_immutable();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );

   // Check unboxed->mutable
   o.make_mutable();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );

   // Check no-op mutable->mutable (unbox and make_mutable)
   o.unbox();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );

   o.make_mutable();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );

   // Check mutable->boxed
   o.box();
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );

   // Check boxed->mutable
   o.make_mutable();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );

   // Check mutable->unboxed
   o.make_immutable();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );

   // Check unboxed->boxed
   o.box();
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );

   BOOST_TEST_MESSAGE( "Test modification of native type" );

   o.make_mutable();
   o->a = 3;
   o.make_immutable();
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_a.begin(), mod_bin_a.end() ) );

   o.make_mutable();
   o->b = 4;
   o.box();

   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_b.begin(), mod_bin_b.end() ) );


   BOOST_TEST_MESSAGE( "Containment of incompatible binary serialization" );

   o = bad_bin;
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );
   BOOST_CHECK_THROW( o.unbox(), std::runtime_error );
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), bad_bin.begin(), bad_bin.end() ) );


   BOOST_TEST_MESSAGE( "Assignment and constructors" );

   o = variable_blob{ 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02 };
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );
   o.unbox();
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), good_bin.begin(), good_bin.end() ) );

   opaque_test_object ote{ .a = 3, .b = 2 };
   o = ote;
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 3 );
      BOOST_CHECK_EQUAL( const_o->b, 2 );
   }

   o.box();
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_a.begin(), mod_bin_a.end() ) );

   o = opaque_test_object{ .a = 3, .b = 4 };

   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 3 );
      BOOST_CHECK_EQUAL( const_o->b, 4 );
   }

   o.box();
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_b.begin(), mod_bin_b.end() ) );

   o = opaque< opaque_test_object >();

   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );
   BOOST_CHECK_EQUAL( o->a, 0 );
   BOOST_CHECK_EQUAL( o->b, 0 );

   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), default_bin.begin(), default_bin.end() ) );

   o = opaque< opaque_test_object >( ote );
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 3 );
      BOOST_CHECK_EQUAL( const_o->b, 2 );
   }

   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_a.begin(), mod_bin_a.end() ) );

   o = opaque< opaque_test_object >( opaque_test_object{ .a = 3, .b = 4 } );
   BOOST_CHECK( o.is_unboxed() );
   BOOST_CHECK( o.is_mutable() );

   {
      const auto& const_o = o;
      BOOST_CHECK_EQUAL( const_o->a, 3 );
      BOOST_CHECK_EQUAL( const_o->b, 4 );
   }

   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_b.begin(), mod_bin_b.end() ) );

   o = opaque< opaque_test_object >( mod_bin_a );
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), mod_bin_a.begin(), mod_bin_a.end() ) );

   o = opaque< opaque_test_object >( variable_blob{ 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02 } );
   BOOST_CHECK( !o.is_unboxed() );
   BOOST_CHECK( !o.is_mutable() );
   BOOST_CHECK( std::equal( o.get_blob().begin(), o.get_blob().end(), good_bin.begin(), good_bin.end() ) );
}

BOOST_AUTO_TEST_SUITE_END()
