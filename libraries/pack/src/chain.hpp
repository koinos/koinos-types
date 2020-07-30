namespace koinos { namespace types { namespace system {

struct head_info
{
   koinos::types::multihash              id;
   koinos::types::block_height_type      height;
};

struct block_part
{
   types::variable_blob                  active_data;
   types::variable_blob                  passive_data;
   types::variable_blob                  sig_data;
};

} } } // koinos::types::system
