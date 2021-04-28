namespace koinos { namespace broadcast {

struct transaction_accepted
{
   protocol::transaction transaction;
   chain::account_type   payer;
   uint128               max_payer_resources;
   uint128               trx_resource_limit;
   block_height_type     height;
};

struct block_accepted
{
   protocol::block block;
};

struct block_irreversible
{
   block_topology topology;
};

struct fork_heads
{
   std::vector< block_topology > fork_heads;
   block_topology                last_irreversible_block;
};

} } // koinos::broadcast
