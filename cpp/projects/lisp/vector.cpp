#include "vector.hpp"

void Vector::push(Value* value) {
    assert(value);
    list.push_back(value);
}

bool Vector::isLikeList() const {
    return true;
}

bool Vector::isVector() const {
    return true;
}

bool Vector::operator==(const Vector& b) const {
    if (size() != b.size()) return false;
    for (std::size_t i = 0; i < size(); ++i) {
        if (at(i) != b.at(i)) {
            return false;
        }
    }
    return true;
}

bool Vector::operator!=(const Vector& b) const {
    return !(*this == b);
}

std::string Vector::inspect(const bool printReadably) const {
    // std::string ans = printReadably ? "\"[" : "[";
    std::string ans = "[";
    for (const auto* value : list) {
        assert(value);
        ans += value->inspect(printReadably) + ' ';
    }
    if (!list.empty()) {
        ans.back() = ']';
    } else {
        ans += ']';
    }
    // if (printReadably) ans.push_back('\"');
    return ans;
}

Value::Type Vector::type() const {
    return Value::Type::Vector;
}
