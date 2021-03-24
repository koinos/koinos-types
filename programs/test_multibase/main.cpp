
#include <koinos/pack/rt/util/multibase.hpp>

#include <iostream>

int main(int argc, char** argv, char** envp)
{
   std::string msg = "hello";
   std::vector<char> dest;

   koinos::pack::util::encode_multibase(msg.c_str(), msg.size(), dest, 'z');
   std::string result(dest.begin(), dest.end());

   std::cout << "result: " << result << std::endl;

   return 0;
}
