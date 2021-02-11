namespace koinos { namespace system {

struct head_info
{
   koinos::multihash              id;
   koinos::block_height_type      height;
};

struct block_part
{
   variable_blob                  active_data;
   variable_blob                  passive_data;
   variable_blob                  sig_data;
};

} } // koinos::system
