#include <iostream>
#include <string>
#include <optional>
#include <unordered_map>

#include "readline.hpp"
#include "reader.hpp"
#include "printer.hpp"
#include "types.hpp"
#include "core.hpp"

const auto HISTORY_PATH = "history.txt";

std::tuple<Value*, bool> EvalAst(Value* ast, Env& env);

std::tuple<Value*, bool> Read(const std::string& input) {
    const auto ast = ReadInput(input);
    return {ast, ast == nullptr};
}

std::tuple<Fn*, std::vector<Value*>> SplitIntoHeadAndTail(List* L) {
    const auto head = L->get().at(0)->asFn();
    std::vector<Value*> tail(L->size() - 1);
    for (std::size_t i = 0; i < L->size() - 1; ++i) {
        tail[i] = L->get().at(i + 1);
    }
    return {head, tail};
}

std::tuple<Value*, bool> Eval(Value* ast, Env& env) {
    if (ast->type() != Value::Type::List) {
        return EvalAst(ast, env);
    } else if (ast->asList()->isEmpty()) {
        return {ast, false};
    } else {
        const auto L = ast->asList();
        const auto head = L->get().at(0);
        if (head->isSymbol() && head->asSymbol()->matches("def!")) {
            const auto key = L->get().at(1);
            const auto [value, err] = Eval(L->get().at(2), env);
            if (err) return {value, err};
            env.set(key->asSymbol(), value);
            return {key, false};
        } else if (head->isSymbol() && head->asSymbol()->matches("let*")) {
            const auto newEnv = new Env(&env);
            const auto bindings = L->get().at(1)->asList();
            for (std::size_t i = 0; i < bindings->size(); i += 2) {
                const auto key = bindings->get().at(i);
                if (i + 1 >= bindings->size()) {
                    return {new Exception("Error! Wrong number of params"), true};
                }
                const auto [value, err] = Eval(bindings->get().at(i + 1), *newEnv);
                if (err) return {value, err};
                newEnv->set(key->asSymbol(), value);
                return Eval(L->get().at(2), *newEnv);
            }
        } else {
            const auto [value, err] = EvalAst(ast, env);
            if (err) return {value, err};
            const auto [head, tail] = SplitIntoHeadAndTail(value->asList());
            return head->get()(value->asList()->size() - 1, tail.data());
        }
    }
    return {ast, false};
}

std::tuple<Value*, bool> EvalAst(Value* ast, Env& env) {
    switch (ast->type()) {
        case Value::Type::Symbol: {
            return env.get(ast->asSymbol());
        }

        case Value::Type::List: {
            const auto ans = new List();
            for (const auto& value : *ast->asList()) {
                const auto [res, err] = Eval(value, env);
                if (err) return {res, true};
                ans->push(res);
            }
            return {ans, false};
        }

        case Value::Type::Vector: {
            const auto ans = new Vector();
            for (const auto& value : *ast->asVector()) {
                const auto [res, err] = Eval(value, env);
                if (err) return {res, true};
                ans->push(res);
            }
            return {ans, false};
        }

        case Value::Type::HashMap: {
            const auto ans = new HashMap();
            for (const auto& [key, value] : *ast->asHashMap()) {
                const auto [res, err] = Eval(value, env);
                if (err) return {res, true};
                ans->insert(key, res);
            }
            return {ans, false};
        }

        default: {
            return {ast, false};
        }
    }
    return {ast, false};
}

void Print(const std::string& input) {
    std::cout << input << std::endl;
}

bool REP(const std::string& input, Env& env) {
    if (const auto [ast, err] = Read(input); err) {
        return false;
    } else if (const auto [result, err] = Eval(ast, env); err) {
        if (result && result->type() == Value::Type::Exception) {
            std::cerr << result->inspect(true) << std::endl;
            return false;
        } else {
            std::cerr << "Unexpected eval error happened" << std::endl;
            return false;
        }
    } else {
        Print(result->inspect(true));
    }
    return true;
}

std::tuple<std::string, bool> GetLine() {
    std::string input;

    const auto linenoise = Linenoise::New(HISTORY_PATH);
        
    if (const auto quit = linenoise.Readline("user> ", input)) {
        return {"", true};
    }

    linenoise.AddHistory(input);
    linenoise.SaveHistory();

    return {input, false};
}

int main() {
    Env env = Env(nullptr);
    env.set(new Symbol("+"), new Fn(Add));
    env.set(new Symbol("-"), new Fn(Sub));
    env.set(new Symbol("*"), new Fn(Mul));
    env.set(new Symbol("/"), new Fn(Div));

    for (;;) {
        const auto [line, ok] = GetLine();
        if (ok) {
            break;
        }
        REP(line, env);
    }

    return 0;
}
