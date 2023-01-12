#include <iostream>

template<class T> struct remove_reference { using type = T; };
template<class T> struct remove_reference<T&> { using type = T; };
template<class T> struct remove_reference<T&&> { using type = T; };
template<class T> using remove_reference_t = typename remove_reference<T>::type;

template<class T> struct is_lvalue_reference { bool value = false; };
template<class T> struct is_lvalue_reference<T&> { bool value = true; };
template<class T> bool is_lvalue_reference_v = is_lvalue_reference<T>::value;

template<class T>
constexpr T&& forward(remove_reference_t<T>& arg) noexcept {
    return static_cast<T&&>(arg);
}

template<class T>
constexpr T&& forward(remove_reference_t<T>&& arg) noexcept {
    static_assert(!is_lvalue_reference_v<T>, "invalid rvalue to lvalue conversion!");
    return static_cast<T&&>(arg);
}

int main() {
    return 0;
}