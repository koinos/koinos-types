namespace koinos { namespace broadcast {

struct transaction_accepted
{
   multihash                    id;
   protocol::transaction        transaction;
};

struct block_accepted
{
   block_topology               topology;
   protocol::block              block;
};

} } // koinos::broadcast
