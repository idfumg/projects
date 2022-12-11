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
}

int main(int argc, char** argv) {
    // function
    std::thread t1(fun, 10);
    if (t1.joinable()) t1.join();

    // lambda
    std::thread t2([](i64 x){
        fun(x);
    }, 10);
    if (t2.joinable()) t2.join();

    // functor
    struct Functor {
        void operator()(i64 x) const {
            fun(x);
        }
    };
    std::thread t3(Functor(), 10);
    if (t3.joinable()) t3.join();

    // non-static member function
    struct Foo {
        void run(i64 x) const {
            fun(10);
        }
    };
    std::thread t4(&Foo::run, Foo(), 10);
    if (t4.joinable()) t4.join();

    // static member function
    struct Bar {
        static void run(i64 x) {
            fun(10);
        }
    };
    std::thread t5(&Bar::run, 10);
    if (t5.joinable()) t5.join();

    return EXIT_SUCCESS;
}