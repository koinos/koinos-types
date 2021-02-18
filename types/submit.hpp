namespace koinos { namespace types { namespace rpc {

typedef koinos::rpc::chain::chain_reserved_request reserved_query_params;

typedef koinos::rpc::chain::get_head_info_request get_head_info_params;

typedef koinos::rpc::chain::get_chain_id_request get_chain_id_params;

typedef std::variant<
   reserved_query_params,
   get_head_info_params,
   get_chain_id_params > query_param_item;

typedef opaque< query_param_item > query_submission;

typedef koinos::rpc::chain::chain_reserved_response reserved_query_result;

typedef koinos::rpc::chain::chain_error_response query_error;

typedef koinos::rpc::chain::get_head_info_response get_head_info_result;

typedef koinos::rpc::chain::get_chain_id_response get_chain_id_result;

typedef std::variant<
   reserved_query_result,
   query_error,
   get_head_info_result,
   get_chain_id_result > query_item_result;

typedef opaque< query_item_result > query_submission_result;

typedef koinos::rpc::chain::chain_reserved_request reserved_submission;

typedef koinos::rpc::chain::submit_block_request block_submission;

typedef koinos::rpc::chain::submit_transaction_request transaction_submission;

typedef std::variant<
   reserved_submission,
   block_submission,
   transaction_submission,
   query_submission > submission_item;

typedef koinos::rpc::chain::chain_reserved_response reserved_submission_result;

typedef koinos::rpc::chain::submit_block_response block_submission_result;

typedef koinos::rpc::chain::submit_transaction_response transaction_submission_result;

typedef koinos::rpc::chain::chain_error_response submission_error_result;

typedef std::variant<
   reserved_submission_result,
   block_submission_result,
   transaction_submission_result,
   query_submission_result,
   submission_error_result > submission_result;

} } } // koinos::types::rpc
