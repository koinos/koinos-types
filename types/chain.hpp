namespace koinos { namespace chain {

enum class privilege : uint8
{
   kernel_mode,
   user_mode
};

struct head_info
{
   block_topology    head_topology;
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
   boolean                                     check_passive_data;
   boolean                                     check_block_signature;
   boolean                                     check_transaction_signatures;
};

typedef void_type apply_block_return;

struct apply_transaction_args
{
   protocol::transaction transaction;
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
   protocol::call_contract_operation op;
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

typedef void_type get_entry_point_args;

typedef uint32 get_entry_point_return;

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

struct recover_public_key_args
{
   variable_blob signature_data;
   multihash     digest;
};

typedef account_type recover_public_key_return;

struct get_transaction_payer_args
{
   protocol::transaction transaction;
};

typedef account_type get_transaction_payer_return;

struct get_max_account_resources_args
{
   account_type account;
};

typedef int128 get_max_account_resources_return;

struct get_transaction_resource_limit_args
{
   protocol::transaction transaction;
};

typedef int128 get_transaction_resource_limit_return;

struct get_last_irreversible_block_args {};

typedef block_height_type get_last_irreversible_block_return;

struct get_caller_args {};

struct get_caller_return
{
   account_type caller;
   privilege    caller_privilege;
};

struct require_authority_args
{
   account_type account;
};

typedef void_type require_authority_return;

struct get_transaction_signature_args {};

typedef variable_blob get_transaction_signature_return;

typedef void_type get_contract_id_args;

typedef contract_id_type get_contract_id_return;

typedef void_type get_head_block_time_args;

typedef timestamp_type get_head_block_time_return;

/**
 * Represents limits on a single resource.
 */
struct resource_limit
{
   /**
    * Max number of resources that can be used per transaction.
    */
   int128 max_usage;

   /**
    * Minimum number of resources that can be used per transaction.
    *
    * Usually will be zero or negative.  Setting a negative value here
    * will allow transactions to get credit by freeing resources.
    */
   int128 min_usage;

   /**
    * Price of the resource in mana.
    *
    * If zero, consuming the resource does not consume mana.
    * If negative, consuming the resource gives mana.
    */
   int64 price;

   /**
    * Price is multiplied by quantity, then divided by 10**price_shift.
    */
   uint8 price_shift;
};

/**
 * A numerical identifier for resources.
 *
 * We can't call it resource_id because _id suffix is special-cased in the code generator.
 * Resource 0 is always mana.
 */
typedef uint16 resource_num;

struct set_resource_limits_args
{
   std::vector< resource_limit > resource_limits;
};

typedef void_type set_resource_limits_return;

struct use_resource_args
{
   resource_num resource;
   int128 amount;
};

typedef void_type use_resource_return;

} } // koinos::chain
