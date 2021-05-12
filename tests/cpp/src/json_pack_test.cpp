#ifdef JSON_ENABLED
#include <boost/test/unit_test.hpp>

#include <koinos/tests/pack_fixture.hpp>

#include <koinos/pack/rt/json.hpp>

struct json_pack_fixture {};

BOOST_FIXTURE_TEST_SUITE( json_pack_tests, pack_fixture )

using namespace koinos::pack;

BOOST_AUTO_TEST_CASE( basic_test )
{
   json j;
   to_json( j, int64_t(-256) );

   std::string expected = "-256";
   BOOST_REQUIRE_EQUAL( expected, j.dump() );

   int64_t res = 0;
   from_json( j, res );
   BOOST_REQUIRE_EQUAL( res, -256 );
}

BOOST_AUTO_TEST_CASE( integer_bounds )
{
   json j;
   j = json::parse( "4294967296" ); // 2^32

#ifdef EXCEPTIONS_ENABLED
   int32_t from_j;
   BOOST_REQUIRE_THROW( from_json( j, from_j ), json_int_out_of_bounds );
#endif
}

BOOST_AUTO_TEST_CASE( json_integer_bounds )
{
   json j;
   const int64_t one = 1;

   to_json( j, (one<<53) - 1 );
   std::string expected = "9007199254740991";
   BOOST_REQUIRE_EQUAL( expected, j.dump() );

   to_json( j, one<<53 );
   expected = "\"9007199254740992\"";
   BOOST_REQUIRE_EQUAL( expected, j.dump() );

   to_json( j, -((one<<53) - 1) );
   expected = "-9007199254740991";
   BOOST_REQUIRE_EQUAL( expected, j.dump() );

   to_json( j, -(one<<53) );
   expected = "\"-9007199254740992\"";
   BOOST_REQUIRE_EQUAL( expected, j.dump() );
}

BOOST_AUTO_TEST_CASE( uint128_test )
{
   uint128_t to_j = 1;
   to_j <<= 65;
   to_j -= 1; // 2^65 - 1

   json j;
   to_json( j, to_j );

   std::string expected = "\"36893488147419103231\"";
   BOOST_REQUIRE_EQUAL( expected, j.dump() );

   uint128_t from_j;
   from_json( j, from_j );
   BOOST_REQUIRE_EQUAL( from_j, to_j );

   to_j = 10;
   to_json( j, to_j );

   expected = "10";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );
}

BOOST_AUTO_TEST_CASE( uint160_test )
{
   uint160_t to_j = 1;
   to_j <<= 129;
   to_j -= 1; // 2^129 - 1

   json j;
   to_json( j, to_j );

   std::string expected = "\"680564733841876926926749214863536422911\"";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   uint160_t from_j;
   from_json( j, from_j );
   BOOST_REQUIRE_EQUAL( from_j, to_j );

   to_j = 10;
   to_json( j, to_j );

   expected = "10";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );
}

BOOST_AUTO_TEST_CASE( uint256_test )
{
   uint256_t to_j = 1;
   to_j <<= 129;
   to_j -= 1; // 2^129 - 1

   json j;
   to_json( j, to_j );

   std::string expected = "\"680564733841876926926749214863536422911\"";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   uint256_t from_j;
   from_json( j, from_j );
   BOOST_REQUIRE_EQUAL( from_j, to_j );

   to_j = 10;
   to_json( j, to_j );

   expected = "10";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );
}

BOOST_AUTO_TEST_CASE( vector_test )
{
   json j;
   vector< int16_t > to_j = { 4, 8, 15, 16, 23, 42 };

   to_json( j, to_j );

   std::string expected = "[4,8,15,16,23,42]";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   vector< int16_t > from_j;
   from_json( j, from_j );

   BOOST_REQUIRE_EQUAL( to_j.size(), from_j.size() );
   for( size_t i = 0; i < to_j.size(); ++i )
   {
      BOOST_REQUIRE_EQUAL( to_j[i], from_j[i] );
   }

#ifdef EXCEPTIONS_ENABLED
   j = json::parse( "[\"foo\",\"bar\"]" );
   BOOST_REQUIRE_THROW( from_json( j, from_j ), std::exception );
#endif
}

BOOST_AUTO_TEST_CASE( array_test )
{
   json j;
   std::array< int16_t, 6 > to_j = { 4, 8, 15, 16, 23, 42 };

   to_json( j, to_j );

   std::string expected = "[4,8,15,16,23,42]";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   std::array< int16_t, 6 > from_j;
   from_json( j, from_j );

   for( size_t i = 0; i < to_j.size(); ++i )
   {
      BOOST_REQUIRE_EQUAL( to_j[i], from_j[i] );
   }

#ifdef EXCEPTIONS_ENABLED
   j = json::parse( "[4,8,15,16,23,42,108]") ;
   BOOST_REQUIRE_THROW( from_json( j, from_j ), json_type_mismatch );

   j = json::parse( "[\"foo\",\"bar\",\"a\",\"b\",\"c\",\"d\"]" );
   BOOST_REQUIRE_THROW( from_json( j, from_j ), std::exception );
#endif
}

