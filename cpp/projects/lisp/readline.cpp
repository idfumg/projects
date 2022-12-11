#include "readline.hpp"

#include "linenoise.hpp"

Linenoise Linenoise::New(const std::string& filename) {
    // Setup completion words every time when a user types
    linenoise::SetCompletionCallback([](const char* editBuffer, std::vector<std::string>& completions) {
        if (editBuffer[0] == 'h') {
            completions.push_back("hello");
            completions.push_back("hello there");
        }
    });

    // Enable the multi-line mode
    linenoise::SetMultiLine(true);

    // Set max length of the history
    linenoise::SetHistoryMaxLen(20);

    // Load history
    linenoise::LoadHistory(filename.c_str());

    Linenoise ans;
    ans.m_filename = filename;
    return ans;
}

void Linenoise::AddHistory(const std::string_view& input) const {
    linenoise::AddHistory(input.data());
}

void Linenoise::SaveHistory() const {
    linenoise::SaveHistory(m_filename.c_str());
}

bool Linenoise::Readline(const std::string_view& prompt, std::string& input) const {
    return linenoise::Readline(prompt.data(), input);
}