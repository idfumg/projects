cmake_minimum_required(VERSION 3.18)
project(hello LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

file(GLOB_RECURSE SRC_FILES src/*.cpp)
add_executable(hello main.cpp ${SRC_FILES})
target_include_directories(hello PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include)

# https://youtu.be/mKZ-i-UfGgQ (Simplified CMake Tutorial)