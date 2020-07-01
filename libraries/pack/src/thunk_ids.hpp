
namespace koinos { namespace types { namespace thunks {

enum class thunk_id
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
   get_contract_args_size = 0x83378e86,
   get_contract_args = 0x8e189d86,
   exit_contract = 0x81f61f9f
};

} } } // koinos::types::thunks