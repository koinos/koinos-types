namespace koinos { namespace types { namespace rpc {

typedef koinos::rpc::koinosd::reserved_rpc_params reserved_query_params;

typedef koinos::rpc::koinosd::get_head_info_params get_head_info_params;

typedef koinos::rpc::koinosd::get_chain_id_params get_chain_id_params;

typedef std::variant<
   reserved_query_params,
   get_head_info_params,
   get_chain_id_params > query_param_item;

typedef opaque< query_param_item > query_submission;

typedef koinos::rpc::koinosd::reserved_rpc_result reserved_query_result;

typedef koinos::rpc::koinosd::rpc_error query_error;

typedef koinos::rpc::koinosd::get_head_info_result get_head_info_result;

typedef koinos::rpc::koinosd::get_chain_id_result get_chain_id_result;

typedef std::variant<
   reserved_query_result,
   query_error,
   get_head_info_result,
   get_chain_id_result > query_item_result;

typedef opaque< query_item_result > query_submission_result;

typedef koinos::rpc::koinosd::reserved_rpc_params reserved_submission;

typedef koinos::rpc::koinosd::submit_block_params block_submission;

typedef koinos::rpc::koinosd::submit_transaction_params transaction_submission;

typedef std::variant<
   reserved_submission,
   block_submission,
   transaction_submission,
   query_submission > submission_item;

typedef koinos::rpc::koinosd::reserved_rpc_result reserved_submission_result;

typedef koinos::rpc::koinosd::submit_block_result block_submission_result;

typedef koinos::rpc::koinosd::submit_transaction_result transaction_submission_result;

typedef koinos::rpc::koinosd::rpc_error submission_error_result;

typedef std::variant<
   reserved_submission_result,
   block_submission_result,
   transaction_submission_result,
   query_submission_result,
   submission_error_result > submission_result;

} } } // koinos::types::rpc
