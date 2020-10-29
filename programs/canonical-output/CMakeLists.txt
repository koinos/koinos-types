find_package( Boost 1.57 COMPONENTS program_options REQUIRED )
add_executable( canonical-output main.cpp )
target_link_libraries( canonical-output koinos_types Boost::program_options )
target_include_directories( canonical-output PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include )
