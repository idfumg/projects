#include <iostream>
#include <type_traits>

using namespace std;

struct Shape {};

template<class T>
using IsShape = enable_if_t<is_base_of_v<Shape, T>>;

template<class T, typename = IsShape<T>>
void munge_shape(const T& shape) {

}

int main() {
    return 0;
}