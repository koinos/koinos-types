namespace koinos { namespace rpc { namespace block_store {

struct block_store_reserved_request {};

struct block_store_reserved_response {};

struct get_blocks_by_id_request
{
   /**
    * The ID's of the blocks to get.
    */
   std::vector< multihash >              block_id;

   /**
    * If true, returns the blocks' contents.
    */
   boolean                               return_block_blob;

   /**
    * If true, returns the blocks' receipts.
    */
   boolean                               return_receipt_blob;
};

struct get_blocks_by_id_response
{
   std::vector< block_store::block_item >             block_items;
};

struct get_blocks_by_height_request
{
   multihash                             head_block_id;
   block_height_type                     ancestor_start_height;
   uint32                                num_blocks;

   boolean                               return_block;
   boolean                               return_receipt;
};

struct get_blocks_by_height_response
{
   std::vector< block_store::block_item >             block_items;
};

struct add_block_request
{
   block_store::block_item block_to_add;
};

struct add_block_response
{
};

struct get_highest_block_request {};

struct get_highest_block_response
{
   block_topology topology;
};

struct block_store_error_response
{
   std::string error_text;
   std::string error_data;
};

typedef std::variant<
   block_store_reserved_request,
   get_blocks_by_id_request,
   get_blocks_by_height_request,
   add_block_request,
   get_highest_block_request
   > block_store_request;

typedef std::variant<
   block_store_reserved_response,
   block_store_error_response,
   get_blocks_by_id_response,
   get_blocks_by_height_response,
   add_block_response,
   get_highest_block_response
   > block_store_response;

} } }
