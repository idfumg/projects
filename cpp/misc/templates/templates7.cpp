#include <iostream>

using namespace std;

template<class T>
inline constexpr T* copy_array(const T* arr, const size_t n) {
    const size_t amt = sizeof(T) * n;
    T* d = static_cast<T*>(::operator new(amt));
    if constexpr (is_trivially_copyable<T>::value) {
        d = static_cast<T*>(memcpy(d, arr, amt));
    }
    else {
        for (size_t i = 0; i != n; ++i) {
            new (&d[i]) T(arr[i]);
        }
    }
}

int main() {
    return 0;
}