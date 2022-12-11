#pragma once

#include "value.hpp"

class False : public Value {
public:
    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;
    bool isTruth() const override;
    bool isFalse() const override;
    bool operator==(const False& b) const;
    bool operator!=(const False& b) const;
};
