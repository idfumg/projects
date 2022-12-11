#include "integer.hpp"
#include <iostream>

std::int64_t Integer::get() const {
    return value;
}

bool Integer::operator==(const Integer& b) const {
    return get() == b.get();
}

bool Integer::operator!=(const Integer& b) const {
    return !(*this == b);
}

bool Integer::operator<(const Integer& b) const {
    return get() < b.get();
}

bool Integer::operator<=(const Integer& b) const {
    return get() <= b.get();
}

bool Integer::operator>(const Integer& b) const {
    return get() > b.get();
}

bool Integer::operator>=(const Integer& b) const {
    return get() >= b.get();
}

bool Integer::isInteger() const {
    return true;
}

std::string Integer::inspect(const bool) const {
    return std::to_string(value);
}

Value::Type Integer::type() const {
    return Value::Type::Integer;
}
