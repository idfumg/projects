#include "reader.hpp"
#include "value.hpp"
#include "tokenizer.hpp"

std::vector<std::string> Tokenize(const std::string& input) {
    std::vector<std::string> ans;

    Tokenizer tokenizer(input);
    for (;;) {
        const auto [token, err] = tokenizer.Next();
        if (err) {
            break;
        }
        ans.push_back(token);
    }
    assert(tokenizer.isDone());

    return ans;
}

Reader::Reader(const std::vector<std::string>& tokens) : tokens(tokens) {

}

bool Reader::HasNext() const {
    return idx < tokens.size();
}

std::string Reader::Next() {
    if (!HasNext()) {
        return "";
    }
    const auto ans = Peek();
    ++idx;
    return ans;
}

std::string Reader::Peek() const {
    if (!HasNext()) {
        return "";
    }
    return tokens[idx];
}

bool IsIntegerWithASign(Reader& reader) {
    const auto token = reader.Peek();
    return (token[0] == '-' || token[0] == '+') && token.size() > 1 && std::isdigit(token[1]);
}

bool IsIntegerWithoutASign(Reader& reader) {
    const auto token = reader.Peek();
    return std::isdigit(token[0]);
}

bool IsInteger(Reader& reader) {
    return IsIntegerWithASign(reader) || IsIntegerWithoutASign(reader);
}

bool IsTrue(Reader& reader) {
    const auto token = reader.Peek();
    return token == "true";
}

bool IsFalse(Reader& reader) {
    const auto token = reader.Peek();
    return token == "false";
}

bool IsNil(Reader& reader) {
    const auto token = reader.Peek();
    return token == "nil";
}

bool IsString(Reader& reader) {
    const auto token = reader.Peek();
    return token[0] == '\"';
}

bool IsKeyword(Reader& reader) {
    const auto token = reader.Peek();
    return token[0] == ':' || (token[0] == '-' && token.size() > 1 && !std::isdigit(token[1]));
}

Value* ReadInteger(Reader& reader) {
    const auto token = reader.Next();
    std::int64_t ans = 0;
    int8_t negative = 1;
    bool stopReading = false;
    for (std::size_t i = 0; i < token.size(); ++i) {
        const char c = token[i];

        if (i == 0 && c == '-') negative = -1;
        else if (i == 0 && c == '+') negative = 1;
        else if (!stopReading && std::isdigit(c)) ans = ans * 10 + c - '0';
        else if (c == '.') {stopReading = true; continue; }
        else break;
    }
    return new Integer(ans * negative);
}

Value* ReadTrue(Reader& reader) {
    reader.Next();
    return new True();
}

Value* ReadFalse(Reader& reader) {
    reader.Next();
    return new False();
}

Value* ReadNil(Reader& reader) {
    reader.Next();
    return new Nil();
}

Value* ReadSymbol(Reader& reader) {
    const auto token = reader.Next();
    return new Symbol(token);
}

Value* ReadString(Reader& reader) {
    const auto token = reader.Next();
    assert(token.size() >= 2);
    if (token.size() == 2) return new String("");

    const std::string str = token.substr(1, token.size() - 2);
    std::string s = "";
    for (std::size_t i = 0; i < str.size(); ++i) {
        if (str[i] == '"') {
            s += '\\';
            s += str[i];
        } else if (str[i] == '\\') {
            if (str[++i] == 'n') {
                s += '\n';
            } else {
                s += str[i];
            }
        } else {
            s += str[i];
        }
    }
    return new String(s);
}

Value* ReadKeyword(Reader& reader) {
    const auto token = reader.Next();
    return new Keyword(token);
}

Value* ReadAtom(Reader& reader) {
    if      (IsInteger(reader)) return ReadInteger(reader);
    else if (IsTrue(reader))    return ReadTrue(reader);
    else if (IsFalse(reader))   return ReadFalse(reader);
    else if (IsNil(reader))     return ReadNil(reader);
    else if (IsString(reader))  return ReadString(reader);
    else if (IsKeyword(reader)) return ReadKeyword(reader);
    return ReadSymbol(reader);
}

Value* ReadList(Reader& reader) {
    reader.Next(); // (

    auto* ans = new List();

    for (; reader.HasNext();) {
        if (const auto token = reader.Peek(); token == ")") {
            reader.Next(); // )
            return ans;
        }
        ans->push(ReadForm(reader));
    }

    std::cerr << "EOF\n";
    return ans;
}

Value* ReadVector(Reader& reader) {
    reader.Next(); // [

    auto* ans = new Vector();

    for (; reader.HasNext();) {
        if (const auto token = reader.Peek(); token == "]") {
            reader.Next(); // ]
            return ans;
        }
        ans->push(ReadForm(reader));
    }

    std::cerr << "EOF\n";
    return ans;
}

Value* ReadHashMap(Reader& reader) {
    reader.Next(); // {

    auto* ans = new HashMap();

    for (; reader.HasNext();) {
        if (const auto token = reader.Peek(); token == "}") {
            reader.Next(); // }
            return ans;
        }

        const auto k = ReadForm(reader);

        if (const auto token = reader.Peek(); token == "}") {
            std::cerr << "EOF\n";
            reader.Next(); // }
            return ans;
        }

        const auto v = ReadForm(reader);

        ans->insert(k, v);
    }

    std::cerr << "EOF\n";
    return ans;
}

Value* ReadQuoted(Reader& reader) {
    reader.Next(); // '

    auto* ans = new List();
    ans->push(new Symbol("quote"));
    ans->push(ReadForm(reader));

    return ans;
}

Value* ReadQuasiQuoted(Reader& reader) {
    reader.Next(); // `

    auto* ans = new List();
    ans->push(new Symbol("quasiquote"));
    ans->push(ReadForm(reader));

    return ans;
}

Value* ReadUnquoted(Reader& reader) {
    assert(reader.Peek()[0] == '~');

    auto* ans = new List();
    if (reader.Peek()[1] == '@') {
        reader.Next(); // ~
        ans->push(new Symbol("splice-unquote"));
    } else {
        reader.Next(); // ~
        ans->push(new Symbol("unquote"));
    }
    ans->push(ReadForm(reader));

    return ans;
}

Value* ReadDeref(Reader& reader) {
    reader.Next(); // @

    auto* ans = new List();
    ans->push(new Symbol("deref"));
    ans->push(ReadForm(reader));

    return ans;
}

Value* ReadMeta(Reader& reader) {
    reader.Next(); // ^

    const auto meta = ReadForm(reader);
    const auto value = ReadForm(reader);
    auto* ans = new List();
    ans->push(new Symbol("with-meta"));
    ans->push(value);
    ans->push(meta);

    return ans;
}

Value* ReadForm(Reader& reader) {
    if (!reader.HasNext()) {
        return nullptr;
    }

    const auto token = reader.Peek();
    switch (token[0]) {
        case '(':  return ReadList(reader);
        case '[':  return ReadVector(reader);
        case '{':  return ReadHashMap(reader);
        case '\'': return ReadQuoted(reader);
        case '`':  return ReadQuasiQuoted(reader);
        case '~':  return ReadUnquoted(reader);
        case '@':  return ReadDeref(reader);
        case '^':  return ReadMeta(reader);
        default:   return ReadAtom(reader);
    }

    return nullptr;
}

Value* ReadInput(const std::string& input) {
    Reader reader(Tokenize(input));
    return ReadForm(reader);
}

