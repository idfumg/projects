#pragma once

#include <string>

class Tokenizer {
private:
    const std::string& input;
    size_t idx;

public:
    Tokenizer(const std::string& input);
    bool isDone() const;
    std::tuple<std::string, bool> Next();
};