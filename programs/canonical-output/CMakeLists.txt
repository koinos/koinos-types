MACRO(SUBDIRLIST result curdir)
  FILE(GLOB children RELATIVE ${curdir}/lang ${curdir}/lang/*)
  SET(dirlist "")
  FOREACH(child ${children})
    IF(IS_DIRECTORY ${curdir}/lang/${child})
      LIST(APPEND dirlist ${child})
    ENDIF()
  ENDFOREACH()
  SET(${result} ${dirlist})
ENDMACRO()

SET(LANG_TARGETS "")

SUBDIRLIST(SUBDIRS ${CMAKE_CURRENT_SOURCE_DIR})

FOREACH(subdir ${SUBDIRS})
  ADD_SUBDIRECTORY(lang/${subdir})
  LIST(APPEND LANG_TARGETS "canonical-output-${subdir}")
ENDFOREACH()

message(STATUS "${LANG_TARGETS}")

add_custom_target( canonical-output
   DEPENDS ${LANG_TARGETS}
)
