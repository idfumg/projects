#include "fn.hpp"

Fn::Fn(const Func& fn) : fn(fn){

}

Func Fn::get() const {
    return fn;
}

std::string Fn::inspect(const bool) const {
    return "#<function>";
}

Value::Type Fn::type() const {
    return Value::Type::Fn;
}
