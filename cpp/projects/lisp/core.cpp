#include "core.hpp"
#include "types.hpp"
#include "printer.hpp"
#include <iostream>

std::unordered_map<std::string, Func> BuildCoreNamespace() {
    std::unordered_map<std::string, Func> ns;
    ns["+"] = Add;
    ns["-"] = Sub;
    ns["*"] = Mul;
    ns["/"] = Div;
    ns["prn"] = Prn;
    ns["pr-str"] = PrStr;
    ns["str"] = Str;
    ns["list"] = ListFn;
    ns["vector"] = VectorFn;
    ns["hashmap"] = HashMapFn;
    ns["list?"] = ListQ;
    ns["vector?"] = VectorQ;
    ns["hashmap?"] = HashMapQ;
    ns["empty?"] = EmptyQ;
    ns["count"] = Count;
    ns["="] = Eq;
    ns["<"] = Lt;
    ns["<="] = Lte;
    ns[">"] = Gt;
    ns[">="] = Gte;
    ns["not"] = Not;
    ns["println"] = Println;
    return ns;
}

std::tuple<Value*, bool> Add(const std::size_t argc, Value*const* argv) {
    if (argc != 2) {
        return {new Exception("Error! Not enough params!"), true};
    }
    const auto a = argv[0];
    const auto b = argv[1];
    if (a->type() != Value::Type::Integer || b->type() != Value::Type::Integer) {
        return {new Exception("Error! Wrong params types!"), true};
    }
    return {new Integer(a->asInteger()->get() + b->asInteger()->get()), false};
}

std::tuple<Value*, bool> Sub(const std::size_t argc, Value*const* argv) {
    if (argc != 2) {
        return {new Exception("Error! Not enough params!"), true};
    }
    const auto a = argv[0];
    const auto b = argv[1];
    if (a->type() != Value::Type::Integer || b->type() != Value::Type::Integer) {
        return {new Exception("Error! Wrong params types!"), true};
    }
    return {new Integer(a->asInteger()->get() - b->asInteger()->get()), false};
}

std::tuple<Value*, bool> Mul(const std::size_t argc, Value*const* argv) {
    if (argc != 2) {
        return {new Exception("Error! Not enough params!"), true};
    }
    const auto a = argv[0];
    const auto b = argv[1];
    if (a->type() != Value::Type::Integer || b->type() != Value::Type::Integer) {
        return {new Exception("Error! Wrong params types!"), true};
    }
    return {new Integer(a->asInteger()->get() * b->asInteger()->get()), false};
}

std::tuple<Value*, bool> Div(const std::size_t argc, Value*const* argv) {
    if (argc != 2) {
        return {new Exception("Error! Not enough params!"), true};
    }
    const auto a = argv[0];
    const auto b = argv[1];
    if (a->type() != Value::Type::Integer || b->type() != Value::Type::Integer) {
        return {new Exception("Error! Wrong params types!"), true};
    }
    if (b->asInteger()->get() == 0) {
        return {new Exception("Error! Division by zero!"), true};
    }
    return {new Integer(a->asInteger()->get() / b->asInteger()->get()), false};
}

std::tuple<Value*, bool> Prn(const std::size_t argc, Value*const* argv) {
    if (argc == 0) {
        std::cout << std::endl;
        return {new Nil(), false};
    }
    std::string s;
    for (std::size_t i = 0; i < argc; ++i) {
        s += PrintAst(argv[i], true);
        if (i != argc - 1) s += ' ';
    }
    s += '\n';
    std::cout << s;
    return {new Nil(), false};
}

std::tuple<Value*, bool> PrStr(const std::size_t argc, Value*const* argv) {
    if (argc == 0) return {new String(""), false};
    std::string ans;
    for (std::size_t i = 0; i < argc; ++i) {
        ans += PrintAst(argv[i], true);
        if (i < argc - 1) ans += ' ';
    }
    return {new String(ans), false};
}

