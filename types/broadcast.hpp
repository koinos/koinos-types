namespace koinos { namespace broadcast {

struct transaction_accepted
{
   protocol::transaction transaction;
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
