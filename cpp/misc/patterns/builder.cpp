#include <iostream>
#include <vector>

using namespace std;

class Product {
public:
    vector<string> parts;
    void listParts() const {
        for_each(parts.begin(), parts.end(), [](const auto& part){cout << part << ' ';});
        cout << endl;
    }
};

class Builder {
public:
    virtual ~Builder(){}
    virtual void buildPartA() const = 0;
    virtual void buildPartB() const = 0;
    virtual void buildPartC() const = 0;
    virtual Product* getProduct() const = 0;
};

class ConcreteBuilder1 : public Builder {
private:
    Product* product = {};
public:
    ConcreteBuilder1() { product = new Product(); }
    virtual void buildPartA() const override { product->parts.push_back("a"); }
    virtual void buildPartB() const override { product->parts.push_back("b"); }
    virtual void buildPartC() const override { product->parts.push_back("c"); }
    virtual Product* getProduct() const override { return product; }
};

Product* buildMinimalViableProduct(Builder* builder) {
    builder->buildPartA();
    return builder->getProduct();
}

Product* buildFullFeaturedViableProduct(Builder* builder) {
    builder->buildPartA();
    builder->buildPartB();
    builder->buildPartC();
    builder->buildPartA();
    return builder->getProduct();
}

int main() {
    Product* product1 = buildMinimalViableProduct(new ConcreteBuilder1());
    product1->listParts();
    Product* product2 = buildFullFeaturedViableProduct(new ConcreteBuilder1());
    product2->listParts();
    return 0;
}