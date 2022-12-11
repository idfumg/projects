#include <iostream>
#include <chrono>
#include <thread>
#include <mutex>

using namespace std;
using i32 = std::int32_t;
using i64 = std::int64_t;

std::recursive_mutex m;
int cnt = 0;

void rec(int n) {
    if (not n) return;
    m.lock();
    cout << string(n - 1, ' ') << std::this_thread::get_id() << ' ' << cnt++ << endl;
    rec(n - 1);
    m.unlock();
}

int main(int argc, char** argv) {
    thread t1(rec, 5);
    thread t2(rec, 5);
    t1.join();
    t2.join();

    for (i32 i = 0; i < 5; ++i) {
        m.lock();
    }
    for (i32 i = 0; i < 5; ++i) {
        m.unlock();
    }
    return EXIT_SUCCESS;
}