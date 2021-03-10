namespace koinos { namespace broadcast {

struct transaction_accepted
{
   transaction_topology  topology;
   protocol::transaction transaction;
};

struct block_accepted
{
   block_topology  topology;
   protocol::block block;
};

struct block_irreversible
{
   block_topology topology;
};

} } // koinos::broadcast
