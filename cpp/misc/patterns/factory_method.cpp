#include <iostream>

using namespace std;

class Product {
public:
    virtual ~Product() {}
    virtual string msg() const = 0;
};

class ConcreteProduct1 : public Product {
public:
    virtual string msg() const override { return "ConcreteProduct1"; }
};

class ConcreteProduct2 : public Product {
public:
    virtual string msg() const override { return "ConcreteProduct2"; }
};

class Creator {
public:
    virtual ~Creator() {}
    virtual Product* factoryMethod() const = 0;
};

class ConcreteCreator1 : public Creator {
public:
    virtual Product* factoryMethod() const override { return new ConcreteProduct1(); }
};

class ConcreteCreator2 : public Creator {
public:
    virtual Product* factoryMethod() const override { return new ConcreteProduct2(); }
};

void client_code(Creator* creator) {
    cout << creator->factoryMethod()->msg() << endl;
}

int main() {
    client_code(new ConcreteCreator1());
    client_code(new ConcreteCreator2());
    return 0;
}