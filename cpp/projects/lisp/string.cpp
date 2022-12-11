#include "string.hpp"

bool String::isString() const {
    return true;
}

bool String::operator==(const String& rhs) const {
    return str == rhs.str;
}

bool String::operator!=(const String& rhs) const {
    return !(*this == rhs);
}

std::string String::inspect(const bool printReadably) const {
    if (printReadably) {
        std::string s = "\"";

        for (const char c : str) {
            switch (c) {
                case '"': 
                    s += '\\';
                    s += c;
                    break;
                case '\\':
                    s += '\\';
                    s += '\\';
                    break;
                case '\n':
                    s += '\\';
                    s += 'n';
                    break;
                default:
                    s += c;
            }
        }

        return s + "\"";
    }
    return str;
}

Value::Type String::type() const {
    return Value::Type::String;
}

