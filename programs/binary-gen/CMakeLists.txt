add_executable( binary-gen main.cpp )
target_link_libraries( binary-gen koinos_types )
target_include_directories( binary-gen PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include )