std::tuple<Value*, bool> Str(const std::size_t argc, Value*const* argv) {
    if (argc == 0) return {new String(""), false};
    std::string ans;
    for (std::size_t i = 0; i < argc; ++i) {
        ans += PrintAst(argv[i], false);
    }
    return {new String(ans), false};
}

std::tuple<Value*, bool> ListFn(const std::size_t argc, Value*const* argv) {
    const auto ans = new List();
    for (std::size_t i = 0; i < argc; ++i) {
        ans->push(argv[i]);
    }
    return {ans, false};
}

std::tuple<Value*, bool> VectorFn(const std::size_t argc, Value*const* argv) {
    const auto ans = new Vector();
    for (std::size_t i = 0; i < argc; ++i) {
        ans->push(argv[i]);
    }
    return {ans, false};
}

std::tuple<Value*, bool> HashMapFn(const std::size_t argc, Value*const* argv) {
    assert(argc % 2 == 0);
    const auto ans = new HashMap();
    for (std::size_t i = 0; i < argc; i += 2) {
        ans->insert(argv[i], argv[i + 1]);
    }
    return {ans, false};
}

std::tuple<Value*, bool> ListQ(const std::size_t argc, Value*const* argv) {
    assert(argc > 0);
    if (argv[0]->isList()) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> VectorQ(const std::size_t argc, Value*const* argv) {
    assert(argc > 0);
    if (argv[0]->isVector()) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> HashMapQ(const std::size_t argc, Value*const* argv) {
    assert(argc > 0);
    if (argv[0]->isHashMap()) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> EmptyQ(const std::size_t argc, Value*const* argv) {
    assert(argc > 0);
    if (IsEmpty(argv[0])) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> Count(const std::size_t argc, Value*const* argv) {
    assert(argc > 0);
    if (argv[0]->isList()) return {new Integer(argv[0]->asList()->size()), false};
    if (argv[0]->isVector()) return {new Integer(argv[0]->asVector()->size()), false};
    if (argv[0]->isHashMap()) return {new Integer(argv[0]->asHashMap()->size()), false};
    return {new Integer(0), false};
}

std::tuple<Value*, bool> Eq(const std::size_t argc, Value*const* argv) {
    assert(argc > 1);
    const auto a = argv[0];
    const auto b = argv[1];
    if (IsEqual(a, b)) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> Lt(const std::size_t argc, Value*const* argv) {
    assert(argc > 1);
    const auto a = argv[0];
    const auto b = argv[1];
    assert(a->isInteger());
    assert(b->isInteger());
    if (*a->asInteger() < *b->asInteger()) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> Lte(const std::size_t argc, Value*const* argv) {
    assert(argc > 1);
    const auto a = argv[0];
    const auto b = argv[1];
    assert(a->isInteger());
    assert(b->isInteger());
    if (*a->asInteger() <= *b->asInteger()) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> Gt(const std::size_t argc, Value*const* argv) {
    assert(argc > 1);
    const auto a = argv[0];
    const auto b = argv[1];
    assert(a->isInteger());
    assert(b->isInteger());
    if (*a->asInteger() > *b->asInteger()) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> Gte(const std::size_t argc, Value*const* argv) {
    assert(argc > 1);
    const auto a = argv[0];
    const auto b = argv[1];
    assert(a->isInteger());
    assert(b->isInteger());
    if (*a->asInteger() >= *b->asInteger()) return {new True(), false};
    return {new False(), false};
}

std::tuple<Value*, bool> Not(const std::size_t argc, Value*const* argv) {
    assert(argc > 0);
    const auto a = argv[0];
    if (a->isTruth()) return {new False(), false};
    return {new True(), false};
}

std::tuple<Value*, bool> Println(const std::size_t argc, Value*const* argv) {
    if (argc == 0) {
        std::cout << std::endl;
        return {new Nil(), false};
    }
    for (std::size_t i = 0; i < argc; ++i) {
        std::cout << PrintAst(argv[i], false);
        if (i != argc - 1) {
            std::cout << ' ';
        }
    }
    std::cout << std::endl;
    return {new Nil(), false};
}
