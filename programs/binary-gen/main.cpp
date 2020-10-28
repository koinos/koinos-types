#include <cstdlib>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

#include <koinos/pack/rt/binary.hpp>
#include <koinos/pack/rt/json.hpp>
#include <koinos/pack/classes.hpp>

#include <nlohmann/json.hpp>

using namespace koinos::types;
using nlohmann::json;

static std::ofstream bin_out, json_out;

template< class T > void write( T&& t )
{
   json j;
   koinos::pack::to_json( j, t );
   json_out << j;

   koinos::pack::to_binary( bin_out, t );
}

int main( int argc, char** argv )
{
   using namespace std::string_literals;

   bin_out.open( "types.bin", std::ios::binary );
   json_out.open( "types.json" );

   write( unused_extensions_type{} );
   write( protocol::reserved_operation{} );
   write( protocol::nop_operation{} );
   write( protocol::create_system_contract_operation {
      .contract_id = contract_id_type{ '1' },
      .bytecode = koinos::pack::to_variable_blob( "bytecode"s )
   } );
   write( system::system_call_target_reserved{} );
   write( system::contract_call_bundle {
      .contract_id = contract_id_type{ '2' },
      .entry_point = 400
   } );
   write( protocol::set_system_call_operation {
      .call_id = 300,
      .target = system::system_call_target{}
   } );
   write( protocol::contract_call_operation {
      .contract_id = contract_id_type{ '3' },
      .entry_point = 500,
      .args = koinos::pack::to_variable_blob( "arguments"s )
   } );
   write( protocol::active_transaction_data{} );
   write( protocol::passive_transaction_data{} );
   write( protocol::transaction {
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
   } );
   write( protocol::active_block_data {
      .header_hashes = 10,
      .height = block_height_type{ 20 },
      .timestamp = timestamp_type{ 30 }
   } );
   write( protocol::passive_block_data{} );
   write( protocol::block {
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
   } );
   write( rpc::reserved_query_params{} );
   write( rpc::get_head_info_params{} );
   write( rpc::reserved_query_result{} );
   write( rpc::query_error {
      .error_text = koinos::pack::to_variable_blob( "failure"s )
   } );
   write( system::head_info {
      .id = multihash {
         .id = 50,
         .digest = koinos::pack::to_variable_blob( "50-digest"s )
      },
      .height = block_height_type{ 100 }
   } );
   write( rpc::block_topology {
      .id = multihash {
         .id = 50,
         .digest = koinos::pack::to_variable_blob( "50-digest"s )
      },
      .height = block_height_type{ 200 },
      .previous = multihash {
         .id = 60,
         .digest = koinos::pack::to_variable_blob( "60-digest"s )
      }
   } );
   write( rpc::reserved_submission{} );
   write( rpc::block_submission {
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
   } );
   write( rpc::transaction_submission {
      .active_bytes = koinos::pack::to_variable_blob( protocol::active_transaction_data {} ),
      .passive_bytes = koinos::pack::to_variable_blob( protocol::passive_transaction_data {} )
   } );
   write( rpc::reserved_submission_result{} );
   write( rpc::block_submission_result{} );
   write( rpc::transaction_submission_result{} );
   write( rpc::submission_error_result {
      .error_text = koinos::pack::to_variable_blob( "failure text"s )
   } );
   write( thunks::void_type{} );
   write( thunks::prints_args {
      .message = "prints arguments"
   } );
   write( thunks::verify_block_sig_args {
      .sig_data = koinos::pack::to_variable_blob( "signature data"s ),
      .digest = multihash {
         .id = 800,
         .digest = koinos::pack::to_variable_blob( "digest"s )
      }
   } );
   write( thunks::verify_merkle_root_args {
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
   } );
   write( thunks::apply_block_args {
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
   } );
   write( thunks:: apply_transaction_args {
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
   } );
   write( thunks::apply_upload_contract_operation_args {
      .op = protocol::create_system_contract_operation {
         .contract_id = contract_id_type{ '1', '2', '3' },
         .bytecode = koinos::pack::to_variable_blob( "contract bytecode"s )
      }
   } );
   write( thunks::apply_reserved_operation_args {
      .op = protocol::reserved_operation{}
   } );
   write( thunks::apply_execute_contract_operation_args {
      .op = protocol::contract_call_operation {
         .contract_id = contract_id_type{ '4', '5', '6' },
         .entry_point = 7,
         .args = koinos::pack::to_variable_blob( "contract arguments"s )
      }
   } );
   write( thunks::apply_set_system_call_operation_args {
      .op = protocol::set_system_call_operation {
         .call_id = 8,
         .target = system::thunk_id_type{ 9 }
      }
   } );
   write( thunks::db_put_object_args {
      .space = 100,
      .key = 200,
      .obj = koinos::pack::to_variable_blob( "object"s )
   } );
   write( thunks::db_get_object_args {
      .space = 100,
      .key = 200,
      .object_size_hint = 300
   } );
   write( thunks::execute_contract_args {
      .contract_id = contract_id_type{ '7', '8', '9' },
      .entry_point = 10,
      .args = koinos::pack::to_variable_blob( "contract argument"s )
   } );
   write( thunks::set_contract_return_args {
      .ret = koinos::pack::to_variable_blob( "return"s )
   } );
   write( thunks::exit_contract_args {
      .exit_code = EXIT_SUCCESS
   } );
   write( thunks::hash_args {
      .code = 120,
      .obj = koinos::pack::to_variable_blob( "object"s ),
      .size = 7
   } );
   write( system::block_part {
      .active_data = koinos::pack::to_variable_blob( "active"s ),
      .passive_data = koinos::pack::to_variable_blob( "passive"s ),
      .sig_data = koinos::pack::to_variable_blob( "signature"s )
   } );

   bin_out.close();
   json_out.close();

   return EXIT_SUCCESS;
}
