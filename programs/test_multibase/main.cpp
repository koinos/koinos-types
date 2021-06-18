
#include <koinos/pack/rt/util/multibase.hpp>

#include <nlohmann/json.hpp>

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

int main(int argc, char** argv, char** envp)
{
   while(true)
   {
      std::string line;
      std::getline(std::cin, line);

      nlohmann::json j = nlohmann::json::parse(line);

      std::string hex_str = j.at(0).get<std::string>();
      std::vector<char> bin;
      koinos::pack::util::decode_multibase( hex_str.c_str(), hex_str.size(), bin );

      for( size_t i=0; i<j.size(); i++ )
      {
         std::string current = j.at(i);

         //
         // Two checks performed here:
         //
         // - Decode the string and make sure it matches bin (which was decoded from the first string)
         // - Encode bin and make sure it matches current
         //

         std::vector<char> decoded;
         std::optional< std::string > err;
         try
         {
            koinos::pack::util::decode_multibase( current.c_str(), current.size(), decoded );
         }
         catch( const koinos::pack::base_decode_error& e )
         {
            err = e.what();
         }

         if( err.has_value() || (decoded != bin) )
         {
            std::cerr << "Dumping mismatch:  decoded != bin" << std::endl;
            std::cerr << "input       is \"" << current << "\"" << std::endl;
            if( err.has_value() )
            {
               std::cerr << "error was " << (*err) << std::endl;
            }
            else
            {
               std::cerr << "decoded     is ";
               dump_str(decoded);
               dump_bytes(decoded);
            }
            std::cerr << "bin         is ";
            dump_str(bin);
            dump_bytes(bin);

            return 1;
         }

         std::vector<char> encoded;
         std::vector<char> current_bin;
         bool ok = false;

         try
         {
            koinos::pack::util::encode_multibase( bin.data(), bin.size(), encoded, current[0] );
            current_bin = string_to_vector( current );
            ok = true;
         }
         catch( const koinos::pack::base_decode_error& e )
         {
            err = e.what();
         }

         if( encoded != current_bin )
         {
            std::cerr << "Dumping mismatch:  encoded != current_bin with encoding " << current[0] << std::endl;
            std::cerr << "input       is ";
            dump_str(bin);
            dump_bytes(bin);
            if( err.has_value() )
            {
               std::cerr << "error was " << (*err) << std::endl;
            }
            else
            {
               std::cerr << "encoded     is ";
               dump_str(encoded);
               dump_bytes(encoded);
            }
            std::cerr << "current_bin is ";
            dump_str(current_bin);
            dump_bytes(current_bin);
            return 1;
         }
      }
   }
   return 0;   // unreachable
}
