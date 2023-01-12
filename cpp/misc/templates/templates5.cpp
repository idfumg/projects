#include <iostream>

using namespace std;

template<class T>
struct IsDefaultConstructible {
private:
    template<class U, typename = decltype(U())>
    static constexpr std::true_type test(const void*);

    template<class U>
    static constexpr std::false_type test(...);
public:
    static constexpr bool value = decltype(test<T>(nullptr))::value;
};

struct NoDefaultConstructibleStruct {
    NoDefaultConstructibleStruct(int) {}
};

int main() {
    static_assert(IsDefaultConstructible<int>::value, "");
    static_assert(!IsDefaultConstructible<NoDefaultConstructibleStruct>::value, "");
    return 0;
}