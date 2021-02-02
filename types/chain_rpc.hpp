namespace koinos { namespace rpc { namespace chain {

struct reserved_rpc_params {};

struct block_topology
{
   types::multihash                      id;
   types::block_height_type              height;
   types::multihash                      previous;
};

struct submit_block_params
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
   types::protocol::block                     block;

   types::boolean                             verify_passive_data;
   types::boolean                             verify_block_signature;
   types::boolean                             verify_transaction_signatures;
};

struct submit_transaction_params
{
   types::variable_blob                       active_bytes;
   types::variable_blob                       passive_bytes;
};

struct get_head_info_params {};

struct get_chain_id_params {};

typedef std::variant<
   reserved_rpc_params,
   submit_block_params,
   submit_transaction_params,
   get_head_info_params,
   get_chain_id_params > koinosd_rpc_params;

struct reserved_rpc_result {};

struct rpc_error
{
   std::string error_text;
};

struct submit_block_result {};

struct submit_transaction_result {};

typedef types::system::head_info get_head_info_result;

struct get_chain_id_result
{
   types::multihash chain_id;
};

typedef std::variant<
   reserved_rpc_result,
   rpc_error,
   submit_block_result,
   submit_transaction_result,
   get_head_info_result,
   get_chain_id_result > koinosd_rpc_result;

} } } // koinos::rpc::koinosd
