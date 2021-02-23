# Typescript
set(KOINOS_TYPESCRIPT_OUTPUT_DIR "${CMAKE_BINARY_DIR}/generated/typescript")

set(TYPESCRIPT_MODULE_NAME github.com/koinos/koinos-types-typescript)
set(TYPESCRIPT_MODULE_SOURCE_DIR ${KOINOS_TYPESCRIPT_OUTPUT_DIR}/src/${TYPESCRIPT_MODULE_NAME})

set(KOINOS_REFLECT_PYTHONPATH "${PROJECT_SOURCE_DIR}/programs/koinos-types")
set(KOINOS_REFLECT_TEMPLATE_DIR "${PROJECT_SOURCE_DIR}/programs/koinos-types/lang")
set(KOINOS_REFLECT_SRC_DIR "${CMAKE_CURRENT_SOURCE_DIR}/src")

add_custom_target(typescript ALL)
add_dependencies(typescript make_schema)
add_custom_command(TARGET typescript
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   ${PYTHON_BINARY} -m koinos_codegen.codegen
   --target-path "${KOINOS_REFLECT_TEMPLATE_DIR}"
   --target typescript
   -p src/${TYPESCRIPT_MODULE_NAME}
   -o "${KOINOS_TYPESCRIPT_OUTPUT_DIR}"
   ${KOINOS_SCHEMA_FILES}
)
