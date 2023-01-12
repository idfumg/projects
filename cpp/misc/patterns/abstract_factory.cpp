#include <iostream>

using namespace std;

class ProductA {
public:
    virtual ~ProductA() {}
    virtual string msg() const = 0;
};

class ConcreteProductA1 : public ProductA {
public:
    virtual string msg() const override { return "ConcreteProductA1"; }
};

class ConcreteProductA2 : public ProductA {
public:
    virtual string msg() const override { return "ConcreteProductA2"; }
};

class ProductB {
public:
    virtual ~ProductB() {}
    virtual string msg() const = 0;
};

class ConcreteProductB1 : public ProductB {
public:
    virtual string msg() const override { return "ConcreteProductB1"; }
};

class ConcreteProductB2 : public ProductB {
public:
    virtual string msg() const override { return "ConcreteProductB2"; }
};

class Factory {
public:
    virtual ~Factory() {}
    virtual ProductA* createProductA() const = 0;
    virtual ProductB* createProductB() const = 0;
};

class ConcreteFactory1 : public Factory {
public:
    ProductA* createProductA() const override { return new ConcreteProductA1(); }
    ProductB* createProductB() const override { return new ConcreteProductB1(); }
};

class ConcreteFactory2 : public Factory {
public:
    ProductA* createProductA() const override { return new ConcreteProductA2(); }
    ProductB* createProductB() const override { return new ConcreteProductB2(); }
};

void client_code(Factory* factory) {
    const ProductA* a = factory->createProductA();
    const ProductB* b = factory->createProductB();
    cout << a->msg() << ' ' << b->msg() << endl;
}

int main() {
    client_code(new ConcreteFactory1());
    client_code(new ConcreteFactory2());
    return 0;
}