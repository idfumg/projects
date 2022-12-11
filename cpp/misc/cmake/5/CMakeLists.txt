cmake_minimum_required(VERSION 3.18)
project(hello LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

FIND_PACKAGE(Boost 1.74 COMPONENTS REQUIRED)

add_subdirectory(foo)
add_subdirectory(websocketpp)
add_executable(${PROJECT_NAME} main.cpp)

target_link_libraries(${PROJECT_NAME} PRIVATE foo ${Boost_LIBRARIES})
target_include_directories(${PROJECT_NAME} PRIVATE ${Boost_INCLUDE_DIR})
target_include_directories(${PROJECT_NAME} PRIVATE websocketpp)
