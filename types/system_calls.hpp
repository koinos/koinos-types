
namespace koinos { namespace types { namespace system {

struct system_call_target_reserved {};

struct contract_call_bundle
{
   contract_id_type contract_id;
   uint32           entry_point;
};

typedef std::variant<
   system_call_target_reserved,
   thunks::thunk_id,
   contract_call_bundle
   > system_call_target;

} } } // koinos::types::system
