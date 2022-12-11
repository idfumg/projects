#pragma once

#include "value.hpp"

#include <unordered_map>

struct HashMapHash {
    std::size_t operator()(Value* key) const noexcept {
        return std::hash<std::string>{}(key->inspect(false));
    }
};

struct HashMapPred {
    bool operator()(Value* lhs, Value* rhs) const {
        return lhs->inspect(false) == rhs->inspect(false);
    }
};

class HashMap : public Value {
public:
    HashMap(){}

    void insert(Value* k, Value* v);
    const Value* get(Value* k) const;
    std::size_t size() const;
    auto begin() { return std::begin(map); }
    auto end() { return std::end(map); }

    bool operator==(const HashMap& b) const;
    bool operator!=(const HashMap& b) const;
    bool isHashMap() const override;

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

private:
    std::unordered_map<Value*, Value*, HashMapHash, HashMapPred> map;
};
