#include <iostream>

class string {
private:
    char* ptr = {};
    int size = 0;

public:
    ~string() {
        cleanup();
    }
    string() : ptr(nullptr), size() {

    }
    string(const char* buffer) {
        size = strlen(buffer);
        ptr = new char[size + 1];
        std::copy(buffer, buffer + size, ptr);
        ptr[size] = '\0';
    }
    string(const string& rhs) {
        size = rhs.size;
        ptr = new char[size + 1];
        std::copy(rhs.ptr, rhs.ptr + size, ptr);
        ptr[size] = '\0';
    }
    string& operator=(const string& rhs) {
        if (this != &rhs) {
            cleanup();
            size = rhs.size;
            ptr = new char[size + 1];
            std::copy(rhs.ptr, rhs.ptr + size, ptr);
            ptr[size] = '\0';
        }
        return *this;
    }
    string(string&& rhs) {
        size = std::exchange(rhs.size, 0);
        ptr = std::exchange(rhs.ptr, {});
    }
    string& operator=(string&& rhs) {
        if (this != &rhs) {
            cleanup();
            size = std::exchange(rhs.size, 0);
            ptr = std::exchange(rhs.ptr, {});
        }
        return *this;
    }
    string operator+(const string& s) const {
        string ans;
        ans.size = size + s.size;
        ans.ptr = new char[ans.size + 1];
        std::copy(ptr, ptr + size, ans.ptr);
        std::copy(s.ptr, s.ptr + s.size, ans.ptr + size);
        ans.ptr[ans.size] = '\0';
        return ans;
    }
    const char* c_str() const { return ptr; }

private:
    void cleanup() {
        if (ptr) {
            delete[] ptr;
            ptr = {};
            size = 0;
        }
    }
};

std::ostream& operator<<(std::ostream& os, const string& s) {
    return os << s.c_str();
}

int main() {
    const string s = "asdasd";
    const string t = "12123";
    const string p = s + t;
    std::cout << s << ' ' << t << ' ' << p << std::endl;
    return 0;
}