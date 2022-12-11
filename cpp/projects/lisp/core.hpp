#pragma once

#include "fn.hpp"

#include <unordered_map>

std::unordered_map<std::string, Func> BuildCoreNamespace();

std::tuple<Value*, bool> Add(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Sub(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Mul(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Div(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Prn(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> PrStr(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Str(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> ListFn(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> VectorFn(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> HashMapFn(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> ListQ(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> VectorQ(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> HashMapQ(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> EmptyQ(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Count(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Eq(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Lt(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Lte(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Gt(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Gte(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Not(const std::size_t argc, Value*const* argv);
std::tuple<Value*, bool> Println(const std::size_t argc, Value*const* argv);