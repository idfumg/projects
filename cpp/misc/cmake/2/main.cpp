#include <iostream>
#include <optional>
#include <foo.hpp>

int main() {
    std::optional op = 1;
    std::cout << "Hello, world!" << std::endl;
    std::cout << *op << std::endl;
    Foo().print();
    return 0;
}