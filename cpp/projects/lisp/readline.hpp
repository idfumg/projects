#pragma once

#include <string>
#include <string_view>

class Linenoise {
private:
    std::string m_filename;

public:
    static Linenoise New(const std::string& filename);
    void AddHistory(const std::string_view& input) const;
    void SaveHistory() const;
    bool Readline(const std::string_view& prompt, std::string& input) const;
};