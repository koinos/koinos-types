namespace koinos { namespace types { namespace protocol {

struct reserved_operation
{
   unused_extensions_type         extensions;
};

struct nop_operation
{
   unused_extensions_type         extensions;
};

struct create_system_contract_operation
{
   contract_id_type               contract_id;
   variable_blob                  bytecode;
   unused_extensions_type         extensions;
};

struct set_system_call_operation
{
   uint32                         call_id;
   system::system_call_target     target;
   unused_extensions_type         extensions;
};

struct contract_call_operation
{
   contract_id_type               contract_id;
   uint32                         entry_point;
   variable_blob                  args;
   unused_extensions_type         extensions;
};

typedef std::variant<
   reserved_operation,
   nop_operation,
   create_system_contract_operation,
   contract_call_operation,
   set_system_call_operation
   > operation;

struct active_transaction_data {};

struct passive_transaction_data {};

struct transaction
{
   opaque< active_transaction_data >   active_data;
   opaque< passive_transaction_data >  passive_data;
   variable_blob                       signature_data;

   std::vector< operation >            operations;
};

enum class header_hash_index : uint32
{
   /**
    * Hash of the previous block.
    */
   previous_block_hash_index = 0,

   /**
    * Hash of Merkle root of transactions.
    */
   transaction_merkle_root_hash_index = 1,

   /**
    * Hash of Merkle root of passive data.
    * Includes transaction passives, transaction signatures and block passives.
    */
   passive_data_merkle_root_hash_index = 2,

   /**
    * Number of header hashes.
    */
   NUM_HEADER_HASHES = 3
};

struct active_block_data
{
   /**
    * Hashes included in the header.
    * All hashes must use the same algorithm.
    */
   multihash_vector               header_hashes;

   /**
    * Block height.  The genesis block has height=1.
    */
   block_height_type              height;

   /**
    * The timestamp.  Must be zero.
    */
   timestamp_type                 timestamp;
};

struct passive_block_data {};

typedef opaque< transaction > opaque_transaction;

struct block
{
   opaque< active_block_data >   active_data;
   opaque< passive_block_data >  passive_data;
   variable_blob                 signature_data;

   std::vector< opaque_transaction >    transactions;
};

struct block_receipt {};

} } } // koinos::types::protocol
