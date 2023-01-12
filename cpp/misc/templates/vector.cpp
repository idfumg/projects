#include <initializer_list>
#include <iostream>
#include <utility>

template<class T>
class vector {
private:
    T* ptr = {};
    int size = 0;

private:
    void cleanup() {
        if (ptr) {
            delete [] ptr;
            ptr = {};
            size = 0;
        }
    }

public:
    vector(const std::initializer_list<T>& list) {
        size = list.size();
        ptr = new T[size];
        std::copy(list.begin(), list.end(), ptr);
    }
    vector(const vector& rhs) {
        size = rhs.size;
        ptr = new T[size];
        std::copy(rhs.ptr, rhs.ptr + size, ptr);
    }
    vector(vector&& rhs) {
        cleanup();
        ptr = std::exchange(rhs.ptr, {});
        size = std::exchange(rhs.size, 0);
    }
    vector& operator=(const vector& rhs) {
        if (this != &this) {
            cleanup();
            size = rhs.size;
            ptr = new T[size];
            std::copy(rhs.ptr, rhs.ptr + size, ptr);
        }
    }
    vector& operator=(vector&& rhs) {
        if (this != &this) {
            cleanup();
            ptr = std::exchange(rhs.ptr, {});
            size = std::exchange(rhs.size, 0);
        }
    }
    T* begin() const {
        return ptr;
    }
    T* end() const {
        return ptr + size;
    }
};

int main() {
    const vector<int> v = {1,2,3,4,5,6,7,8,9};
    for (const int num : v) {
        std::cout << num << ' ';
    }
    std::cout << std::endl;
    return 0;
}