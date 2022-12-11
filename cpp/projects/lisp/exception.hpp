#pragma once

#include "value.hpp"

class Exception : public Value {
public:
    Exception(const std::string& msg) : msg(msg) {}

    const std::string& get() const;

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

private:
    std::string msg;
};
