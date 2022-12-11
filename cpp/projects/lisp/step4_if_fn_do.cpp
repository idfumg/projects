#include <exception>
#include <iostream>
#include <string>
#include <optional>
#include <unordered_map>

#include "core.hpp"
#include "exception.hpp"
#include "readline.hpp"
#include "reader.hpp"
#include "printer.hpp"
#include "types.hpp"

const auto HISTORY_PATH = "history.txt";

std::tuple<Value*, bool> EvalAst(Value* ast, Env* env);
std::tuple<Value*, bool> Eval(Value* ast, Env* env);

std::tuple<Value*, bool> Read(const std::string& input) {
    const auto ast = ReadInput(input);
    return {ast, ast == nullptr};
}

std::tuple<Fn*, std::vector<Value*>> SplitIntoHeadAndTail(List* L) {
    const auto head = L->at(0)->asFn();
    std::vector<Value*> tail(L->size() - 1);
    for (std::size_t i = 0; i < L->size() - 1; ++i) {
        tail[i] = L->at(i + 1);
    }
    return {head, tail};
}

std::tuple<Value*, bool> EvalDef(Value* ast, Env* env) {
    const auto L = ast->asList();
    const auto key = L->at(1);
    const auto [value, err] = Eval(L->at(2), env);
    if (err) return {value, err};
    env->set(key->asSymbol(), value);
    return {value, false};
}

std::tuple<Value*, bool> EvalLet(Value* ast, Env* env) {
    const auto L = ast->asList();
    const auto newEnv = new Env(env);
    const auto bindings = L->at(1)->asList();
    for (std::size_t i = 0; i < bindings->size(); i += 2) {
        const auto key = bindings->at(i);
        if (i + 1 >= bindings->size()) {
            return {new Exception("Error! Wrong number of params"), true};
        }
        const auto [value, err] = Eval(bindings->at(i + 1), newEnv);
        if (err) return {value, err};
        newEnv->set(key->asSymbol(), value);
    }
    return Eval(L->get().at(2), newEnv);
}

std::tuple<Value*, bool> EvalList(Value* ast, Env* env) {
    const auto [value, err] = EvalAst(ast, env);
    if (err) return {value, err};
    const auto [head, tail] = SplitIntoHeadAndTail(value->asList());
    return head->get()(value->asList()->size() - 1, tail.data());
}

std::tuple<Value*, bool> EvalDo(Value* ast, Env* env) {
    const auto L = ast->asList();
    Value* ans = nullptr;
    for (std::size_t i = 1; i < L->size(); ++i) {
        const auto [res, err] = Eval(L->at(i), env);
        if (err) return {res, err};
        ans = res;
    }
    return {ans, false};
}

std::tuple<Value*, bool> EvalIf(Value* ast, Env* env) {
    const auto L = ast->asList();
    const auto condition = L->at(1);
    const auto trueExpr = L->at(2);
    const auto falseExpr = L->size() >= 4 ? L->at(3) : new Nil();
    const auto [res, err] = Eval(condition, env);
    if (err) return {res, err};
    if (res->isTruth()) return Eval(trueExpr, env);
    return Eval(falseExpr, env);
}

bool IsThereAmpExist(const std::string& s) {
    for (const char c : s) {
        if (c == '&') {
            return true;
        }
    }
    return false;
}

bool IsThereAnAmp(List* binds) {
    for (std::size_t i = 0; i < binds->size(); ++i) {
        if (IsThereAmpExist(binds->at(i)->asSymbol()->get())) {
            return true;
        }
    }
    return false;
}

std::tuple<Value*, bool> EvalFn(Value* ast, Env* env) {
    const auto L = ast->asList();
    if (L->size() != 3) {
        return {new Exception("Error! Wrong function definition"), true};
    }

    const auto closure = [=](const std::size_t argc, Value*const* argv)->std::tuple<Value*, bool>{
        const auto body = L->at(2);
        const auto binds = L->at(1)->asList();
        const auto exprs = new List(argc, argv);
        if (!binds || !exprs || (binds->size() != exprs->size() && !IsThereAnAmp(binds))) {
            return {new Exception("Error! Not enough function params"), true};
        }
        Env* tmpEnv = new Env(*env);
        Env* funcEnv = new Env(tmpEnv, binds, exprs);
        const auto [res, err] = Eval(body, funcEnv);
        return {res, err};
    };

    const auto [res, err] = Eval(new Fn(closure), env);
    return {res, err};
}

