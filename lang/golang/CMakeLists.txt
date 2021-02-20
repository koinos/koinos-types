# Golang
set(KOINOS_GOLANG_OUTPUT_DIR "${CMAKE_BINARY_DIR}/generated/golang")

set(GOLANG_MODULE_NAME github.com/koinos/koinos-types-golang)
set(GOLANG_MODULE_VERSION 1.15)
set(GOLANG_MODULE_SOURCE_DIR ${KOINOS_GOLANG_OUTPUT_DIR}/src/${GOLANG_MODULE_NAME})
configure_file(${PROJECT_SOURCE_DIR}/cmake/go.mod.in ${GOLANG_MODULE_SOURCE_DIR}/go.mod)

set(KOINOS_REFLECT_PYTHONPATH "${PROJECT_SOURCE_DIR}/programs/koinos-types")
set(KOINOS_REFLECT_TEMPLATE_DIR "${PROJECT_SOURCE_DIR}/programs/koinos-types/lang")
set(KOINOS_REFLECT_SRC_DIR "${CMAKE_CURRENT_SOURCE_DIR}/src")

add_custom_target(golang ALL)
add_dependencies(golang make_schema)
add_custom_command(TARGET golang
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   ${PYTHON_BINARY} -m koinos_codegen.codegen
   --target-path "${KOINOS_REFLECT_TEMPLATE_DIR}"
   --target golang
   -p src/${GOLANG_MODULE_NAME}
   -o "${KOINOS_GOLANG_OUTPUT_DIR}"
   ${KOINOS_SCHEMA_FILES}
)
add_custom_command(TARGET golang POST_BUILD
   COMMAND go mod tidy
   WORKING_DIRECTORY ${GOLANG_MODULE_SOURCE_DIR}
   COMMENT "Generating go.sum"
)

