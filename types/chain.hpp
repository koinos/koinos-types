namespace koinos { namespace chain {

// Use generate_ids.py to generate the system call id
enum class system_call_id : uint32
{
   prints = 0x9b229941,
   verify_block_header = 0x9625504e,
   apply_block = 0x94ab4ff5,
   apply_transaction = 0x9d8b55da,
   apply_reserved_operation = 0x9aa85924,
   apply_upload_contract_operation = 0x9e6ea937,
   apply_execute_contract_operation = 0x92184686,
   apply_set_system_call_operation = 0x9579a45c,
   db_put_object = 0x971ec7a2,
   db_get_object = 0x9766a8fb,
   db_get_next_object = 0x99a398e8,
   db_get_prev_object = 0x9bd3767c,
   execute_contract = 0x98c12cfe,
   get_contract_args_size = 0x9b0d8fd9,
   get_contract_args = 0x9fbba198,
   set_contract_return = 0x9f49cdea,
   exit_contract = 0x98df75b0,
   get_head_info = 0x956fb22d,
   hash = 0x99770e04,
   verify_block_signature = 0x8fb9a59b,
   verify_merkle_root = 0x996e24b9,
   get_transaction_payer = 0x86a87bf5,
   get_max_account_resources = 0x842c6c81,
   get_transaction_resource_limit = 0x9940f685,
   get_last_irreversible_block = 0x953d2e37,
   get_caller = 0x94176c5f,
   require_authority = 0x9491e528,
   get_transaction_signature = 0x9dc3ec34,
   set_user_mode = 0x9dd0a15a
};

// Use generate_ids.py to generate the thunk id
enum class thunk_id : uint32
{
   prints = 0x8f6df54d,
   verify_block_header = 0x8d425aac,
   apply_block = 0x8d6d31a8,
   apply_transaction = 0x8981b0df,
   apply_reserved_operation = 0x8b3c14f6,
   apply_upload_contract_operation = 0x8882a55e,
   apply_execute_contract_operation = 0x85e882eb,
   apply_set_system_call_operation = 0x86f92c8c,
   db_put_object = 0x82038de5,
   db_get_object = 0x8862a0d8,
   db_get_next_object = 0x86e45047,
   db_get_prev_object = 0x8d57e8fd,
   execute_contract = 0x8a43fe83,
   get_contract_args_size = 0x83378e86,
   get_contract_args = 0x8e189d86,
   set_contract_return = 0x86b86275,
   exit_contract = 0x81f61f9f,
   get_head_info = 0x89df34c4,
   hash = 0x8aaaf547,
   verify_block_signature = 0x9d1c3c89,
   verify_merkle_root = 0x8ed9ddcb,
   get_transaction_payer = 0x9db35086,
   get_max_account_resources = 0x90f14f8d,
   get_transaction_resource_limit = 0x8bdf81a1,
   get_last_irreversible_block = 0x80c3b893,
   get_caller = 0x82312501,
   require_authority = 0x8a06717d,
   get_transaction_signature = 0x83441b23,
   set_user_mode = 0x8e887829
};

struct head_info
{
   multihash         id;
   multihash         previous_id;
   block_height_type height;
   block_height_type last_irreversible_height;
};

typedef variable_blob account_type;

struct system_call_target_reserved {};

struct contract_call_bundle
{
   contract_id_type contract_id;
   uint32           entry_point;
};

typedef std::variant<
   system_call_target_reserved,
   thunk_id,
   contract_call_bundle
   > system_call_target;

struct void_type {};

struct prints_args
{
   std::string message;
};

typedef void_type prints_return;

struct verify_block_signature_args
{
   variable_blob signature_data;
   multihash     digest;
};

typedef boolean verify_block_signature_return;

struct verify_merkle_root_args
{
   multihash                                  root;
   std::vector< multihash >                   hashes;
};

typedef boolean verify_merkle_root_return;

struct apply_block_args
{
   protocol::block                             block;
   boolean                                     enable_check_passive_data;
   boolean                                     enable_check_block_signature;
   boolean                                     enable_check_transaction_signatures;
};

typedef void_type apply_block_return;

struct apply_transaction_args
{
   opaque< protocol::transaction >      trx;
};

typedef void_type apply_transaction_return;

struct apply_upload_contract_operation_args
{
   protocol::create_system_contract_operation op;
};

typedef void_type apply_upload_contract_operation_return;

struct apply_reserved_operation_args
{
   protocol::reserved_operation op;
};

typedef void_type apply_reserved_operation_return;

struct apply_execute_contract_operation_args
{
   protocol::contract_call_operation op;
};

typedef void_type apply_execute_contract_operation_return;

struct apply_set_system_call_operation_args
{
   protocol::set_system_call_operation op;
};

typedef void_type apply_set_system_call_operation_return;

struct db_put_object_args
{
   uint256       space;
   uint256       key;
   variable_blob obj;
};

typedef boolean db_put_object_return;

struct db_get_object_args
{
   uint256 space;
   uint256 key;
   int32   object_size_hint;
};

typedef variable_blob db_get_object_return;

typedef db_get_object_args db_get_next_object_args;

typedef db_get_object_return db_get_next_object_return;

typedef db_get_object_args db_get_prev_object_args;

typedef db_get_object_return db_get_prev_object_return;

struct execute_contract_args
{
   contract_id_type contract_id;
   uint32           entry_point;
   variable_blob    args;
};

typedef variable_blob execute_contract_return;

typedef void_type get_contract_args_size_args;

typedef uint32 get_contract_args_size_return;

typedef void_type get_contract_args_args;

typedef variable_blob get_contract_args_return;

struct set_contract_return_args
{
   variable_blob value;
};

typedef void_type set_contract_return_return;

struct exit_contract_args
{
   uint8   exit_code;
};

typedef void_type exit_contract_return;

typedef void_type get_head_info_args;

typedef head_info get_head_info_return;

struct hash_args
{
   uint64        code;
   variable_blob obj;
   uint64        size;
};

typedef multihash hash_return;

struct get_transaction_payer_args
{
   opaque< protocol::transaction > transaction;
};

typedef account_type get_transaction_payer_return;

struct get_max_account_resources_args
{
   account_type account;
};

typedef uint128 get_max_account_resources_return;

struct get_transaction_resource_limit_args
{
   opaque< protocol::transaction > transaction;
};

typedef uint128 get_transaction_resource_limit_return;

struct get_last_irreversible_block_args {};

typedef block_height_type get_last_irreversible_block_return;

struct get_caller_args {};

typedef account_type get_caller_return;

struct require_authority_args
{
   account_type account;
};

typedef void_type require_authority_return;

struct get_transaction_signature_args {};

typedef variable_blob get_transaction_signature_return;

struct set_user_mode_args{};

typedef void_type set_user_mode_return;

} } // koinos::chain
