#include "types.hpp"
#include "value.hpp"

bool Value::isSymbol() const { 
    return false; 
}

bool Value::isTruth() const { 
    return true;
}

bool Value::isTrue() const { 
    return false;
}

bool Value::isFalse() const { 
    return false;
}

bool Value::isLikeList() const { 
    return false;
}

bool Value::isList() const { 
    return false;
}

bool Value::isVector() const { 
    return false;
}

bool Value::isInteger() const { 
    return false;
}

bool Value::isNil() const { 
    return false;
}

bool Value::isHashMap() const { 
    return false;
}

bool Value::isString() const { 
    return false;
}

bool Value::isKeyword() const { 
    return false;
}

Symbol* Value::asSymbol() const {
    assert(type() == Type::Symbol);
    return static_cast<Symbol*>(const_cast<Value*>(this));
}

List* Value::asList() const {
    assert(type() == Type::List || type() == Type::Vector);
    return static_cast<List*>(const_cast<Value*>(this));
}

Vector* Value::asVector() const {
    assert(type() == Type::Vector);
    return static_cast<Vector*>(const_cast<Value*>(this));
}

HashMap* Value::asHashMap() const {
    assert(type() == Type::HashMap);
    return static_cast<HashMap*>(const_cast<Value*>(this));
}

Integer* Value::asInteger() const {
    assert(type() == Type::Integer);
    return static_cast<Integer*>(const_cast<Value*>(this));
}

Fn* Value::asFn() const {
    assert(type() == Type::Fn);
    return static_cast<Fn*>(const_cast<Value*>(this));
}

True* Value::asTrue() const {
    assert(type() == Type::True);
    return static_cast<True*>(const_cast<Value*>(this));
}

False* Value::asFalse() const {
    assert(type() == Type::False);
    return static_cast<False*>(const_cast<Value*>(this));
}

Nil* Value::asNil() const {
    assert(type() == Type::Nil);
    return static_cast<Nil*>(const_cast<Value*>(this));
}

String* Value::asString() const {
    assert(type() == Type::String);
    return static_cast<String*>(const_cast<Value*>(this));
}

Keyword* Value::asKeyword() const {
    assert(type() == Type::Keyword);
    return static_cast<Keyword*>(const_cast<Value*>(this));
}

bool IsEqual(const Value* a, const Value* b) {
    if (a->isLikeList() && b->isLikeList() && *a->asList()    == *b->asList())    return true;
    if (a->isList()     && b->isList()     && *a->asList()    == *b->asList())    return true;
    if (a->isVector()   && b->isVector()   && *a->asVector()  == *b->asVector())  return true;
    if (a->isInteger()  && b->isInteger()  && *a->asInteger() == *b->asInteger()) return true;
    if (a->isNil()      && b->isNil()      && *a->asNil()     == *b->asNil())     return true;
    if (a->isTrue()     && b->isTrue()     && *a->asTrue()    == *b->asTrue())    return true;
    if (a->isFalse()    && b->isFalse()    && *a->asFalse()   == *b->asFalse())   return true;
    if (a->isHashMap()  && b->isHashMap()  && *a->asHashMap() == *b->asHashMap()) return true;
    if (a->isSymbol()   && b->isSymbol()   && *a->asSymbol()  == *b->asSymbol())  return true;
    if (a->isString()   && b->isString()   && *a->asString()  == *b->asString())  return true;
    if (a->isKeyword()  && b->isKeyword()  && *a->asKeyword() == *b->asKeyword()) return true;
    return false;
}

bool IsEmpty(const Value* v) {
    if (!v) return true;
    if (v->isList() && v->asList()->size() == 0) return true;
    if (v->isVector() && v->asVector()->size() == 0) return true;
    if (v->isHashMap() && v->asHashMap()->size() == 0) return true;
    if (v->isString() && v->asString()->size() == 0) return true;
    return false;
}