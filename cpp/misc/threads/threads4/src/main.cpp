#include <iostream>
#include <chrono>
#include <thread>
#include <mutex>

using namespace std;
using i32 = std::int32_t;
using i64 = std::int64_t;

i32 global = 0;
std::mutex m;

void fun() {
    m.lock();
    ++global;
    m.unlock();
}

void fun2() {
    for (i32 i = 0; i < 1'000'000; ++i) {
        if (m.try_lock()) { // don't sleep on lock
            ++global;
            m.unlock();
        }
    }
}

int main(int argc, char** argv) {
    thread t1(fun);
    thread t2(fun);

    t1.join();
    t2.join();

    cout << global << endl;

    thread t3(fun2);
    thread t4(fun2);

    t3.join();
    t4.join();

    cout << global << endl;

    return EXIT_SUCCESS;
}