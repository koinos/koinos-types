#ifdef JSON_ENABLED
#include <boost/test/unit_test.hpp>

#include <koinos/tests/pack_fixture.hpp>

#include <koinos/pack/rt/json.hpp>

#include <cstdint>
#include <iomanip>
#include <iostream>
#include <string>
#include <vector>

std::vector<char> string_to_vector( const std::string& s )
{
   std::vector<char> result;
   for( size_t i=0; i<s.size(); i++ )
   {
      result.push_back(s[i]);
   }
   return result;
}

void dump_bytes( const std::vector< char >& s )
{
   std::cerr << std::hex << std::setfill('0') << std::setw(4) << s.size() << ",";
   for( size_t j=0; j<s.size(); j++ )
   {
      std::cerr << std::hex << std::setfill('0') << std::setw(2) << uint32_t(uint8_t(s[j]));
      if( j+1 < s.size() )
         std::cerr << ",";
   }
   std::cerr << std::endl;
}

void dump_str( const std::vector< char >& s )
{
   std::cerr << "\"" << std::string( s.begin(), s.end() ) << "\"" << std::endl;
}


BOOST_AUTO_TEST_CASE( multibase_test )
{
   std::vector< std::vector< std::string > > test_cases = {
      {"f", "z", "m", "u"},                        // n.b. Python base58 fails this case
      {"f61", "z2g", "mYQ", "uYQ"},
      {"f6161", "z8Qp", "mYWE", "uYWE"},
      {"f616161", "zZi88", "mYWFh", "uYWFh"},
      {"f61616161", "z3VNWTa", "mYWFhYQ", "uYWFhYQ"},
      {"f6161616161", "zBzDw2JL", "mYWFhYWE", "uYWFhYWE"},
      {"f616161616161", "zqVa5SjWY", "mYWFhYWFh", "uYWFhYWFh"},
      {"f61616161616161", "z4h36zcadPW", "mYWFhYWFhYQ", "uYWFhYWFhYQ"},
      {"f6161616161616161", "zHHiHTJ3RcLg", "mYWFhYWFhYWE", "uYWFhYWFhYWE"},
      {"f616161616161616161", "z2EtmDd4Dhcyrp", "mYWFhYWFhYWFh", "uYWFhYWFhYWFh"},
      {"f00", "z1", "mAA", "uAA"},
      {"f0000", "z1", "mAAA", "uAAA"},
      {"f000000", "z1", "mAAAA", "uAAAA"},
      {"f00000000", "z1", "mAAAAAA", "uAAAAAA"},
      {"f0000000000", "z1", "mAAAAAAA", "uAAAAAAA"},
      {"f000000000000", "z1", "mAAAAAAAA", "uAAAAAAAA"},
      {"f00000000000000", "z1", "mAAAAAAAAAA", "uAAAAAAAAAA"},
      {"f0000000000000000", "z1", "mAAAAAAAAAAA", "uAAAAAAAAAAA"},
      {"f000000000000000000", "z1", "mAAAAAAAAAAAA", "uAAAAAAAAAAAA"},
      {"f7468697320697320612074657374", "zjo91waLQA1NNeBmZKUF", "mdGhpcyBpcyBhIHRlc3Q", "udGhpcyBpcyBhIHRlc3Q"},
      {"f6d756c74696261736520656e636f64696e67", "z5YKdSgATEwQMG3eWyqWvhjEhL", "mbXVsdGliYXNlIGVuY29kaW5n", "ubXVsdGliYXNlIGVuY29kaW5n"},
      {"fe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", "zGKot5hBsd81kMupNCXHaqbhv3huEbxAFMLnpcX2hniwn", "m47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU", "u47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU"}
   };

   for( size_t case_num=0; case_num < test_cases.size(); case_num++ )
   {
      std::string hex_str = test_cases[case_num][0];
      std::vector<char> bin;
      koinos::pack::util::decode_multibase( hex_str.c_str(), hex_str.size(), bin );

      for( size_t i=0; i<test_cases[case_num].size(); i++ )
      {
         try
         {
            std::string current = test_cases[case_num][i];

            //
            // Two checks performed here:
            //
            // - Decode the string and make sure it matches bin (which was decoded from the first string)
            // - Encode bin and make sure it matches current
            //

            std::vector<char> decoded;
            koinos::pack::util::decode_multibase( current.c_str(), current.size(), decoded );
            if( decoded != bin )
            {
               std::cerr << "Dumping mismatch:  decoded != bin" << std::endl;
               std::cerr << "input       is \"" << current << "\"" << std::endl;
               std::cerr << "decoded     is ";
               dump_str(decoded);
               dump_bytes(decoded);
               std::cerr << "bin         is ";
               dump_str(bin);
               dump_bytes(bin);
            }
            BOOST_CHECK( decoded == bin );

            std::vector<char> encoded;
            koinos::pack::util::encode_multibase( bin.data(), bin.size(), encoded, current[0] );
            std::vector<char> current_bin = string_to_vector( current );
            if( encoded != current_bin )
            {
               std::cerr << "Dumping mismatch:  encoded != current_bin with encoding " << current[0] << std::endl;
               std::cerr << "input       is ";
               dump_str(bin);
               dump_bytes(bin);
               std::cerr << "encoded     is ";
               dump_str(encoded);
               dump_bytes(encoded);
               std::cerr << "current_bin is ";
               dump_str(current_bin);
               dump_bytes(current_bin);
            }
            BOOST_CHECK( encoded == current_bin );
         }
         catch( const koinos::pack::base_decode_error& e )
         {
            std::cerr << e.what() << std::endl;
            BOOST_CHECK( false );
         }
      }
   }

   // koinos::pack::util::encode_multibase(msg.c_str(), msg.size(), dest, 'z');
}
#endif