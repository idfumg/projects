#pragma once

#include "value.hpp"

class Keyword : public Value {
public:
    Keyword(){}

    Keyword(std::string_view str) : str(str) {}

    std::size_t size() const { return str.size(); }
    virtual bool isKeyword() const override;

    bool operator==(const Keyword& rhs) const;
    bool operator!=(const Keyword& rhs) const;
    
    std::string& get() { return str; }
    const std::string& get() const { return str; };

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

private:
    std::string str{};
};
