
namespace koinos { namespace chain {

// Do not edit this list directly, edit system_call_names.json and run codegen_thunk_ids.py to update
enum class system_call_id : uint32
{
   prints = 0x9b229941,
   apply_block = 0x94ab4ff5,
   apply_transaction = 0x9d8b55da,
   apply_reserved_operation = 0x9aa85924,
   apply_upload_contract_operation = 0x9e6ea937,
   apply_execute_contract_operation = 0x92184686,
   apply_set_system_call_operation = 0x9579a45c,
   db_put_object = 0x971ec7a2,
   db_get_object = 0x9766a8fb,
   db_get_next_object = 0x99a398e8,
   db_get_prev_object = 0x9bd3767c,
   execute_contract = 0x98c12cfe,
   get_entry_point = 0x9ff207dd,
   get_contract_args_size = 0x9b0d8fd9,
   get_contract_args = 0x9fbba198,
   set_contract_return = 0x9f49cdea,
   exit_contract = 0x98df75b0,
   get_head_info = 0x956fb22d,
   hash = 0x99770e04,
   recover_public_key = 0x95724a88,
   verify_block_signature = 0x9d1c3c89,
   verify_merkle_root = 0x996e24b9,
   get_transaction_payer = 0x9db35086,
   get_max_account_resources = 0x90f14f8d,
   get_transaction_resource_limit = 0x9940f685,
   get_last_irreversible_block = 0x953d2e37,
   get_caller = 0x94176c5f,
   require_authority = 0x9491e528,
   get_transaction_signature = 0x9dc3ec34,
   get_contract_id = 0x95e30f50,
   get_head_block_time = 0x941d5ab9
};

} } // koinos::chain
