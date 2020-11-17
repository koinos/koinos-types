namespace koinos { namespace types { namespace rpc {

struct reserved_query_params {};

struct get_head_info_params {};

typedef std::variant<
   reserved_query_params,
   get_head_info_params > query_param_item;

typedef opaque< query_param_item > query_submission;

struct reserved_query_result {};

struct query_error
{
   variable_blob error_text;
};

typedef types::system::head_info get_head_info_result;

typedef std::variant<
   reserved_query_result,
   query_error,
   get_head_info_result > query_item_result;

typedef opaque< query_item_result > query_submission_result;

// Every block has a cryptographic ID.
// Check the claimed ID against the block content.

struct block_topology
{
   types::multihash                      id;
   types::block_height_type              height;
   types::multihash                      previous;
};

struct reserved_submission {};

struct block_submission
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
   protocol::block                            block;

   boolean                                    verify_passive_data;
   boolean                                    verify_block_signature;
   boolean                                    verify_transaction_signatures;
};

struct transaction_submission
{
   types::variable_blob                       active_bytes;
   types::variable_blob                       passive_bytes;
};

typedef std::variant<
   reserved_submission,
   block_submission,
   transaction_submission,
   query_submission > submission_item;

struct reserved_submission_result {};

struct block_submission_result {};

struct transaction_submission_result {};

struct submission_error_result
{
   types::variable_blob error_text;
};

typedef std::variant<
   reserved_submission_result,
   block_submission_result,
   transaction_submission_result,
   query_submission_result,
   submission_error_result > submission_result;

} } } // koinos::types::rpc
