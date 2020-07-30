
namespace koinos { namespace types { namespace thunks {

struct void_type {};

struct prints_args
{
   std::string message;
};

typedef void_type prints_ret;

struct verify_block_sig_args
{
   variable_blob                              sig_data;
   multihash                                  digest;
};

typedef types::boolean verify_block_sig_ret;

struct verify_merkle_root_args
{
   multihash                                  root;
   std::vector< multihash >                   hashes;
};

typedef types::boolean verify_merkle_root_ret;

struct apply_block_args
{
   protocol::block                             block;
   types::boolean                              enable_check_passive_data;
   types::boolean                              enable_check_block_signature;
   types::boolean                              enable_check_transaction_signatures;
};

typedef void_type apply_block_ret;

struct apply_transaction_args
{
   opaque< types::protocol::transaction >      trx;
};

typedef void_type apply_transaction_ret;

struct apply_upload_contract_operation_args
{
   types::protocol::create_system_contract_operation op;
};

typedef void_type apply_upload_contract_operation_ret;

struct apply_reserved_operation_args
{
   types::protocol::reserved_operation op;
};

typedef void_type apply_reserved_operation_ret;

struct apply_execute_contract_operation_args
{
   types::protocol::contract_call_operation op;
};

typedef void_type apply_execute_contract_operation_ret;

struct apply_set_system_call_operation_args
{
   protocol::set_system_call_operation op;
};

typedef void_type apply_set_system_call_operation_ret;

struct db_put_object_args
{
   types::uint256       space;
   types::uint256       key;
   types::variable_blob obj;
};

typedef types::boolean db_put_object_ret;

struct db_get_object_args
{
   types::uint256 space;
   types::uint256 key;
   types::int32   object_size_hint;
};

typedef types::variable_blob db_get_object_ret;

typedef db_get_object_args db_get_next_object_args;

typedef db_get_object_ret db_get_next_object_ret;

typedef db_get_object_args db_get_prev_object_args;

typedef db_get_object_ret db_get_prev_object_ret;

struct execute_contract_args
{
   types::contract_id_type contract_id;
   types::uint32           entry_point;
   types::variable_blob    args;
};

typedef types::variable_blob execute_contract_ret;

typedef void_type get_contract_args_size_args;

typedef types::uint32 get_contract_args_size_ret;

typedef void_type get_contract_args_args;

typedef types::variable_blob get_contract_args_ret;

struct set_contract_return_args
{
   types::variable_blob ret;
};

typedef void_type set_contract_return_ret;

struct exit_contract_args
{
   types::uint8   exit_code;
};

typedef void_type exit_contract_ret;

typedef void_type get_head_info_args;

typedef system::head_info get_head_info_ret;

struct hash_args
{
   types::uint64        code;
   types::variable_blob obj;
   types::uint64        size;
};

typedef types::multihash hash_ret;

} } } // koinos::types::thunks