bool IsHeadMatchSymbol(Value* ast, const std::string& name) {
    const auto L = ast->asList();
    const auto head = L->at(0);
    return head->isSymbol() && head->asSymbol()->matches(name);
}

bool IsDef(Value* ast) {
    return IsHeadMatchSymbol(ast, "def!");
}

bool IsLet(Value* ast) {
    return IsHeadMatchSymbol(ast, "let*");
}

bool IsDo(Value* ast) {
    return IsHeadMatchSymbol(ast, "do");
}

bool IsIf(Value* ast) {
    return IsHeadMatchSymbol(ast, "if");
}

bool IsFn(Value* ast) {
    return IsHeadMatchSymbol(ast, "fn*");
}

std::tuple<Value*, bool> Eval(Value* ast, Env* env) {
    if (ast->type() != Value::Type::List) return EvalAst(ast, env);
    else if (ast->asList()->isEmpty())    return {ast, false};
    else if (IsDef(ast))                  return EvalDef(ast, env);
    else if (IsLet(ast))                  return EvalLet(ast, env);
    else if (IsDo(ast))                   return EvalDo(ast, env);
    else if (IsIf(ast))                   return EvalIf(ast, env);
    else if (IsFn(ast))                   return EvalFn(ast, env);
    else                                  return EvalList(ast, env);
    return {ast, false};
}

std::tuple<Value*, bool> EvalAstSymbol(Value* ast, Env* env) {
    return env->get(ast->asSymbol());
}

std::tuple<Value*, bool> EvalAstList(Value* ast, Env* env) {
    const auto ans = new List();
    for (const auto& value : *ast->asList()) {
        const auto [res, err] = Eval(value, env);
        if (err) return {res, true};
        ans->push(res);
    }
    return {ans, false};
}

std::tuple<Value*, bool> EvalAstVector(Value* ast, Env* env) {
    const auto ans = new Vector();
    for (const auto& value : *ast->asVector()) {
        const auto [res, err] = Eval(value, env);
        if (err) return {res, true};
        ans->push(res);
    }
    return {ans, false};
}

std::tuple<Value*, bool> EvalAstHashMap(Value* ast, Env* env) {
    const auto ans = new HashMap();
    for (const auto& [key, value] : *ast->asHashMap()) {
        const auto [res, err] = Eval(value, env);
        if (err) return {res, true};
        ans->insert(key, res);
    }
    return {ans, false};
}

// std::tuple<Value*, bool> EvalAstFn(Value* ast, Env* env) {
//     std::cout << __PRETTY_FUNCTION__ << ": start" << std::endl;
//     const auto [res, err] = ast->asFn()->get()();
//     return {};
// }

std::tuple<Value*, bool> EvalAst(Value* ast, Env* env) {
    switch (ast->type()) {
        case Value::Type::Symbol:  return EvalAstSymbol(ast, env);
        case Value::Type::List:    return EvalAstList(ast, env);
        case Value::Type::Vector:  return EvalAstVector(ast, env);
        case Value::Type::HashMap: return EvalAstHashMap(ast, env);
        default:                   return {ast, false};
    }
    return {ast, false};
}

void Print(const std::string& input) {
    std::cout << input << std::endl;
}

bool REP(const std::string& input, Env* env) {
    if (const auto [ast, err] = Read(input); err) {
        return false;
    } else if (const auto [result, err] = Eval(ast, env); err) {
        if (result && result->type() == Value::Type::Exception) {
            std::cerr << result->inspect(false) << std::endl;
            return false;
        } else {
            std::cerr << "Unexpected eval error happened" << std::endl;
            return false;
        }
    } else {
        Print(PrintAst(result, true));
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
    const auto ns = BuildCoreNamespace();
    Env env(nullptr);
    for (const auto& [k, v] : ns) {
        env.set(new Symbol(k), new Fn(v));
    }

    for (;;) {
        const auto [line, ok] = GetLine();
        if (ok) {
            break;
        }
        REP(line, &env);
    }

    return 0;
}
