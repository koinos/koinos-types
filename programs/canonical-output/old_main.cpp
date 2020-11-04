#include <cstdlib>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>

#include <koinos/pack/rt/binary.hpp>
#include <koinos/pack/rt/json.hpp>
#include <koinos/pack/classes.hpp>

#include <boost/program_options.hpp>

#include <nlohmann/json.hpp>

using namespace koinos::types;
using namespace boost::program_options;
using nlohmann::json;

template< class T > void append( T&& t, json& arr, std::ofstream& bin )
{
   json j;
   koinos::pack::to_json( j, std::forward< T >( t ) );
   arr.push_back( std::move( j ) );
   koinos::pack::to_binary( bin, std::forward< T >( t ) );
}

int main( int argc, char** argv )
{
   using namespace std::string_literals;

   options_description desc{ "Options" };
   desc.add_options()
      ("help,h", "Help screen")
      ("binary,b", value< std::string >()->default_value( "types.bin" ),  "The binary output file")
      ("json,j",   value< std::string >()->default_value( "types.json" ), "The JSON output file");

   variables_map vm;
   store( parse_command_line( argc, argv, desc ), vm );
   notify( vm );

   if ( vm.count( "help" ) )
   {
      std::cout << desc << std::endl;
      exit( EXIT_SUCCESS );
   }

   std::ofstream bin_out, json_out;
   json arr = json::array();

   bin_out.open( vm[ "binary" ].as< std::string >(), std::ios::binary );
   json_out.open( vm[ "json" ].as< std::string >() );

   append( unused_extensions_type{}, arr, bin_out );
   append( protocol::reserved_operation{}, arr, bin_out );
   append( protocol::nop_operation{}, arr, bin_out );
   append( protocol::create_system_contract_operation {
      .contract_id = contract_id_type{ '1' },
      .bytecode = koinos::pack::to_variable_blob( "bytecode"s )
   }, arr, bin_out );
   append( system::system_call_target_reserved{}, arr, bin_out );
   append( system::contract_call_bundle {
      .contract_id = contract_id_type{ '2' },
      .entry_point = 400
   }, arr, bin_out );
   append( protocol::set_system_call_operation {
      .call_id = 300,
      .target = system::system_call_target{}
   }, arr, bin_out );
   append( protocol::contract_call_operation {
      .contract_id = contract_id_type{ '3' },
      .entry_point = 500,
      .args = koinos::pack::to_variable_blob( "arguments"s )
   }, arr, bin_out );
   append( protocol::active_transaction_data{}, arr, bin_out );
   append( protocol::passive_transaction_data{}, arr, bin_out );
   append( protocol::transaction {
      .active_data = protocol::active_transaction_data{},
      .passive_data = protocol::passive_transaction_data{},
      .signature_data = koinos::pack::to_variable_blob( "signature"s ),
      .operations = std::vector< protocol::operation > {
         protocol::reserved_operation{},
         protocol::nop_operation{},
         protocol::create_system_contract_operation {
            .contract_id = contract_id_type{ '4' },
            .bytecode = koinos::pack::to_variable_blob( "bytecode"s )
         }
      }
   }, arr, bin_out );
   append( protocol::active_block_data {
      .header_hashes = multihash_vector {
         .id = 70,
         .digests = std::vector< variable_blob > {
            koinos::pack::to_variable_blob( "digestA"s ),
            koinos::pack::to_variable_blob( "digestB"s ),
            koinos::pack::to_variable_blob( "digestC"s ),
         }
      },
      .height = block_height_type{ 20 },
      .timestamp = timestamp_type{ 30 }
   }, arr, bin_out );
   append( protocol::passive_block_data{}, arr, bin_out );
   append( protocol::block {
      .active_data = protocol::active_block_data {
         .header_hashes = multihash_vector {
            .id = 70,
            .digests = std::vector< variable_blob > {
               koinos::pack::to_variable_blob( "digestA"s ),
               koinos::pack::to_variable_blob( "digestB"s ),
               koinos::pack::to_variable_blob( "digestC"s ),
            }
         },
         .height = block_height_type{ 30 },
         .timestamp = timestamp_type{ 40 }
      },
      .passive_data = protocol::passive_block_data{},
      .signature_data = koinos::pack::to_variable_blob( "signature"s ),
      .transactions = std::vector< protocol::opaque_transaction > {
         protocol::opaque_transaction{},
         protocol::opaque_transaction{}
      }
   }, arr, bin_out );
   append( rpc::reserved_query_params{}, arr, bin_out );
   append( rpc::get_head_info_params{}, arr, bin_out );
   append( rpc::reserved_query_result{}, arr, bin_out );
   append( rpc::query_error {
      .error_text = koinos::pack::to_variable_blob( "failure"s )
   }, arr, bin_out );
   append( system::head_info {
      .id = multihash {
         .id = 50,
         .digest = koinos::pack::to_variable_blob( "50-digest"s )
      },
      .height = block_height_type{ 100 }
   }, arr, bin_out );
   append( rpc::block_topology {
      .id = multihash {
         .id = 50,
         .digest = koinos::pack::to_variable_blob( "50-digest"s )
      },
      .height = block_height_type{ 200 },
      .previous = multihash {
         .id = 60,
         .digest = koinos::pack::to_variable_blob( "60-digest"s )
      }
   }, arr, bin_out );
   append( rpc::reserved_submission{}, arr, bin_out );
   append( rpc::block_submission {
      .topology = rpc::block_topology {
         .id = multihash {
            .id = 60,
            .digest = koinos::pack::to_variable_blob( "60-digest"s )
         },
         .height = block_height_type{ 300 },
         .previous = multihash {
            .id = 70,
            .digest = koinos::pack::to_variable_blob( "70-digest"s )
         }
      },
      .block = protocol::block {
         .active_data = protocol::active_block_data {
            .header_hashes = multihash_vector {
               .id = 80,
               .digests = std::vector< variable_blob > {
                  koinos::pack::to_variable_blob( "digest1"s ),
                  koinos::pack::to_variable_blob( "digest2"s ),
                  koinos::pack::to_variable_blob( "digest3"s ),
               }
            },
            .height = block_height_type { 700 },
            .timestamp = timestamp_type { 1000 }

         },
         .passive_data = protocol::passive_block_data {},
         .signature_data = koinos::pack::to_variable_blob( "signature"s ),
         .transactions = std::vector< protocol::opaque_transaction >{
            protocol::transaction {
               .active_data = protocol::active_transaction_data{},
               .passive_data = protocol::passive_transaction_data{},
               .signature_data = koinos::pack::to_variable_blob( "tx-signature"s ),
               .operations = std::vector< protocol::operation > {
                  protocol::nop_operation{},
                  protocol::create_system_contract_operation {
                     .contract_id = contract_id_type{ '5' },
                     .bytecode = koinos::pack::to_variable_blob( "contract-bytecode"s )
                  },
                  protocol::reserved_operation{}
               }
            }
         }
      },
      .verify_passive_data = true,
      .verify_block_signature = true,
      .verify_transaction_signatures = true,
   }, arr, bin_out );
   append( rpc::transaction_submission {
      .active_bytes = koinos::pack::to_variable_blob( protocol::active_transaction_data {} ),
      .passive_bytes = koinos::pack::to_variable_blob( protocol::passive_transaction_data {} )
   }, arr, bin_out );
   append( rpc::reserved_submission_result{}, arr, bin_out );
   append( rpc::block_submission_result{}, arr, bin_out );
   append( rpc::transaction_submission_result{}, arr, bin_out );
   append( rpc::submission_error_result {
      .error_text = koinos::pack::to_variable_blob( "failure text"s )
   }, arr, bin_out );
   append( thunks::void_type{}, arr, bin_out );
   append( thunks::prints_args {
      .message = "prints arguments"
   }, arr, bin_out );
   append( thunks::verify_block_sig_args {
      .sig_data = koinos::pack::to_variable_blob( "signature data"s ),
      .digest = multihash {
         .id = 800,
         .digest = koinos::pack::to_variable_blob( "digest"s )
      }
   }, arr, bin_out );
   append( thunks::verify_merkle_root_args {
      .root = multihash {
         .id = 900,
         .digest = koinos::pack::to_variable_blob( "root"s )
      },
      .hashes = std::vector< multihash >{
         multihash {
            .id = 1000,
            .digest = koinos::pack::to_variable_blob( "hash1"s )
         },
         multihash {
            .id = 1100,
            .digest = koinos::pack::to_variable_blob( "hash2"s )
         }
      }
   }, arr, bin_out );
   append( thunks::apply_block_args {
      .block = protocol::block {
         .active_data = protocol::active_block_data {
         .header_hashes = multihash_vector {
            .id = 70,
            .digests = std::vector< variable_blob > {
               koinos::pack::to_variable_blob( "digestA"s ),
               koinos::pack::to_variable_blob( "digestB"s ),
               koinos::pack::to_variable_blob( "digestC"s ),
            }
         },
         .height = block_height_type{ 30 },
         .timestamp = timestamp_type{ 40 }
         },
         .passive_data = protocol::passive_block_data{},
         .signature_data = koinos::pack::to_variable_blob( "signature"s ),
         .transactions = std::vector< protocol::opaque_transaction > {
            protocol::opaque_transaction{},
            protocol::opaque_transaction{}
         }
      },
      .enable_check_passive_data = true,
      .enable_check_block_signature = true,
      .enable_check_transaction_signatures = true
   }, arr, bin_out );
   append( thunks:: apply_transaction_args {
      .trx = protocol::transaction {
         .active_data = protocol::active_transaction_data{},
         .passive_data = protocol::passive_transaction_data{},
         .signature_data = koinos::pack::to_variable_blob( "transaction signature"s ),
         .operations = std::vector< protocol::operation > {
            protocol::reserved_operation{},
            protocol::nop_operation{},
            protocol::create_system_contract_operation {
               .contract_id = contract_id_type{ '5' },
               .bytecode = koinos::pack::to_variable_blob( "contract bytecode"s )
            }
         }
      }
   }, arr, bin_out );
   append( thunks::apply_upload_contract_operation_args {
      .op = protocol::create_system_contract_operation {
         .contract_id = contract_id_type{ '1', '2', '3' },
         .bytecode = koinos::pack::to_variable_blob( "contract bytecode"s )
      }
   }, arr, bin_out );
   append( thunks::apply_reserved_operation_args {
      .op = protocol::reserved_operation{}
   }, arr, bin_out );
   append( thunks::apply_execute_contract_operation_args {
      .op = protocol::contract_call_operation {
         .contract_id = contract_id_type{ '4', '5', '6' },
         .entry_point = 7,
         .args = koinos::pack::to_variable_blob( "contract arguments"s )
      }
   }, arr, bin_out );
   append( thunks::apply_set_system_call_operation_args {
      .op = protocol::set_system_call_operation {
         .call_id = 8,
         .target = system::thunk_id_type{ 9 }
      }
   }, arr, bin_out );
   append( thunks::db_put_object_args {
      .space = 100,
      .key = 200,
      .obj = koinos::pack::to_variable_blob( "object"s )
   }, arr, bin_out );
   append( thunks::db_get_object_args {
      .space = 100,
      .key = 200,
      .object_size_hint = 300
   }, arr, bin_out );
   append( thunks::execute_contract_args {
      .contract_id = contract_id_type{ '7', '8', '9' },
      .entry_point = 10,
      .args = koinos::pack::to_variable_blob( "contract argument"s )
   }, arr, bin_out );
   append( thunks::set_contract_return_args {
      .ret = koinos::pack::to_variable_blob( "return"s )
   }, arr, bin_out );
   append( thunks::exit_contract_args {
      .exit_code = EXIT_SUCCESS
   }, arr, bin_out );
   append( thunks::hash_args {
      .code = 120,
      .obj = koinos::pack::to_variable_blob( "object"s ),
      .size = 7
   }, arr, bin_out );
   append( system::block_part {
      .active_data = koinos::pack::to_variable_blob( "active"s ),
      .passive_data = koinos::pack::to_variable_blob( "passive"s ),
      .sig_data = koinos::pack::to_variable_blob( "signature"s )
   }, arr, bin_out );

   json_out << arr;

   bin_out.close();
   json_out.close();

   return EXIT_SUCCESS;
}
