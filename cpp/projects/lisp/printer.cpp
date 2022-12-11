#include "printer.hpp"

std::string PrintAst(Value* ast, const bool printReadably) {
    assert(ast);
    return ast->inspect(printReadably);
}