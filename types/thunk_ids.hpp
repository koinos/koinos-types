
namespace koinos { namespace chain {

// Use generate_ids.py to generate the thunk id
enum class thunk_id : uint32
{
   prints = 0x8f6df54d,
   verify_block_header = 0x8d425aac,
   apply_block = 0x8d6d31a8,
   apply_transaction = 0x8981b0df,
   apply_reserved_operation = 0x8b3c14f6,
   apply_upload_contract_operation = 0x8882a55e,
   apply_execute_contract_operation = 0x85e882eb,
   apply_set_system_call_operation = 0x86f92c8c,
   db_put_object = 0x82038de5,
   db_get_object = 0x8862a0d8,
   db_get_next_object = 0x86e45047,
   db_get_prev_object = 0x8d57e8fd,
   execute_contract = 0x8a43fe83,
   get_entry_point = 0x807eb765,
   get_contract_args_size = 0x83378e86,
   get_contract_args = 0x8e189d86,
   set_contract_return = 0x86b86275,
   exit_contract = 0x81f61f9f,
   get_head_info = 0x89df34c4,
   hash = 0x8aaaf547,
   recover_public_key = 0x804f5450,
   verify_block_signature = 0x8fb9a59b,
   verify_merkle_root = 0x8ed9ddcb,
   get_transaction_payer = 0x86a87bf5,
   get_max_account_resources = 0x842c6c81,
   get_transaction_resource_limit = 0x8bdf81a1,
   get_last_irreversible_block = 0x80c3b893,
   get_caller = 0x82312501,
   require_authority = 0x8a06717d,
   get_transaction_signature = 0x83441b23,
   get_contract_id = 0x8df4ad21,
   get_head_block_time = 0x8e76d5b4
};

} } // koinos::thunks
