#pragma once

#include "value.hpp"

#include <vector>

class Vector : public Value {
public:
    Vector(){}

    void push(Value* value);
    Value* at(const std::size_t i) const { return list.at(i); }
    std::size_t size() const { return list.size(); }
    std::vector<Value*> get() { return list; }
    const std::vector<Value*> get() const { return list; }
    virtual bool isLikeList() const override;
    virtual bool isVector() const override;
    bool operator==(const Vector& rhs) const;
    bool operator!=(const Vector& rhs) const;

    auto begin() { return std::begin(list); }
    auto end() { return std::end(list); }

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

private:
    std::vector<Value*> list{};
};
