#include "nil.hpp"

std::string Nil::inspect(const bool) const {
    return "nil";
}

Value::Type Nil::type() const {
    return Value::Type::Nil;
}

bool Nil::isTruth() const {
    return false;
}

bool Nil::isNil() const {
    return true;
}

bool Nil::operator==(const Nil&) const {
    return true;
}

bool Nil::operator!=(const Nil&) const {
    return false;
}
