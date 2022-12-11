#include "exception.hpp"

const std::string& Exception::get() const {
    return msg;
}

std::string Exception::inspect(const bool) const {
    return "<exception (" + msg + ")>";
}

Value::Type Exception::type() const {
    return Value::Type::Exception;
}