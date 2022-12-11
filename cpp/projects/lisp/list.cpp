#include "list.hpp"

void List::push(Value* value) {
    assert(value);
    list.push_back(value);
}

bool List::isList() const {
    return true;
}

bool List::isLikeList() const {
    return true;
}

bool List::operator==(const List& b) const {
    if (size() != b.size()) return false;
    for (std::size_t i = 0; i < size(); ++i) {
        if (!IsEqual(at(i), b.at(i))) {
            return false;
        }
    }
    return true;
}

bool List::operator!=(const List& b) const {
    return !(*this == b);
}

std::string List::inspect(const bool printReadably) const {
    std::string ans = "(";
    for (const auto* value : list) {
        assert(value);
        ans += value->inspect(printReadably) + ' ';
    }
    if (!list.empty()) {
        ans.back() = ')';
    } else {
        ans += ')';
    }
    return ans;
}

Value::Type List::type() const {
    return Value::Type::List;
}
