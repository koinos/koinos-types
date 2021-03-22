namespace koinos { namespace broadcast {

struct transaction_accepted
{
   protocol::transaction transaction;
   chain::account_type   payer;
   uint128               max_payer_resources;
   uint128               trx_resource_limit;
};

struct block_accepted
{
   protocol::block block;
};

struct block_irreversible
{
   block_topology topology;
};

} } // koinos::broadcast
