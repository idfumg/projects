file(GLOB_RECURSE FOO_SRC_FILES src/*.cpp)
add_library(foo SHARED ${FOO_SRC_FILES})
target_include_directories(foo PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include)