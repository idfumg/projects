#include "keyword.hpp"

bool Keyword::isKeyword() const {
    return true;
}

bool Keyword::operator==(const Keyword& rhs) const {
    return str == rhs.str;
}

bool Keyword::operator!=(const Keyword& rhs) const {
    return !(*this == rhs);
}

std::string Keyword::inspect(const bool) const {
    return str;
}

Value::Type Keyword::type() const {
    return Value::Type::Keyword;
}

