#include "true.hpp"

std::string True::inspect(const bool) const {
    return "true";
}

Value::Type True::type() const {
    return Value::Type::True;
}

bool True::isTrue() const { 
    return true;
}

bool True::operator==(const True&) const {
    return true;
}

bool True::operator!=(const True&) const {
    return false;
}
