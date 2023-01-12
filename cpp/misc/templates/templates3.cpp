#include <iostream>

using namespace std;

template<class ... Ts> struct Tuple;
template<> struct Tuple<> {};
template<class T, class ... Ts> struct Tuple<T, Ts...> : public Tuple<Ts...> { T value; };

template<size_t idx, class ... Ts> 
struct TupleElement;

template<class T, class ... Ts> 
struct TupleElement<0, Tuple<T, Ts...>> {
    using type = T;
    using TupleT = Tuple<T, Ts...>;
};

template<size_t idx, class T, class ... Ts> 
struct TupleElement<idx, Tuple<T, Ts...>> : public TupleElement<idx - 1, Tuple<Ts...>> {
};

template<size_t idx, class ... Ts> decltype(auto) get(Tuple<Ts...>& tuple) {
    using TupleT = typename TupleElement<idx, Tuple<Ts...>>::TupleT;
    return (static_cast<TupleT&>(tuple).value);
}

int main() {
    Tuple<int, string> tuple;
    get<0>(tuple) = 1;
    get<1>(tuple) = "hello";
    assert(get<0>(tuple) == 1);
    assert(get<1>(tuple) == "hello");
    return 0;
}