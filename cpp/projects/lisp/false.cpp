#include "false.hpp"

std::string False::inspect(const bool) const {
    return "false";
}

Value::Type False::type() const {
    return Value::Type::False;
}

bool False::isTruth() const {
    return false;
}

bool False::isFalse() const {
    return true;
}

bool False::operator==(const False&) const {
    return true;
}

bool False::operator!=(const False&) const {
    return false;
}
