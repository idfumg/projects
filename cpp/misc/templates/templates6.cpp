#include <iostream>

using namespace std;

class Foo {
    string s;
    string t;
    int n;
public:
    template<class T, class U = string, typename = enable_if_t<is_convertible_v<T, string>>, typename = enable_if_t<true>>
    Foo(T&& s, U&& t = "", int n = 0) : s(forward<T>(s)), t(forward<U>(t)), n(n) {}
};

int main() {
    Foo foo("hello", "world", 42);
    return 0;
}
