
cmake_minimum_required (VERSION 3.10.2)

find_program(CCACHE_PROGRAM ccache)
if(CCACHE_PROGRAM)
   set_property(GLOBAL PROPERTY RULE_LAUNCH_COMPILE "${CCACHE_PROGRAM}")
   set(CMAKE_XCODE_ATTRIBUTE_CC           "${CMAKE_SOURCE_DIR}/ci/ccache_clang")
   set(CMAKE_XCODE_ATTRIBUTE_CXX          "${CMAKE_SOURCE_DIR}/ci/ccache_clang++")
   set(CMAKE_XCODE_ATTRIBUTE_LD           "${CMAKE_SOURCE_DIR}/ci/ccache_clang")
   set(CMAKE_XCODE_ATTRIBUTE_LDPLUSPLUS   "${CMAKE_SOURCE_DIR}/ci/ccache_clang++")
endif()

option(HUNTER_ENABLED "Enable Hunter package manager" ON)
option(HUNTER_RUN_UPLOAD "Upload Hunter packages to binary cache server" OFF)

set(
   HUNTER_CACHE_SERVERS
   "https://github.com/koinos/hunter-cache"
   CACHE
   STRING
   "Default cache server"
)

set(
   HUNTER_PASSWORDS_PATH
   "${CMAKE_CURRENT_LIST_DIR}/cmake/Hunter/passwords.cmake"
   CACHE
   FILEPATH
   "Hunter passwords"
)

include("cmake/HunterGate.cmake")

HunterGate(
   URL "https://github.com/cpp-pm/hunter/archive/v0.23.287.tar.gz"
   SHA1 "69371a638d697bbf09f4f08dd0aad996efcc79d9"
   LOCAL
)

project(koinos_types VERSION 0.1.0 LANGUAGES CXX)

#
# CONFIGURATION
#
include(GNUInstallDirs)

set(KOINOS_LIB_TARGET_NAME                ${PROJECT_NAME})
set(KOINOS_LIB_TARGET_SHORT_NAME          "types")
set(KOINOS_LIB_CONFIG_INSTALL_DIR         "${CMAKE_INSTALL_LIBDIR}/cmake/${PROJECT_NAME}" CACHE INTERNAL "")
set(KOINOS_LIB_INCLUDE_INSTALL_DIR        "${CMAKE_INSTALL_INCLUDEDIR}")
set(KOINOS_LIB_TARGETS_EXPORT_NAME        "${PROJECT_NAME}Targets")
set(KOINOS_LIB_CMAKE_CONFIG_TEMPLATE      "cmake/config.cmake.in")
set(KOINOS_LIB_CMAKE_CONFIG_DIR           "${CMAKE_CURRENT_BINARY_DIR}")
set(KOINOS_LIB_CMAKE_VERSION_CONFIG_FILE  "${KOINOS_LIB_CMAKE_CONFIG_DIR}/${PROJECT_NAME}ConfigVersion.cmake")
set(KOINOS_LIB_CMAKE_PROJECT_CONFIG_FILE  "${KOINOS_LIB_CMAKE_CONFIG_DIR}/${PROJECT_NAME}Config.cmake")
set(KOINOS_LIB_CMAKE_PROJECT_TARGETS_FILE "${KOINOS_LIB_CMAKE_CONFIG_DIR}/${PROJECT_NAME}Targets.cmake")
set(KOINOS_LIB_PKGCONFIG_INSTALL_DIR      "${CMAKE_INSTALL_LIBDIR}/pkgconfig")

option(BUILD_TESTS "Build Tests" ON)
option(ENABLE_JSON "Build C++ with JSON support" ON)
option(FORCE_COLORED_OUTPUT "Always produce ANSI-colored output (GNU/Clang only)." OFF)

# This is to force color output when using ccache with Unix Makefiles
if( ${FORCE_COLORED_OUTPUT} )
    if( "${CMAKE_CXX_COMPILER_ID}" STREQUAL "GNU" )
       add_compile_options (-fdiagnostics-color=always)
    elseif( "${CMAKE_CXX_COMPILER_ID}" STREQUAL "Clang" OR "${CMAKE_CXX_COMPILER_ID}" STREQUAL "AppleClang" )
       add_compile_options (-fcolor-diagnostics)
    endif ()
endif ()

list(APPEND CMAKE_MODULE_PATH "${PROJECT_SOURCE_DIR}/cmake")

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
set(CMAKE_CXX_VISIBILITY_PRESET hidden)
set(Boost_NO_BOOST_CMAKE ON)

set(PYTHON_BINARY "python3" CACHE STRING "The default python to invoke")

if (NOT TARGET coverage)
  if (COVERAGE)
    include(CodeCoverage)
    append_coverage_compiler_flags()
    setup_target_for_coverage_lcov(
       NAME coverage
       LCOV_ARGS "--quiet" "--no-external"
       EXECUTABLE koinos_types_tests
       EXCLUDE "lang/cpp/vendor/*")
  endif()
endif()

hunter_add_package(Boost COMPONENTS system test program_options)
hunter_add_package(nlohmann_json)

add_subdirectory(lang)
add_subdirectory(programs)

if(BUILD_TESTS)
   add_subdirectory(tests)
endif()

# Install a pkg-config file, so other tools can find this.
configure_file(
   "${CMAKE_CURRENT_SOURCE_DIR}/cmake/pkg-config.pc.in"
   "${CMAKE_CURRENT_BINARY_DIR}/${PROJECT_NAME}.pc"
)

#
# INSTALL
# install header files, generate and install cmake config files for find_package()
#
include(CMakePackageConfigHelpers)

configure_file(
   "cmake/${PROJECT_NAME}ConfigVersion.cmake.in"
   ${KOINOS_LIB_CMAKE_VERSION_CONFIG_FILE}
   @ONLY
)
configure_file(
   ${KOINOS_LIB_CMAKE_CONFIG_TEMPLATE}
   ${KOINOS_LIB_CMAKE_PROJECT_CONFIG_FILE}
   @ONLY
)

install(
   FILES ${KOINOS_LIB_CMAKE_PROJECT_CONFIG_FILE} ${KOINOS_LIB_CMAKE_VERSION_CONFIG_FILE}
   DESTINATION ${KOINOS_LIB_CONFIG_INSTALL_DIR}
)

install(
   FILES "${CMAKE_CURRENT_BINARY_DIR}/${PROJECT_NAME}.pc"
   DESTINATION ${KOINOS_LIB_PKGCONFIG_INSTALL_DIR}
)
