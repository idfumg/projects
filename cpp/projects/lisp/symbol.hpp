#pragma once

#include "value.hpp"

class Symbol : public Value {
public:
    Symbol(std::string_view str) : m_str(str) {}

    const std::string& get() const;
    bool matches(const std::string& str) const;
    bool isSymbol() const override { return true; }

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

    bool operator==(const Symbol& b) const;
    bool operator!=(const Symbol& b) const;

private:
    std::string m_str;
};
