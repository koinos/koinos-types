find_package( Boost 1.57 COMPONENTS program_options REQUIRED )
add_executable( binary-gen main.cpp )
target_link_libraries( binary-gen koinos_types Boost::program_options )
target_include_directories( binary-gen PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include )
