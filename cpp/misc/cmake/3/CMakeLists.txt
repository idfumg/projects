cmake_minimum_required(VERSION 3.18)
project(hello LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

file(GLOB_RECURSE FOO_SRC_FILES foo/src/*.cpp)
add_library(foo STATIC ${FOO_SRC_FILES})
target_include_directories(foo PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/foo/include)

add_executable(hello main.cpp)
target_link_libraries(hello PUBLIC foo)

# https://youtu.be/mKZ-i-UfGgQ (Simplified CMake Tutorial)