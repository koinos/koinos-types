namespace koinos { namespace transaction_store {

struct transaction_record
{
   protocol::transaction    transaction;
   std::vector< multihash > containing_blocks;
};

} } // koinos::transaction_store
