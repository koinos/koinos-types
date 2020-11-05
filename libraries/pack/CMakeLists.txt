find_package(Boost 1.57 REQUIRED)

# C++
set(KOINOS_CPP_OUTPUT_DIR "${CMAKE_BINARY_DIR}/generated/cpp")
set(KOINOS_CPP_OUTPUT_INCLUDE_DIR "${KOINOS_CPP_OUTPUT_DIR}/include")

# Golang
set(KOINOS_GOLANG_OUTPUT_DIR "${CMAKE_BINARY_DIR}/generated/golang")


set(KOINOS_REFLECT_PYTHONPATH "${CMAKE_CURRENT_SOURCE_DIR}/../../programs/koinos-types")
set(KOINOS_REFLECT_TEMPLATE_DIR "${CMAKE_CURRENT_SOURCE_DIR}/../../programs/koinos-types/lang")
set(KOINOS_REFLECT_SRC_DIR "${CMAKE_CURRENT_SOURCE_DIR}/src")

set(KOINOS_REFLECT_SOURCES
   "${KOINOS_REFLECT_SRC_DIR}/block.bt"
   "${KOINOS_REFLECT_SRC_DIR}/types.hpp"
   "${KOINOS_REFLECT_SRC_DIR}/block.hpp"
   "${KOINOS_REFLECT_SRC_DIR}/submit.hpp"
   "${KOINOS_REFLECT_SRC_DIR}/system_calls.hpp"
   "${KOINOS_REFLECT_SRC_DIR}/thunks.hpp"
   "${KOINOS_REFLECT_SRC_DIR}/system_call_ids.hpp"
   "${KOINOS_REFLECT_SRC_DIR}/thunk_ids.hpp"
   "${KOINOS_REFLECT_SRC_DIR}/chain.hpp"
   )

set(KOINOS_SCHEMA_DIR "${CMAKE_BINARY_DIR}/schema")
set(KOINOS_SCHEMA_FILES "${KOINOS_SCHEMA_DIR}/block.schema")

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

set(KOINOS_PACK_GOLANG_GEN_FILES
   "${KOINOS_CPP_OUTPUT_DIR}/koinos.go"
   )

add_custom_target(make_schema ALL)

add_custom_command(TARGET make_schema
   COMMAND ${CMAKE_COMMAND} -E make_directory ${KOINOS_SCHEMA_DIR}
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   python3 -m koinos_reflect.analyze
   ${KOINOS_REFLECT_SOURCES}
   -s
   -o ${KOINOS_SCHEMA_DIR}/block.schema
#   OUTPUT ${KOINOS_SCHEMA_DIR}/block.schema
   DEPENDS mkdir_schema ${KOINOS_REFLECT_SOURCES}
)

add_custom_command(
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   python3 -m koinos_codegen.codegen
   --target-path "${KOINOS_REFLECT_TEMPLATE_DIR}"
   --target cpp
   -p koinos/pack
   -o "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}"
   ${KOINOS_SCHEMA_FILES}
   OUTPUT ${KOINOS_CPP_OUTPUT_FILES}
   DEPENDS ${KOINOS_SCHEMA_FILES}
)

add_custom_target(golang ALL)
add_dependencies(golang make_schema)

add_custom_command(TARGET golang
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   python3 -m koinos_codegen.codegen
   --target-path "${KOINOS_REFLECT_TEMPLATE_DIR}"
   --target golang
   -p src
   -o "${KOINOS_GOLANG_OUTPUT_DIR}"
   ${KOINOS_SCHEMA_FILES}
#   OUTPUT ${KOINOS_PACK_GOLANG_GEN_FILES}
   DEPENDS ${KOINOS_SCHEMA_FILES}
)

set_source_files_properties(${KOINOS_CPP_OUTPUT_FILES} PROPERTIES GENERATED TRUE)

# Unfortunately cmake doesn't seem to support interface libraries with generated header files,
# so we create a dummy .a file

add_library(koinos_types ${KOINOS_CPP_OUTPUT_FILES}
   ${CMAKE_CURRENT_SOURCE_DIR}/pack.cpp
   )
add_dependencies(koinos_types make_schema)
add_custom_target(cpp DEPENDS koinos_types)

# We specifically do not want to link to boost, only include boost headers.
# Any boost library that is not header only must be build against wasm.
# It is easier to simply to utilize those libraries. Most boost libraries are header only.
target_link_libraries(koinos_types INTERFACE nlohmann_json )
target_include_directories(koinos_types INTERFACE "${CMAKE_CURRENT_SOURCE_DIR}/gen" "${CMAKE_CURRENT_BINARY_DIR}/gen" "${KOINOS_CPP_OUTPUT_INCLUDE_DIR}" ${Boost_INCLUDE_DIR} )

set_target_properties(koinos_types
    PROPERTIES
    ARCHIVE_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}/generated/cpp/lib"
    LIBRARY_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}/generated/cpp/lib"
    RUNTIME_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}/generated/cpp/bin"
)

if( INSTALL_KOINOS_TYPES_HEADERS )
   install( DIRECTORY ${KOINOS_CPP_OUTPUT_INCLUDE_DIR}/ DESTINATION include )
endif()
