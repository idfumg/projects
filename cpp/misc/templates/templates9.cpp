#include <iostream>
#include <type_traits>

using namespace std;

template<class T, class U>
using Similar = is_same<decay_t<T>, decay_t<U>>;

template<class T, class U>
using NotSimilar = enable_if_t<!Similar<T, U>::value>;

template<class T>
class Foo {
public:
    void bar(const T&); // lvalue
    void bar(T&&); // rvalue
    template<class U, typename = NotSimilar<T, U>> void uni(U&&); // uni value because of the type deduction
};

template<class T, template<typename> class ... Ts>
inline constexpr auto satisfies_all = (... && Ts<T>::value);

template<class T, template<typename> class ... Ts>
inline constexpr auto satisfies_some = (... || Ts<T>::value);

template<class T, template<typename> class ... Ts>
inline constexpr auto satisfies_none = (... && !Ts<T>::value);

template<class T>
inline constexpr auto satisfies_my_needs = satisfies_all<T, is_signed, is_pod> && satisfies_none<T, is_polymorphic, is_array>;

int main() {
    static_assert(satisfies_my_needs<int>, "My needs are unsatisfied!");
    return 0;
}