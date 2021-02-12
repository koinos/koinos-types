namespace koinos { namespace rpc { namespace chain {

struct chain_reserved_request {};

struct submit_block_request
{
   block_topology                             topology;

   /**
    * block_parts[0].active_data     -> active_block_data
    * block_parts[0].passive_data    -> passive_block_data
    * block_parts[0].sig_data        -> sig_block_data
    *
    * block_parts[1..n].active_data  -> active_transaction_data (transaction_type)
    * block_parts[1..n].passive_data -> passive_transaction_data
    * block_parts[1..n].sig_data     -> sig_transaction_data
    */
   protocol::block                     block;

   boolean                             verify_passive_data;
   boolean                             verify_block_signature;
   boolean                             verify_transaction_signatures;
};

struct submit_transaction_request
{
   opaque< protocol::active_transaction_data >   active_data;
   opaque< protocol::passive_transaction_data >  passive_data;
};

struct get_head_info_request {};

struct get_chain_id_request {};

typedef std::variant<
   chain_reserved_request,
   submit_block_request,
   submit_transaction_request,
   get_head_info_request,
   get_chain_id_request > chain_rpc_request;

struct chain_reserved_response {};

struct chain_error_response
{
   std::string error_text;
};

struct submit_block_response {};

struct submit_transaction_response {};

struct get_head_info_response
{
   multihash              id;
   block_height_type      height;
};

struct get_chain_id_response
{
   multihash chain_id;
};

typedef std::variant<
   chain_reserved_response,
   chain_error_response,
   submit_block_response,
   submit_transaction_response,
   get_head_info_response,
   get_chain_id_response > chain_rpc_response;

} } } // koinos::rpc::chain
