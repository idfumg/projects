#pragma once

#include <string>

class Symbol;
class List;
class Vector;
class HashMap;
class Integer;
class Fn;
class True;
class False;
class Nil;
class String;
class Keyword;

class Value {
public:
    enum class Type {
        Symbol,
        List,
        Vector,
        HashMap,
        Integer,
        Fn,
        Exception,
        True,
        False,
        Nil,
        String,
        Keyword,
    };

public:
    virtual std::string inspect(const bool printReadably) const = 0;
    virtual Type type() const = 0;

    virtual bool isSymbol() const;
    virtual bool isTruth() const;
    virtual bool isTrue() const;
    virtual bool isFalse() const;
    virtual bool isList() const;
    virtual bool isLikeList() const;
    virtual bool isVector() const;
    virtual bool isInteger() const;
    virtual bool isNil() const;
    virtual bool isHashMap() const;
    virtual bool isString() const;
    virtual bool isKeyword() const;

    Symbol* asSymbol() const;
    List* asList() const;
    Vector* asVector() const;
    HashMap* asHashMap() const;
    Integer* asInteger() const;
    Fn* asFn() const;
    True* asTrue() const;
    False* asFalse() const;
    Nil* asNil() const;
    String* asString() const;
    Keyword* asKeyword() const;
};

bool IsEqual(const Value* a, const Value* b);
bool IsEmpty(const Value* v);