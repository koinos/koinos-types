find_package( Boost 1.57 COMPONENTS program_options REQUIRED )

set(KOINOS_REFLECT_PYTHONPATH "${CMAKE_CURRENT_SOURCE_DIR}/../../../koinos-types")

set(KOINOS_REFLECT_TEMPLATE_DIR "${KOINOS_REFLECT_PYTHONPATH}/lang")

message(STATUS "${KOINOS_REFLECT_PYTHONPATH}")

set(KOINOS_JSON_DATA_FILES
   "${KOINOS_REFLECT_PYTHONPATH}/json/test_data.json"
   )

set(KOINOS_TEST_GEN_FILES
   "${CMAKE_CURRENT_BINARY_DIR}/test.go"
)

message(STATUS "${KOINOS_JSON_DATA_FILES}")

add_custom_command(
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   ${PYTHON_BINARY} -m koinos_codegen.testgen
   --target-path "${KOINOS_REFLECT_TEMPLATE_DIR}"
   --target golang
   -o "${CMAKE_CURRENT_BINARY_DIR}"
   ${KOINOS_JSON_DATA_FILES}
   OUTPUT ${KOINOS_TEST_GEN_FILES}
)

set_source_files_properties(${KOINOS_TEST_GEN_FILES} PROPERTIES GENERATED TRUE)

add_custom_target( canonical-output-golang
   DEPENDS ${KOINOS_TEST_GEN_FILES}
)

add_custom_command(
   TARGET canonical-output-golang POST_BUILD
   COMMAND ${CMAKE_COMMAND} -E copy
           ${CMAKE_CURRENT_SOURCE_DIR}/driver.py
           ${CMAKE_CURRENT_BINARY_DIR}/driver.py)
