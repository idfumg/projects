#include <iostream>
#include <utility>

template<class T>
class unique_ptr {
private:
    T* ptr = {};

public:
    ~unique_ptr() { cleanup(); }
    unique_ptr(T* ptr) : ptr{ptr} {}
    unique_ptr(const unique_ptr& rhs) = delete;
    unique_ptr& operator=(const unique_ptr& rhs) = delete;
    unique_ptr(unique_ptr&& rhs) { 
        ptr = std::exchange(rhs.ptr, {}); 
    }
    unique_ptr& operator=(unique_ptr&& rhs) {
        if (this != &rhs) {
            cleanup();
            ptr = std::exchange(rhs.ptr, {});
        }
        return *this;
    }

    T* operator->() const { return ptr; }
    T& operator*() const { return *ptr; }
    T* get() const { return ptr; }

private:
    void cleanup() {
        if (ptr) {
            delete ptr;
            ptr = {};
        }
    }
};

class Box {
public:
    int length = 5;
    int width = 2;
    int height = 3;
    ~Box() { std::cout << "~Box()" << std::endl; }
    Box() { std::cout << "Box()" << std::endl; }
    int square() const { return length * width * height; }
};

int main() {
    unique_ptr<Box> box1(new Box());
    std::cout << box1->square() << std::endl;
    auto box2 = std::move(box1);
    return 0;
}