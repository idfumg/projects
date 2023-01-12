#include <iostream>

using namespace std;

void print() {
    cout << endl;
}

template<class Arg, class ... Args>
void print(Arg&& arg, Args&& ... args) {
    cout << arg << ' ';
    print(forward<Args>(args)...);
}

int main() {
    print(1, 2, 3, "hello!", 'Q');
    return 0;
}