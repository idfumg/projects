#include <iostream>
#include <chrono>
#include <thread>
#include <mutex>

using namespace std;
using i32 = std::int32_t;
using i64 = std::int64_t;

i32 x = 0, y = 0, cnt = 5;
std::mutex m1, m2;

void fun1(int& value, std::mutex& m) {
    while (cnt) {
        m.lock();
        cout << "incrementing the value int a thread #" << std::this_thread::get_id() << endl;
        ++value;
        m.unlock();
        std::this_thread::sleep_for(std::chrono::microseconds(1'000'000));
    }
}

void fun2() {
    while (cnt) {
        if (int res = std::try_lock(m1, m2); res == -1) {
            if (x > 0 and y > 0) {
                x = y = 0;
                --cnt;
                cout << "consume x and y" << endl;
            }
            m1.unlock();
            m2.unlock();
        }
    }
}

int main(int argc, char** argv) {
    thread t1(fun1, std::ref(x), std::ref(m1));
    thread t2(fun1, std::ref(y), std::ref(m2));
    thread t3(fun2);

    t1.join();
    t2.join();
    t3.join();

    return EXIT_SUCCESS;
}