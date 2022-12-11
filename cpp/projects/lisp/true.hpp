#pragma once

#include "value.hpp"

class True : public Value {
public:
    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;
    bool isTrue() const override;
    bool operator==(const True& b) const;
    bool operator!=(const True& b) const;
};
