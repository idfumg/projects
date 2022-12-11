#include "tokenizer.hpp"
#include <iostream>

Tokenizer::Tokenizer(const std::string& input) : input(input), idx(0) {
}

bool Tokenizer::isDone() const {
    return idx >= input.size();
}

std::tuple<std::string, bool> Tokenizer::Next() {
    const auto& view = input;

    const auto readComment = [&](){
        while (idx < input.size()) {
            const char c = input[idx];
            if (c == '\n') {
                ++idx;
                break;
            }
            ++idx;
        }
    };

    const auto readString = [&]()->std::tuple<std::string, bool>{
        size_t start = idx;
        
        for (++idx; idx < input.size(); ++idx) {
            const char c = input[idx];
            switch (c) {
                case '\\': {
                    ++idx;
                    break;
                }
                case '"': {
                    ++idx;
                    return {view.substr(start, idx - start), false};
                }
            }
        }
        std::cerr << "EOF\n";
        return {view.substr(start, idx - start), true};
    };

    const auto readInteger = [&]()->std::tuple<std::string, bool>{
        size_t start = idx;
        for (; idx < input.size(); ++idx) {
            switch (input[idx]) {
                case '-': case '+': case '.':
                case '0': case '1': case '2': case '3': case '4':
                case '5': case '6': case '7': case '8': case '9': break;
                default: return {view.substr(start, idx - start), false};
            }
        }
        return {view.substr(start, idx - start), false};
    };

    const auto readAtom = [&]()->std::tuple<std::string, bool>{
        size_t start = idx;
        for (; idx < input.size(); ++idx) {
            const char c = input[idx];
            switch (c) {
                case '-': case '+': {
                    if (idx + 1 < input.size() && std::isdigit(input[idx + 1])) {
                        return readInteger();
                    }
                    break;
                }
                case ' ': case '\t': case '\n': case '[': case ']':
                case '{': case '}':  case '(':  case ')': case '"':
                case '`': case ',':  case ';':  case '~': case '\'': {
                    return {view.substr(start, idx - start), false};
                }
            }
        }
        return {view.substr(start, idx - start), false};
    };

    const auto readUnquote = [&]()->std::tuple<std::string, bool>{
        if (idx + 1 < input.size() && input[idx + 1] == '@') {
            const auto x = view.substr(idx, 2);
            idx += 2;
            return {x, false};
        }
        return {view.substr(idx++, 1), false};
    };

    const auto readOneCharSymbol = [&]()->std::tuple<std::string, bool>{
        return {view.substr(idx++, 1), false};
    };

    while (idx < input.size()) {
        const char c = input.at(idx);

        switch (c) {
            case ' ': case '\t': case '\n': case ',': break;
            case '~':                                          return readUnquote();
            case '"':                                          return readString();
            case ';':                                          readComment();
            case '[': case ']':  case '(': case ')': case '{': 
            case '}': case '\'': case '`': case '^': case '@': return readOneCharSymbol();
            case '0': case '1':  case '2': case '3': case '4':
            case '5': case '6':  case '7': case '8': case '9': return readInteger();
            default:                                           return readAtom();
        }

        ++idx;
    }
    return {"", true};
}