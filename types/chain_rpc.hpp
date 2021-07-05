namespace koinos { namespace rpc { namespace chain {

struct chain_reserved_request {};

struct submit_block_request
{
   protocol::block block;
   boolean         verify_passive_data;
   boolean         verify_block_signature;
   boolean         verify_transaction_signatures;
};

struct submit_transaction_request
{
   protocol::transaction transaction;
   boolean               verify_passive_data;
   boolean               verify_transaction_signatures;
};

struct get_head_info_request {};

struct get_chain_id_request {};

struct get_fork_heads_request {};

struct read_contract_request
{
   contract_id_type contract_id;
   uint32           entry_point;
   variable_blob    args;
};

struct get_account_nonce_request
{
   protocol::account_type account;
};

typedef std::variant<
   chain_reserved_request,
   submit_block_request,
   submit_transaction_request,
   get_head_info_request,
   get_chain_id_request,
   get_fork_heads_request,
   read_contract_request,
   get_account_nonce_request > chain_rpc_request;

struct chain_reserved_response {};

struct chain_error_response
{
   std::string error_text;
   std::string error_data;
};

struct submit_block_response {};

struct submit_transaction_response {};

struct get_head_info_response
{
   block_topology    head_topology;
   block_height_type last_irreversible_height;
};

struct get_chain_id_response
{
   multihash chain_id;
};

struct get_fork_heads_response
{
   std::vector< block_topology > fork_heads;
   block_topology                last_irreversible_block;
};

struct read_contract_response
{
   variable_blob result;
   std::string   logs;
};

struct get_account_nonce_response
{
   uint64 nonce;
};

typedef std::variant<
   chain_reserved_response,
   chain_error_response,
   submit_block_response,
   submit_transaction_response,
   get_head_info_response,
   get_chain_id_response,
   get_fork_heads_response,
   read_contract_response,
   get_account_nonce_response > chain_rpc_response;

} } } // koinos::rpc::chain
