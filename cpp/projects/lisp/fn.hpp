#pragma once

#include "value.hpp"

#include <functional>

using Func = std::function<std::tuple<Value*, bool>(const std::size_t, Value*const*)>;

class Fn : public Value {
public:
    Fn(const Func& func);

    Func get() const;

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

private:
    Func fn;
};
