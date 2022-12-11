#pragma once

#include "value.hpp"

class Nil : public Value {
public:
    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;
    bool isTruth() const override;
    bool isNil() const override;

    bool operator==(const Nil& b) const;
    bool operator!=(const Nil& b) const;
};