BOOST_AUTO_TEST_CASE( variant_test )
{
   typedef std::variant< int16_t, int32_t > test_variant;
   test_variant to_j = int16_t( 10 );

   json j;
   to_json( j, to_j );

   std::string expected = "{\"type\":\"int16_t\",\"value\":10}";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   test_variant from_j;
   from_json( j, from_j );

   std::visit( koinos::overloaded {
      []( int16_t v ){ BOOST_REQUIRE_EQUAL( v, 10 ); },
      []( int32_t v ){ BOOST_FAIL( "variant contains unexpected type" ); },
      []( auto& v ){ BOOST_FAIL( "variant contains unexpected type" ); }
   }, from_j );

   j["type"] = 0;
   from_json( j, from_j );

   std::visit( koinos::overloaded {
      []( int16_t v ){ BOOST_REQUIRE_EQUAL( v, 10 ); },
      []( int32_t v ){ BOOST_FAIL( "variant contains unexpected type" ); },
      []( auto& v ){ BOOST_FAIL( "variant contains unexpected type" ); }
   }, from_j );

   to_j = int32_t( 20 );
   to_json( j, to_j );

   expected = expected = "{\"type\":\"int32_t\",\"value\":20}";

   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   from_json( j, from_j );

   std::visit( koinos::overloaded {
      []( int16_t v ){ BOOST_FAIL( "variant contains unexpected type" ); },
      []( int32_t v ){ BOOST_REQUIRE_EQUAL( v, 20 ); },
      []( auto& v ){ BOOST_FAIL( "variant contains unexpected type" ); }
   }, from_j );

   j["type"] = 1;
   from_json( j, from_j );

   std::visit( koinos::overloaded {
      []( int16_t v ){ BOOST_FAIL( "variant contains unexpected type" ); },
      []( int32_t v ){ BOOST_REQUIRE_EQUAL( v, 20 ); },
      []( auto& v ){ BOOST_FAIL( "variant contains unexpected type" ); }
   }, from_j );

#ifdef EXCEPTIONS_ENABLED
   j["type"] = 2;
   BOOST_REQUIRE_THROW( from_json( j, from_j ), parse_error );

   j["type"] = "uint64_t";
   BOOST_REQUIRE_THROW( from_json( j, from_j ), json_type_mismatch );
#endif
}

BOOST_AUTO_TEST_CASE( optional_test )
{
   typedef std::optional< int16_t > test_optional;
   test_optional to_j;

   json j;
   to_json( j, to_j );

   BOOST_REQUIRE( j.is_null() );

   test_optional from_j;
   from_json( j, from_j );
   BOOST_REQUIRE( !from_j.has_value() );

   to_j = 10;
   to_json( j, to_j );

   std::string expected = "10";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   from_json( j, from_j );
   BOOST_REQUIRE( from_j.has_value() );
   BOOST_REQUIRE_EQUAL( *from_j, *to_j );
}

BOOST_AUTO_TEST_CASE( variable_blob_test )
{
   variable_blob to_j;
   to_j = { 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A };

   json j;
   to_json( j, to_j );

   std::string expected = "\"z31SRtpx1\"";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   variable_blob from_j;
   from_json( j, from_j );
   BOOST_REQUIRE_EQUAL( to_j.size(), from_j.size() );
   for( size_t i = 0; i < to_j.size(); ++i )
   {
      BOOST_REQUIRE_EQUAL( to_j[i], from_j[i] );
   }
}

BOOST_AUTO_TEST_CASE( fixed_blob_test )
{
   fixed_blob< 6 > to_j;
   to_j = { 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A };

   json j;
   to_json( j, to_j );

   std::string expected = "\"z31SRtpx1\"";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   fixed_blob< 6 > from_j;
   from_json( j, from_j );
   for( size_t i = 0; i < to_j.size(); ++i )
   {
      BOOST_CHECK_EQUAL( to_j[i], from_j[i] );
   }
}

void hex_to_vector(
   std::vector<char>& out,
   std::initializer_list<uint8_t> values)
{
   for( auto it=values.begin(); it!=values.end(); ++it )
   {
      out.push_back( uint8_t( *it ) );
   }
}

