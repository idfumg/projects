#pragma once

#include "value.hpp"

#include <vector>
#include <iostream>

class List : public Value {
public:
    List(){}

    List(const std::size_t argc, Value*const* argv) {
        for (std::size_t i = 0; i < argc; ++i) {
            list.push_back(argv[i]);
        }
    }

    void push(Value* value);
    auto begin() { return std::begin(list); }
    auto end() { return std::end(list); }
    bool isEmpty() const { return list.empty(); }
    std::size_t size() const { return list.size(); }
    Value* at(const std::size_t i) const { return list.at(i); }
    virtual bool isList() const override;
    virtual bool isLikeList() const override;
    bool operator==(const List& rhs) const;
    bool operator!=(const List& rhs) const;
    
    std::vector<Value*>& get() { return list; }
    const std::vector<Value*>& get() const { return list; };

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

private:
    std::vector<Value*> list{};
};
