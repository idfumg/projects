#include <iostream>

using namespace std;

template<size_t> struct fib;
template<> struct fib<0> { static const auto value = 1; };
template<> struct fib<1> { static const auto value = 1; };
template<size_t n> struct fib { static const auto value = fib<n - 1>::value + fib<n - 2>::value; };

constexpr size_t fib14(size_t n) {
    if (n <= 1) return 1;
    return fib14(n - 1) + fib14(n - 2);
}

int main() {
    static_assert(fib<5>::value == 8, "");
    static_assert(fib14(5) == 8, "");
    return 0;
}