BOOST_AUTO_TEST_CASE( multihash_test )
{
   multihash to_j;
   to_j.id = 1;
   to_j.digest = { 0x04, 0x08, 0x0F, 0x10, 0x17, 0x2A };

   json j;
   to_json( j, to_j );

   // Expected string generated using Python base58 (pip install base58) as follows:
   // echo -n z; echo -ne '\x01\x06\x04\x08\x0f\x10\x17\x2a' | base58; echo
   std::string expected = "\"zAvtuU8Lw8u\"";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   multihash from_j;
   from_json( j, from_j );
   BOOST_REQUIRE_EQUAL( to_j.digest.size(), from_j.digest.size() );

   j = std::string("zkpt6ajQywrZ");
#ifdef EXCEPTIONS_ENABLED
   try
   {
      from_json( j, from_j );
      BOOST_REQUIRE( false );
   }
   catch( json_type_mismatch& e )
   {
      BOOST_REQUIRE_EQUAL( e.what(), "Multihash JSON had extra bytes" );
   }
#endif

   for( size_t i = 0; i < to_j.digest.size(); ++i )
   {
      BOOST_REQUIRE_EQUAL( to_j.digest[i], from_j.digest[i] );
   }

   to_j.digest.clear();
   to_j.id = 0x1220;
   hex_to_vector( to_j.digest, {
      0xBA, 0x78, 0x16, 0xBF, 0x8F, 0x01, 0xCF, 0xEA,
      0x41, 0x41, 0x40, 0xDE, 0x5D, 0xAE, 0x22, 0x23,
      0xB0, 0x03, 0x61, 0xA3, 0x96, 0x17, 0x7A, 0x9C,
      0xB4, 0x10, 0xFF, 0x61, 0xF2, 0x00, 0x15, 0xAD
   } );

   to_json( j, to_j );

   // Expected string generated using Python base58 (pip install base58) as follows:
   // echo -n z; echo -ne '\xa0\x24\x20\xba\x78\x16\xbf\x8f\x01\xcf\xea\x41\x41\x40\xde\x5d\xae\x22\x23\xb0\x03\x61\xa3\x96\x17\x7a\x9c\xb4\x10\xff\x61\xf2\x00\x15\xad' | base58; echo
   expected = "\"zGymuoch75iSPbDGPsMPFA5QTVJxkGY5Qz4eiTLNbkJY9EvMA\"";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );
}

BOOST_AUTO_TEST_CASE( reflect_test )
{
   test_object to_j;
   to_j.id = { 0, 4, 8, 15, 16, 23, 42, 0 };
   to_j.key.id = 1;
   to_j.key.digest = { 'f', 'o', 'o', 'b', 'a', 'r' };
   to_j.vals = { 108 };

   json j;
   to_json( j, to_j );
   // Expected string generated using Python base58 (pip install base58) as follows:
   // echo -n z; echo -ne '\x01\x06foobar' | base58; echo
   std::string expected = "{\"ext\":{},\"id\":\"z19rwEskdm1\",\"key\":\"zAwjubcV5mT\",\"vals\":[108]}";
   BOOST_REQUIRE_EQUAL( j.dump(), expected );

   test_object from_j;
   from_json( j, from_j );
   BOOST_REQUIRE( memcmp( from_j.id.data(), to_j.id.data(), to_j.id.size() ) == 0 );
   BOOST_REQUIRE_EQUAL( from_j.key.digest.size(), to_j.key.digest.size() );
   BOOST_REQUIRE( memcmp( from_j.key.digest.data(), to_j.key.digest.data(), to_j.key.digest.size() ) == 0 );
   BOOST_REQUIRE_EQUAL( from_j.vals.size(), to_j.vals.size() );
   BOOST_REQUIRE( memcmp( (char*)from_j.vals.data(), (char*)to_j.vals.data(), sizeof(uint32_t) * to_j.vals.size() ) == 0 );
}

BOOST_AUTO_TEST_CASE( empty_case_test )
{
   //
   // Empty vector<T>        should be []
   // Empty set<T>           should be []
   // Empty array< T, N >    should be []
   // Empty optional         should be null
   // Empty struct           should be {}
   //
   std::string json_earray = "[]";
   std::string json_eobject = "{}";
   std::string json_null = "null";
   std::string json_emhv = "{\"digests\":[],\"hash\":0}";

   std::vector< uint32_t >   empty_vector;
   std::set< uint32_t >      empty_set;
   std::array< uint32_t, 0 > empty_array;
   std::optional< uint32_t > empty_optional;
   extensions                empty_struct;

   { json j; to_json( j, empty_vector   ); BOOST_REQUIRE_EQUAL( j.dump(), json_earray  ); }
   { json j; to_json( j, empty_set      ); BOOST_REQUIRE_EQUAL( j.dump(), json_earray  ); }
   { json j; to_json( j, empty_array    ); BOOST_REQUIRE_EQUAL( j.dump(), json_earray  ); }
   { json j; to_json( j, empty_optional ); BOOST_REQUIRE_EQUAL( j.dump(), json_null    ); }
   { json j; to_json( j, empty_struct   ); BOOST_REQUIRE_EQUAL( j.dump(), json_eobject ); }
}

/*
( depth_test )
{
   // The compiler is not happy about the template deduction I am about to make it do,
   // but it will complete it...
   vector< vector< vector< vector< vector< vector< vector< vector< vector< vector< vector<
      vector< vector< vector< vector< vector< vector< vector< vector< vector< vector< uint8_t >
   > > > > > > > > > > > > > > > > > > > > from_j;

#ifdef EXCEPTIONS_ENABLED
   json j = json::parse( "[[[[[[[[[[[[[[[[[[[[[1]]]]]]]]]]]]]]]]]]]]]" );
   BOOST_REQUIRE_THROW( from_json( j, from_j ), depth_violation );
#endif
}
*/

BOOST_AUTO_TEST_SUITE_END()
#endif
