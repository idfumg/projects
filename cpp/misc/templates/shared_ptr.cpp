#include <iostream>

template<class T>
class shared_ptr {
private:
    T* ptr = {};
    int* cnt = new int(0);

public:
    ~shared_ptr() {
        cleanup();
    }
    shared_ptr(T* ptr) : ptr(ptr), cnt(new int(1)) {

    }
    shared_ptr(const shared_ptr& rhs) {
        cnt = rhs.cnt;
        if (ptr = rhs.ptr; ptr) ++(*cnt);
    }
    shared_ptr(shared_ptr&& rhs) {
        ptr = std::exchange(rhs.ptr, {});
        cnt = std::exchange(rhs.cnt, {});
    }
    shared_ptr& operator=(const shared_ptr& rhs) {
        if (this != &rhs) {
            cleanup();
            cnt = rhs.cnt;
            if (ptr = rhs.ptr; ptr) ++(*cnt);
        }
        return *this;
    }
    shared_ptr& operator=(shared_ptr&& rhs) {
        if (this != &rhs) {
            cleanup();
            ptr = std::exchange(rhs.ptr, {});
            cnt = std::exchange(rhs.cnt, {});
        }
        return *this;
    }
    T* operator->() const { return ptr; }
    T& operator*() const { return *ptr; }
    int get_count() const { return *cnt; }
    T* get() const { return ptr; }

private:
    void cleanup() {
        if (cnt and --(*cnt) == 0) {
            if (ptr) {
                delete ptr;
                ptr = {};
            }
            if (cnt) {
                delete cnt;
                cnt = {};
            }
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
    shared_ptr<Box> box1(new Box);
    std::cout << box1.get_count() << std::endl;
    std::cout << box1->square() << std::endl;
    auto box2 = box1;
    std::cout << box2.get_count() << std::endl;
    box2 = shared_ptr<Box>(new Box);
    std::cout << box2.get_count() << std::endl;
    return 0;
}