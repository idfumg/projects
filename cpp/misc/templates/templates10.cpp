#include "../../template.hpp"
#include <type_traits>

template<class T, class U>
using is_enum_of_t = typename enable_if<is_enum_v<T> && is_same_v<underlying_type_t<T>, U>>::type;

template<char c>
struct m3_sig {
    static const char value = c;
};

template<class T, class = void> 
struct m3_type_to_sig;

template<class T>
struct m3_type_to_sig<T, is_enum_of_t<T, int32_t>> : m3_sig<'i'> {};

template<class T>
struct m3_type_to_sig<T, is_enum_of_t<T, int64_t>> : m3_sig<'I'> {};

template<> struct m3_type_to_sig<int32_t> : m3_sig<'i'> {};
template<> struct m3_type_to_sig<int64_t> : m3_sig<'I'> {};
template<> struct m3_type_to_sig<float>   : m3_sig<'f'> {};
template<> struct m3_type_to_sig<double>  : m3_sig<'F'> {};
template<> struct m3_type_to_sig<void>    : m3_sig<'v'> {};
template<> struct m3_type_to_sig<void *>  : m3_sig<'*'> {};
template<> struct m3_type_to_sig<const void *> : m3_sig<'*'> {};

enum Foo32 : int32_t {};
static_assert(m3_type_to_sig<Foo32>::value == 'i', "");

enum Foo64 : int64_t {};
static_assert(m3_type_to_sig<Foo64>::value == 'I', "");

static_assert(m3_type_to_sig<int32_t>::value == 'i', "");
static_assert(m3_type_to_sig<int64_t>::value == 'I', "");
static_assert(m3_type_to_sig<float>::value == 'f', "");
static_assert(m3_type_to_sig<double>::value == 'F', "");
static_assert(m3_type_to_sig<void>::value == 'v', "");
static_assert(m3_type_to_sig<void*>::value == '*', "");
static_assert(m3_type_to_sig<const void*>::value == '*', "");

int main() {

}
