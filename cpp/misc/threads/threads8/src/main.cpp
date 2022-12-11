#include <algorithm>
#include <condition_variable>
#include <iostream>
#include <chrono>
#include <thread>
#include <mutex>
#include <future>

using namespace std;
using i32 = std::int32_t;
using i64 = std::int64_t;

std::mutex m;

void fun(i32 n) {
    std::lock_guard guard(m);
    cout << std::this_thread::get_id() << ": " << endl;
    for (i32 i = 0; i < n; ++i) {
        cout << i << ' ';
    }
    cout << endl;
}

std::condition_variable cv;
i64 balance = 0;

void waitForMoney(i64 amount) {
    std::unique_lock lock(m);
    cv.wait(lock, [](){return balance > 0;});
    if (balance >= amount) {
        balance -= amount;
    }
    cout << "balance was decreased and now it equals: " << balance << endl;
}

void increaseMoney(i64 amount) {
    std::lock_guard lock(m);
    balance += amount;
    cout << "balance was increased by: " << amount << " and now it equals: " << balance << endl;
    cv.notify_all();
}

std::mutex m1, m2, m3;

void getSeveralLocks() {
    std::lock(m1, m2, m3);
    cout << "Thread#" << this_thread::get_id() << " running" << endl;
    m1.unlock();
    m2.unlock();
    m3.unlock();
}

void workingWithPromise(std::promise<i32>& resultPromise) {
    i32 result = 0;
    for (int i = 0; i < 1'000'000; ++i) {
        result += 1;
    }
    std::this_thread::sleep_for(std::chrono::seconds(2));
    resultPromise.set_value(result);
}

int getSum() {
    int sum = 0;
    for (int i = 0; i < 10; ++i) {
        sum += i;
    }
    std::this_thread::sleep_for(std::chrono::seconds(1));
    return sum;
}

std::condition_variable cvx;
std::mutex mx;
const auto buffer_max_size = 5;
int buffer_size = 0;
atomic_bool is_running = true;

void producer(i32 n) {
    while (n--) {
        std::unique_lock lock(mx);
        cvx.wait(lock, [](){return buffer_size < buffer_max_size;});
        ++buffer_size;
        cout << "produce. now buffer size = " << buffer_size << endl;
        lock.unlock();
        cvx.notify_one();
    }
    is_running = false;
}

void consumer() {
    while (is_running or buffer_size > 0) {
        std::unique_lock lock(mx);
        cvx.wait(lock, [](){return buffer_size > 0;});
        --buffer_size;
        cout << "consume. now buffer size = " << buffer_size << endl;
        lock.unlock();
        cvx.notify_one();
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }
}

int main(int argc, char** argv) {
    thread t1(fun, 5);
    thread t2(fun, 5);
    t1.join();
    t2.join();

    thread t3(waitForMoney, 500);
    thread t4(increaseMoney,500);
    t3.join();
    t4.join();

    thread t5(getSeveralLocks);
    thread t6(getSeveralLocks);
    t5.join();
    t6.join();

    std::promise<i32> resultPromise;
    std::future f = resultPromise.get_future();
    thread t7(workingWithPromise, std::ref(resultPromise));
    cout << "waiting for the future" << endl;
    cout << "result: " << f.get() << endl;
    t7.join();

    std::future f2 = std::async(std::launch::async, getSum);
    cout << "do something important" << endl;
    cout << "sum = " << f2.get() << endl;

    thread t8(producer, 15);
    thread t9(consumer);
    t8.join();
    t9.join();

    return EXIT_SUCCESS;
}