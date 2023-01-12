#include <iostream>
#include <type_traits>

using namespace std;

template<class T> T constexpr Min(const T& a, const T& b) {
    return a < b ? a : b;
}

template<class T1, class T2> constexpr bool Less(const T1& a, const T2& b) {
    static_assert(std::is_arithmetic_v<T1> && std::is_arithmetic_v<T2>, "");
    return a < b;
}

template<class T> void Swap(T& a, T& b) {
    T temp = a;
    a = b;
    b = temp;
}

int main() {
    static_assert(Min(101, 202) == 101, "");
    static_assert(Less(101, 202.0), "");
    int a = 1, b = 2;
    Swap(a, b);
    assert(a == 2);
    return 0;
}