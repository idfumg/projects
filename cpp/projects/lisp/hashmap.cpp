#include "hashmap.hpp"

void HashMap::insert(Value* k, Value* v) {
    map[k] = v;
}

const Value* HashMap::get(Value* k) const {
    if (map.count(k)) {
        return map.at(k);
    }
    return {};
}

std::size_t HashMap::size() const {
    return map.size();
}

std::string HashMap::inspect(const bool printReadably) const {
    std::string ans = "{";
    for (const auto [k, v] : map) {
        assert(k && v);
        ans += k->inspect(printReadably) + ' ';
        ans += v->inspect(printReadably) + ' ';
    }
    if (map.size() > 0) {
        ans.back() = '}';
    } else {
        ans += '}';
    }
    return ans;
}

Value::Type HashMap::type() const {
    return Value::Type::HashMap;
}

bool HashMap::operator==(const HashMap& b) const {
    if (size() != b.size()) return false;
    for (const auto& [k, v] : this->map) {
        if (!b.map.count(k) || map.at(k) != b.map.at(k)) {
            return false;
        }
    }
    return true;
}

bool HashMap::operator!=(const HashMap& b) const {
    return !(*this == b);
}

bool HashMap::isHashMap() const {
    return true;
}