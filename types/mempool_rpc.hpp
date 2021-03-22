namespace koinos { namespace rpc { namespace mempool {

struct mempool_reserved_request {};

struct check_pending_account_resources_request
{
   chain::account_type payer;
   uint128             max_payer_resources;
   uint128             trx_resource_limit;
};

struct get_pending_transactions_request
{
   uint64 limit;
};

typedef std::variant<
   mempool_reserved_request,
   check_pending_account_resources_request,
   get_pending_transactions_request > mempool_rpc_request;

struct mempool_reserved_response {};

struct mempool_error_response
{
   std::string error_text;
   std::string error_data;
};

struct check_pending_account_resources_response
{
   boolean success;
};

struct get_pending_transactions_response
{
   std::vector< protocol::transaction > transactions;
};

typedef std::variant<
   mempool_reserved_response,
   mempool_error_response,
   check_pending_account_resources_response,
   get_pending_transactions_response > mempool_rpc_response;

} } } // koinos::rpc::mempool
