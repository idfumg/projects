#include <iostream>
#include <string>

#include "readline.hpp"

const auto HISTORY_PATH = "history.txt";

std::string Read(const std::string& input) {
    return input;
}

std::string Eval(const std::string& input) {
    return input;
}

std::string Print(const std::string& input) {
    return input;
}

std::string REP(const std::string& input) {
    const auto ast = Read(input);
    const auto result = Eval(ast);
    return Print(result);
}

int main() {
    const auto linenoise = Linenoise::New(HISTORY_PATH);

    std::string input;
    while (true) {
        if (const auto quit = linenoise.Readline("user> ", input)) {
            break;
        }

        std::cout << REP(input) << std::endl;
        linenoise.AddHistory(input);
    }

    linenoise.SaveHistory();

    return 0;
}