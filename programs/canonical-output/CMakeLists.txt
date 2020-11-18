macro(SUBDIRLIST result curdir)
  file(GLOB children RELATIVE ${curdir}/lang ${curdir}/lang/*)
  set(dirlist "")
  foreach(child ${children})
    if(IS_DIRECTORY ${curdir}/lang/${child})
      list(APPEND dirlist ${child})
    endif()
  endforeach()
  set(${result} ${dirlist})
endmacro()

set(LANG_TARGETS "")

subdirlist(SUBDIRS ${CMAKE_CURRENT_SOURCE_DIR})

foreach(subdir ${SUBDIRS})
  add_subdirectory(lang/${subdir})
  list(APPEND LANG_TARGETS "canonical-output-${subdir}")
endforeach()

add_custom_target( canonical-output ALL
   DEPENDS generate ${LANG_TARGETS}
)
