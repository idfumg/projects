#include <iostream>
using namespace std;

template<size_t n> struct fact;

template<> struct fact<0> { static const auto value = 1; };
template<size_t n> struct fact { static const auto value = n * fact<n - 1>::value; };

int main() {
    static_assert(fact<5>::value == 120, "");
    return 0;
}