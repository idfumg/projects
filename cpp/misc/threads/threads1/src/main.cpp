#include <iostream>
#include <chrono>
#include <thread>

using namespace std;
using u64 = unsigned long long;

u64 even = 0;
u64 odd = 0;

void sumEven(u64 a, u64 b) {
    for (; a <= b; ++a) {
        if (a % 2 == 0) {
            even += a;
        }
    }
}

void sumOdd(u64 a, u64 b) {
    for (; a <= b; ++a) {
        if (a & 1) {
            odd += a;
        }
    }
}

int main(int argc, char** argv) {
    const u64 start = 0;
    const u64 stop = 19'000'00;
    const auto startTime = chrono::high_resolution_clock::now();

    std::thread t1(sumEven, start, stop);
    std::thread t2(sumOdd, start, stop);

    t1.join();
    t2.join();

    const auto stopTime = chrono::high_resolution_clock::now();
    const auto duration = chrono::duration_cast<chrono::seconds>(stopTime - startTime);
    cout << "even: " << even << ", odd = " << odd << ", (" << duration.count() << " seconds)" << endl;

    return EXIT_SUCCESS;
}