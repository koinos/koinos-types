add_executable(test_multibase main.cpp)
target_link_libraries(test_multibase)
# TODO To avoid compiling, we hack it into the correct include path with a symlink -- is there a better way to do this
target_include_directories(test_multibase PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include)
install(TARGETS
   test_multibase
   RUNTIME DESTINATION bin
   LIBRARY DESTINATION lib
   ARCHIVE DESTINATION lib
)
