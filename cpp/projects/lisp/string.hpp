#pragma once

#include "value.hpp"

class String : public Value {
public:
    String(){}

    String(std::string_view str) : str(str) {}

    auto begin() { return std::begin(str); }
    auto end() { return std::end(str); }
    bool isEmpty() const { return str.empty(); }
    std::size_t size() const { return str.size(); }
    virtual bool isString() const override;

    bool operator==(const String& rhs) const;
    bool operator!=(const String& rhs) const;
    
    std::string& get() { return str; }
    const std::string& get() const { return str; };

    std::string inspect(const bool printReadably) const override;
    virtual Type type() const override;

private:
    std::string str{};
};
