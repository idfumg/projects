#pragma once

#include <iostream>
#include <string>
#include <vector>

#include "types.hpp"

class Reader {
private:
    std::vector<std::string> tokens;
    size_t idx{};

public:
    Reader(const std::vector<std::string>& tokens);

    bool HasNext() const;
    std::string Next();
    std::string Peek() const;
};

std::vector<std::string> Tokenize(const std::string& input);
Value* ReadForm(Reader& reader);
Value* ReadInput(const std::string& input);