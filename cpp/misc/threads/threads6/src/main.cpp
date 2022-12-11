#include <iostream>
#include <chrono>
#include <thread>
#include <mutex>

using namespace std;
using i32 = std::int32_t;
using i64 = std::int64_t;

i32 x = 0, y = 0, cnt = 5;
std::timed_mutex m;

void fun1() {
    if (m.try_lock_for(std::chrono::seconds(1))) {
        cout << "thread #" << std::this_thread::get_id() << " entered" << endl;
        std::this_thread::sleep_for(std::chrono::seconds(2));
    }
    else {
        cout << "thread #" << std::this_thread::get_id() << " could have not entered" << endl;
    }
}

int main(int argc, char** argv) {
    thread t1(fun1);
    thread t2(fun1);

    t1.join();
    t2.join();

    return EXIT_SUCCESS;
}