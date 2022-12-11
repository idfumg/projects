#pragma once

#include "value.hpp"

class Integer : public Value {
public:
    Integer(const std::int64_t value) : value(value){}

    std::int64_t get() const;
    bool operator==(const Integer& b) const;
    bool operator!=(const Integer& b) const;
    bool operator<(const Integer& b) const;
    bool operator<=(const Integer& b) const;
    bool operator>(const Integer& b) const;
    bool operator>=(const Integer& b) const;
    bool isInteger() const override;

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

private:
    std::int64_t value{};
};
