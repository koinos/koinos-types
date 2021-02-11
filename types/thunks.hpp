
namespace koinos { namespace thunk {

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

typedef boolean verify_block_sig_ret;

struct verify_merkle_root_args
{
   multihash                                  root;
   std::vector< multihash >                   hashes;
};

typedef boolean verify_merkle_root_ret;

struct apply_block_args
{
   protocol::block                             block;
   boolean                                     enable_check_passive_data;
   boolean                                     enable_check_block_signature;
   boolean                                     enable_check_transaction_signatures;
};

typedef void_type apply_block_ret;

struct apply_transaction_args
{
   opaque< protocol::transaction >      trx;
};

typedef void_type apply_transaction_ret;

struct apply_upload_contract_operation_args
{
   protocol::create_system_contract_operation op;
};

typedef void_type apply_upload_contract_operation_ret;

struct apply_reserved_operation_args
{
   protocol::reserved_operation op;
};

typedef void_type apply_reserved_operation_ret;

struct apply_execute_contract_operation_args
{
   protocol::contract_call_operation op;
};

typedef void_type apply_execute_contract_operation_ret;

struct apply_set_system_call_operation_args
{
   protocol::set_system_call_operation op;
};

typedef void_type apply_set_system_call_operation_ret;

struct db_put_object_args
{
   uint256       space;
   uint256       key;
   variable_blob obj;
};

typedef boolean db_put_object_ret;

struct db_get_object_args
{
   uint256 space;
   uint256 key;
   int32   object_size_hint;
};

typedef variable_blob db_get_object_ret;

typedef db_get_object_args db_get_next_object_args;

typedef db_get_object_ret db_get_next_object_ret;

typedef db_get_object_args db_get_prev_object_args;

typedef db_get_object_ret db_get_prev_object_ret;

struct execute_contract_args
{
   contract_id_type contract_id;
   uint32           entry_point;
   variable_blob    args;
};

typedef variable_blob execute_contract_ret;

typedef void_type get_contract_args_size_args;

typedef uint32 get_contract_args_size_ret;

typedef void_type get_contract_args_args;

typedef variable_blob get_contract_args_ret;

struct set_contract_return_args
{
   variable_blob ret;
};

typedef void_type set_contract_return_ret;

struct exit_contract_args
{
   uint8   exit_code;
};

typedef void_type exit_contract_ret;

typedef void_type get_head_info_args;

typedef system::head_info get_head_info_ret;

struct hash_args
{
   uint64        code;
   variable_blob obj;
   uint64        size;
};

typedef multihash hash_ret;

} } // koinos::thunk
