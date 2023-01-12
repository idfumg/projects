#include <iostream>

using namespace std;

class Target {
public:
    virtual ~Target(){}
    virtual string request() const { return "Target's default behavior"; }
};

class Adaptee {
public:
    string specificRequest() const { return ".eetpadA eht fo roivaheb laicepS"; }
};

class Adapter : public Target {
private:
    Adaptee adaptee;
public:
    virtual string request() const override {
        string value = adaptee.specificRequest();
        reverse(value.begin(), value.end());
        return value;
    }
};

void client_code(const Target* target) {
    cout << target->request() << endl;
}

int main() {
    client_code(new Target());
    client_code(new Adapter());
    return 0;
}