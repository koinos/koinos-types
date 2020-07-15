
namespace koinos { namespace types { namespace system {

// Use generate_ids.py to generate the system call id
enum class system_call_id : uint32
{
   prints = 0x9b229941,
   verify_block_header = 0x9625504e,
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
   get_contract_args_size = 0x9b0d8fd9,
   get_contract_args = 0x9fbba198,
   set_contract_return = 0x9f49cdea,
   exit_contract = 0x98df75b0,
   get_head_info = 0x956fb22d,
   hash = 0x99770e04,
   verify_block_sig = 0x95836910,
   verify_merkle_root = 0x996e24b9
};

} } } // koinos::types::system
