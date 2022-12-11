#include "symbol.hpp"

const std::string& Symbol::get() const {
    return m_str;
}

bool Symbol::matches(const std::string& str) const {
    return str == m_str;
}

std::string Symbol::inspect(const bool) const {
    return m_str;
}

Value::Type Symbol::type() const {
    return Value::Type::Symbol;
}

bool Symbol::operator==(const Symbol& b) const {
    return get() == b.get();
}

bool Symbol::operator!=(const Symbol& b) const {
    return !(*this == b);
}