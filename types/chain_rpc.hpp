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

struct get_pending_transactions_request
{
   multihash start;
   uint64    limit;
};

struct get_fork_heads_request {};

typedef std::variant<
   chain_reserved_request,
   submit_block_request,
   submit_transaction_request,
   get_head_info_request,
   get_chain_id_request,
   get_pending_transactions_request,
   get_fork_heads_request > chain_rpc_request;

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
   multihash         id;
   multihash         previous_id;
   block_height_type height;
   block_height_type last_irreversible_height;
};

struct get_chain_id_response
{
   multihash chain_id;
};

struct get_pending_transactions_response
{
   std::vector< protocol::transaction > transactions;
};

struct get_fork_heads_response
{
   std::vector< block_topology > fork_heads;
   block_topology                last_irreversible_block;
};

typedef std::variant<
   chain_reserved_response,
   chain_error_response,
   submit_block_response,
   submit_transaction_response,
   get_head_info_response,
   get_chain_id_response,
   get_pending_transactions_response,
   get_fork_heads_response > chain_rpc_response;

} } } // koinos::rpc::chain
