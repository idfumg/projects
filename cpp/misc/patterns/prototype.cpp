#include <iostream>
#include <unordered_map>

using namespace std;

class Prototype {
protected:
    string name;
public:
    virtual ~Prototype(){}
    Prototype(const string& name) : name(name) {}
    const string& getName() const { return name; }
    virtual Prototype* clone() const = 0;
};

class ConcretePrototype1 : public Prototype {
    [[maybe_unused]] int value = 0;
public:
    ConcretePrototype1(const string& name, int value) : Prototype(name), value(value) {}
    virtual Prototype* clone() const override { return new ConcretePrototype1(*this); }
};

class ConcretePrototype2 : public Prototype {
    [[maybe_unused]] int value = 0;
public:
    ConcretePrototype2(const string& name, int value) : Prototype(name), value(value) {}
    virtual Prototype* clone() const override { return new ConcretePrototype2(*this); }
};

enum Type {
    Type1,
    Type2,
};

class PrototypeFactory {
private:
    unordered_map<Type, Prototype*, hash<int>> prototypes;
public:
    PrototypeFactory() {
        prototypes[Type::Type1] = new ConcretePrototype1("type1", 10);
        prototypes[Type::Type2] = new ConcretePrototype2("type2", 25);
    }
    Prototype* create(Type type) {
        return prototypes[type]->clone();
    }
};

int main() {
    PrototypeFactory* factory = new PrototypeFactory();
    Prototype* prototype = factory->create(Type::Type1);
    cout << prototype->getName() << endl;
    return 0;
}