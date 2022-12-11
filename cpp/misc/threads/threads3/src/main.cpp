#include <iostream>
#include <chrono>
#include <thread>

using namespace std;
using i32 = std::int32_t;
using i64 = std::int64_t;

void fun(i64 x) {
    while (x--) {
        cout << x << ' ';
    }
    cout << endl;
    std::this_thread::sleep_for(std::chrono::seconds(1));
}

int main(int argc, char** argv) {
    cout << "Thread#1" << endl;
    std::thread t1(fun, 10);
    if (t1.joinable()) t1.join();

    cout << "Thread#2" << endl;
    std::thread t2(fun, 10);
    if (t2.joinable()) t2.detach();

    cout << "Need to wait thread#2 because it was detached" << endl;
    std::this_thread::sleep_for((chrono::seconds(1)));

    return EXIT_SUCCESS;
}