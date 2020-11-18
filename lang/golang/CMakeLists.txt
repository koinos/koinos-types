# Golang
set(KOINOS_GOLANG_OUTPUT_DIR "${CMAKE_BINARY_DIR}/generated/golang")

set(GOLANG_MODULE_NAME koinos)
set(GOLANG_MODULE_VERSION 1.15)
configure_file(${CMAKE_MODULE_PATH}/go.mod.in ${KOINOS_GOLANG_OUTPUT_DIR}/src/${GOLANG_MODULE_NAME}/go.mod)

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

add_custom_command(
   TARGET golang POST_BUILD
   COMMAND ${CMAKE_COMMAND} -E copy
           ${KOINOS_GOLANG_OUTPUT_DIR}/src/koinos/koinos_test.go
           ${KOINOS_GOLANG_OUTPUT_DIR}/tests/koinos_test.go
   COMMAND ${CMAKE_COMMAND} -E remove
           ${KOINOS_GOLANG_OUTPUT_DIR}/src/koinos/koinos_test.go
)
