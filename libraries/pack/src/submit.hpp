namespace koinos { namespace types { namespace rpc {

// Every block has a cryptographic ID.
// Check the claimed ID against the block content.

struct block_topology
{
   types::multihash_type                 id;
   types::block_height_type              height;
   types::multihash_type                 previous;
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
   std::vector< system::block_part >          block_parts;

   boolean                                    verify_passive_data;
   boolean                                    verify_block_signature;
   boolean                                    verify_transaction_signatures;
};

struct transaction_submission
{
   types::variable_blob                       active_bytes;
   types::variable_blob                       passive_bytes;
};

struct query_submission
{
   types::variable_blob                       query;
};

typedef std::variant<
   reserved_submission,
   block_submission,
   transaction_submission,
   query_submission > submission_item;

struct get_head_info_params {};

typedef std::variant<
   get_head_info_params > query_param_item;

struct reserved_submission_result {};

struct block_submission_result {};

struct transaction_submission_result {};

struct query_submission_result
{
   types::variable_blob result;
};

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

typedef query_submission_result query_error;

typedef types::system::head_info get_head_info_result;

typedef std::variant<
   query_error,
   get_head_info_result > query_item_result;

} } } // koinos::types::rpc
