find_package( Boost 1.57 COMPONENTS program_options REQUIRED )

set(KOINOS_REFLECT_PYTHONPATH "${CMAKE_CURRENT_SOURCE_DIR}/../../../koinos-types")

set(KOINOS_REFLECT_TEMPLATE_DIR "${KOINOS_REFLECT_PYTHONPATH}/lang")

message(STATUS "${KOINOS_REFLECT_PYTHONPATH}")

set(KOINOS_JSON_DATA_FILES
   "${KOINOS_REFLECT_PYTHONPATH}/json/test_data.json"
   )

set(KOINOS_TEST_GEN_FILES
   "${CMAKE_CURRENT_BINARY_DIR}/main.cpp"
)

message(STATUS "${KOINOS_JSON_DATA_FILES}")

add_custom_command(
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   python3 -m koinos_codegen.testgen
   --target-path "${KOINOS_REFLECT_TEMPLATE_DIR}"
   --target cpp
   -o "${CMAKE_CURRENT_BINARY_DIR}"
   ${KOINOS_JSON_DATA_FILES}
   OUTPUT ${KOINOS_TEST_GEN_FILES}
)

set_source_files_properties(${KOINOS_TEST_GEN_FILES} PROPERTIES GENERATED TRUE)

add_executable( canonical-output ${KOINOS_TEST_GEN_FILES} )
target_link_libraries( canonical-output koinos_types Boost::program_options )
target_include_directories( canonical-output PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include )
