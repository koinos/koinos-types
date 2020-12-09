macro(SUBDIRLIST result curdir)
  file(GLOB children RELATIVE ${curdir} ${curdir}/*)
  set(dirlist "")
  foreach(child ${children})
    if(IS_DIRECTORY ${curdir}/${child})
      list(APPEND dirlist ${child})
    endif()
  endforeach()
  set(${result} ${dirlist})
endmacro()

macro(LISTDIR result curdir)
  file(GLOB children RELATIVE ${curdir} ${curdir}/*)
  set(filelist "")
  foreach(child ${children})
    if(NOT IS_DIRECTORY ${curdir}/${child})
      list(APPEND filelist "${curdir}/${child}")
    endif()
  endforeach()
  set(${result} ${filelist})
endmacro()

set(KOINOS_REFLECT_PYTHONPATH "${PROJECT_SOURCE_DIR}/programs/koinos-types")
set(KOINOS_REFLECT_TEMPLATE_DIR "${PROJECT_SOURCE_DIR}/programs/koinos-types/lang")
set(KOINOS_REFLECT_SRC_DIR "${PROJECT_SOURCE_DIR}/types")

listdir(KOINOS_REFLECT_SOURCES ${KOINOS_REFLECT_SRC_DIR})

set(KOINOS_SCHEMA_DIR "${CMAKE_BINARY_DIR}/schema")
set(KOINOS_SCHEMA_FILES "${KOINOS_SCHEMA_DIR}/block.schema")

add_custom_target(make_schema ALL)
add_custom_command(TARGET make_schema
   COMMAND ${CMAKE_COMMAND} -E make_directory ${KOINOS_SCHEMA_DIR}
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   ${PYTHON_BINARY} -m koinos_reflect.analyze
   ${KOINOS_REFLECT_SOURCES}
   -s
   -o ${KOINOS_SCHEMA_DIR}/block.schema
)

set(LANG_TARGETS "")

subdirlist(SUBDIRS ${CMAKE_CURRENT_SOURCE_DIR})

foreach(subdir ${SUBDIRS})
  add_subdirectory(${subdir})
  list(APPEND LANG_TARGETS "${subdir}")
endforeach()

message(STATUS "${LANG_TARGETS}")

add_custom_target( generate
   DEPENDS ${LANG_TARGETS}
)
