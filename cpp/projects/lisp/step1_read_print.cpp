#include <iostream>
#include <string>
#include <optional>

#include "readline.hpp"
#include "reader.hpp"
#include "printer.hpp"

const auto HISTORY_PATH = "history.txt";

struct Parsed {
    Value* ast;
    bool isError;
};

Parsed Read() {
    const auto linenoise = Linenoise::New(HISTORY_PATH);

    std::string input;
        
    if (const auto quit = linenoise.Readline("user> ", input)) {
        return Parsed{
            .ast = {},
            .isError = true,
        };
    }

    linenoise.AddHistory(input);
    linenoise.SaveHistory();

    return Parsed{
        .ast = ReadInput(input),
        .isError = false,
    };
}

std::string Eval(Value* ast) {
    return PrintAst(ast, true);
}

void Print(const std::string& input) {
    std::cout << input << std::endl;
}

bool REP() {
    const auto read = Read();
    if (read.isError) {
        return false;
    }
    if (!read.ast) return true;
    const auto result = Eval(read.ast);
    Print(result);
    return true;
}

int main() {
    for (; REP();) {
    }

    return 0;
}
