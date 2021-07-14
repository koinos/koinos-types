# Typescript
set(KOINOS_TYPESCRIPT_OUTPUT_DIR "${CMAKE_BINARY_DIR}/generated/typescript")

set(TYPESCRIPT_MODULE_NAME github.com/koinos/koinos-types-typescript)
set(TYPESCRIPT_MODULE_VERSION 1.15)
set(TYPESCRIPT_MODULE_SOURCE_DIR ${KOINOS_TYPESCRIPT_OUTPUT_DIR}/src/${TYPESCRIPT_MODULE_NAME})
configure_file(${PROJECT_SOURCE_DIR}/cmake/typescript/.editorconfig ${TYPESCRIPT_MODULE_SOURCE_DIR}/.editorconfig)
configure_file(${PROJECT_SOURCE_DIR}/cmake/typescript/.eslintrc.js ${TYPESCRIPT_MODULE_SOURCE_DIR}/.eslintrc.js)
configure_file(${PROJECT_SOURCE_DIR}/cmake/typescript/jest.config.js ${TYPESCRIPT_MODULE_SOURCE_DIR}/jest.config.js)
configure_file(${PROJECT_SOURCE_DIR}/cmake/typescript/package-lock.json ${TYPESCRIPT_MODULE_SOURCE_DIR}/package-lock.json)
configure_file(${PROJECT_SOURCE_DIR}/cmake/typescript/package.json ${TYPESCRIPT_MODULE_SOURCE_DIR}/package.json)
configure_file(${PROJECT_SOURCE_DIR}/cmake/typescript/tsconfig-build.json ${TYPESCRIPT_MODULE_SOURCE_DIR}/tsconfig-build.json)
configure_file(${PROJECT_SOURCE_DIR}/cmake/typescript/tsconfig.eslint.json ${TYPESCRIPT_MODULE_SOURCE_DIR}/tsconfig.eslint.json)
configure_file(${PROJECT_SOURCE_DIR}/cmake/typescript/tsconfig.json ${TYPESCRIPT_MODULE_SOURCE_DIR}/tsconfig.json)

set(KOINOS_REFLECT_PYTHONPATH "${PROJECT_SOURCE_DIR}/programs/koinos-types")
set(KOINOS_REFLECT_TEMPLATE_DIR "${PROJECT_SOURCE_DIR}/programs/koinos-types/lang")
set(KOINOS_REFLECT_SRC_DIR "${CMAKE_CURRENT_SOURCE_DIR}/src")

add_custom_target(typescript ALL)
add_dependencies(typescript make_schema)
add_custom_command(TARGET typescript
   COMMAND ${CMAKE_COMMAND} -E env PYTHONPATH=${KOINOS_REFLECT_PYTHONPATH}
   ${PYTHON_BINARY} -m koinos_codegen.codegen
   --target-path "${KOINOS_REFLECT_TEMPLATE_DIR}"
   --target typescript
   -p src/${TYPESCRIPT_MODULE_NAME}
   -o "${KOINOS_TYPESCRIPT_OUTPUT_DIR}"
   ${KOINOS_SCHEMA_FILES}
)
add_custom_command(TARGET typescript POST_BUILD
   WORKING_DIRECTORY ${TYPESCRIPT_MODULE_SOURCE_DIR}
   COMMAND npm install
   COMMAND npx prettier ${KOINOS_TYPESCRIPT_OUTPUT_DIR}
   COMMAND npm run build
)