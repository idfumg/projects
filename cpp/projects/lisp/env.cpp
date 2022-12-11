#include "env.hpp"
#include "symbol.hpp"
#include "exception.hpp"
#include "list.hpp"

Env::Env(Env* outer) : outer(outer) {
    
}

Env::Env(Env* outer, List* binds, List* exprs) : outer(outer) {
    for (size_t i = 0 ; i < binds->size(); ++i) {
        const auto key = binds->at(i)->asSymbol();
        if (key->get()[0] == '&') {
            const auto key = binds->at(i + 1)->asSymbol();
            const auto value = new List();
            for (std::size_t j = i; j < exprs->size(); ++j) {
                value->push(exprs->at(j));
            }
            set(key, value);
            return;
        } else {
            const auto value = exprs->at(i);
            set(key, value);
        }
    }
}

void Env::set(Symbol* key, Value* value) {
    data[key] = value;
}

Env* Env::find(Symbol* key) const {
    const auto it = data.find(key);
    if (it != data.end()) {
        return const_cast<Env*>(this);
    } else if (outer) {
        return outer->find(key);
    } else {
        return nullptr;
    }
}

std::tuple<Value*, bool> Env::get(Symbol* key) const {
    const auto env = find(key);
    if (!env) {
        return {new Exception("Error! Atom "  + key->get() + " not found"), true};
    }
    return {env->data[key], false};
}