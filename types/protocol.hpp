namespace koinos { namespace protocol {

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
   chain::system_call_target      target;
   unused_extensions_type         extensions;
};

struct call_contract_operation
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
   call_contract_operation,
   set_system_call_operation
   > operation;

struct active_transaction_data
{
   uint128                  resource_limit;
   uint64                   nonce;
   std::vector< operation > operations;
};

struct passive_transaction_data {};

struct transaction
{
   multihash                           id;
   opaque< active_transaction_data >   active_data;
   opaque< passive_transaction_data >  passive_data;
   variable_blob                       signature_data;
};

struct active_block_data
{
   multihash                      transaction_merkle_root;
   multihash                      passive_data_merkle_root;
   multihash                      signer_address;
};

struct passive_block_data {};

struct block_header
{
   multihash                     previous;
   block_height_type             height;
   timestamp_type                timestamp;
};

/**
 * Topological constraints:  Constraints for a new block b that builds on a given block a
 *
 * b.header.previous_block == a.block_id
 * b.header.height         == a.header.height+1
 * b.header.timestamp      >  a.header.timestamp
 *
 * Cryptographic constraints:  Internal constraints for block b.
 *
 * b.block_id                             == H(b.header, b.active_data)
 * b.active_data.transaction_merkle_root  == Hm(b.transactions)
 * b.active_data.passive_data_merkle_root == Hm(b.passive_data)
 * b.active_data.signer_address           == H(recover(b.signature_data, b.block_id))
 */

struct block
{
   multihash                     id;
   block_header                  header;
   opaque< active_block_data >   active_data;
   opaque< passive_block_data >  passive_data;
   variable_blob                 signature_data;

   std::vector< transaction >    transactions;
};

struct block_receipt {};

} } // koinos::protocol
