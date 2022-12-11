cmake_minimum_required(VERSION 3.18)
project(project LANGUAGES CXX VERSION 1.0)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(Boost 1.74 COMPONENTS REQUIRED)
find_package(Threads REQUIRED)
find_package(OpenSSL REQUIRED)

# Download all the submodules
find_package(Git QUIET)
if (GIT_FOUND AND EXISTS "${PROJECT_SOURCE_DIR}/.git")
    option(GIT_SUBMODULE "Check submodules during build" ON)
    if (GIT_SUBMODULE)
        message(STATUS "Submodule update")
        execute_process(COMMAND ${GIT_EXECUTABLE} submodule update --init --recursive
            WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
            RESULT_VARIABLE GIT_SUBMOD_RESULT)
        if(NOT GIT_SUBMOD_RESULT EQUAL "0")
            message(FATAL_ERROR "git submodule update --init failed with ${GIT_SUBMOD_RESULT}, please check")
        endif()
    endif()
endif()

# Check if all the submodules exist
if(NOT EXISTS "${PROJECT_SOURCE_DIR}/external/json/CMakeLists.txt")
    message(FATAL_ERROR "The json submodules was not downloaded! GIT_SUBMODULE was turned off or failed.")
endif()

add_subdirectory(external/json)

add_executable(${PROJECT_NAME} src/main.cpp src/Client.cpp)

target_link_libraries(${PROJECT_NAME} PRIVATE ${Boost_LIBRARIES})
target_link_libraries(${PROJECT_NAME} PRIVATE Threads::Threads)
target_link_libraries(${PROJECT_NAME} PRIVATE OpenSSL::SSL)
target_include_directories(${PROJECT_NAME} PRIVATE ${Boost_INCLUDE_DIR})
target_include_directories(${PROJECT_NAME} PRIVATE external/json/include)
