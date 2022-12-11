cmake_minimum_required(VERSION 3.18)
project(hello LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_subdirectory(foo)
add_executable(hello main.cpp)
target_link_libraries(hello PRIVATE foo)

# https://youtu.be/mKZ-i-UfGgQ (Simplified CMake Tutorial)