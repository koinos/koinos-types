find_package(Boost CONFIG REQUIRED)
find_package(nlohmann_json CONFIG REQUIRED)

# C++
set(KOINOS_CPP_OUTPUT_DIR "${CMAKE_BINARY_DIR}/generated/cpp")
set(KOINOS_CPP_OUTPUT_INCLUDE_DIR "${KOINOS_CPP_OUTPUT_DIR}/include")

set(KOINOS_CPP_OUTPUT_FILES
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/classes.hpp"

   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/thunk_ids.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/system_call_ids.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/thunk_ids.h"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/system_call_ids.h"

   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/basetypes.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/binary_fwd.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/binary.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/binary_serializer.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/exceptions.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/json_fwd.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/json.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/pack_fwd.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/reflect.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/typename.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/varint.hpp"

   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/util/base58.hpp"
   "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/koinos/pack/rt/util/variant_helpers.hpp"
)

add_custom_target(cpp ALL)
add_dependencies(cpp make_schema)
add_custom_command(TARGET cpp
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   ${PYTHON_BINARY} -m koinos_codegen.codegen
   --target-path "${KOINOS_REFLECT_TEMPLATE_DIR}"
   --target cpp
   -p koinos/pack
   -o "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}"
   ${KOINOS_SCHEMA_FILES}
)

set_source_files_properties(${KOINOS_CPP_OUTPUT_FILES} PROPERTIES GENERATED TRUE)

add_library(koinos_types INTERFACE)
add_dependencies(koinos_types cpp)
add_library(Koinos::types ALIAS koinos_types)

# We specifically do not want to link to boost, only include boost headers.
# Any boost library that is not header only must be build against wasm.
# It is easier to simply to utilize those libraries. Most boost libraries are header only.
target_link_libraries(koinos_types INTERFACE nlohmann_json::nlohmann_json)
target_include_directories(koinos_types INTERFACE ${KOINOS_CPP_OUTPUT_INCLUDE_DIR} ${Boost_INCLUDE_DIR} )

if( INSTALL_KOINOS_TYPES_HEADERS )
   install( DIRECTORY ${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/ DESTINATION include )
endif()

