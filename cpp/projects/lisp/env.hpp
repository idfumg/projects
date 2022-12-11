#pragma once

#include "hashmap.hpp"
#include "list.hpp"

class Env {
public:
    Env(Env* outer);
    Env(Env* outer, List* binds, List* exprs);

    void set(Symbol* key, Value* value);
    Env* find(Symbol* key) const;
    std::tuple<Value*, bool> get(Symbol* key) const;

private:
    Env* outer{};
    std::unordered_map<Symbol*, Value*, HashMapHash, HashMapPred> data;
};