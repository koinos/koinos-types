namespace koinos { namespace rpc { namespace transaction_store {

struct transaction_store_reserved_request {};

struct transaction_store_reserved_response {};

struct get_transactions_by_id_request
{
   std::vector< multihash > transaction_ids;
};

typedef std::optional< transaction_store::transaction_record > opt_transaction_record;

struct get_transactions_by_id_response
{
   std::vector< opt_transaction_record > transactions;
};

struct transaction_store_error_response
{
   std::string error_text;
   std::string error_data;
};

typedef std::variant<
   transaction_store_reserved_request,
   get_transactions_by_id_request
   > transaction_store_request;

typedef std::variant<
   transaction_store_reserved_response,
   transaction_store_error_response,
   get_transactions_by_id_response
   > transaction_store_response;

} } }
