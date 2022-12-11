#include <cstdint>
#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <random>

#define UNIMPLEMENTED \
    do { \
        fprintf(stderr, "%s:%d:%s: UNIMPLEMENTED\n", __FILE__, __LINE__, __PRETTY_FUNCTION__); \
        exit(1); \
    } while (0);

#define UNUSED(x) \
    (void)(x);

static const auto NODE_POOL_CAP = 1024;
static const std::string DIGITS = "123456789";

struct Node {
    std::int32_t value;
    Node* left;
    Node* right;
};

struct NodePool {
    size_t size{};
    Node nodes[NODE_POOL_CAP];
};

static NodePool globalNodePool = {0, {}};

static Node* NodePoolAlloc(NodePool* pool) {
    assert(pool->size < NODE_POOL_CAP);
    Node* result = &pool->nodes[pool->size++];
    memset(result, 0, sizeof(Node));
    return result;
}

static void NodeSetValue(Node* root, const std::int32_t value) {
    root->value = value;
}

static Node* NodePoolAllocWithValue(NodePool* pool, const std::int32_t value) {
    Node* root = NodePoolAlloc(pool);
    NodeSetValue(root, value);
    return root;
}

static void PrintValueWithIndent(FILE* stream, const std::int32_t value, const std::size_t level) {
    for (std::size_t i = 0; i < level; ++i) {
        fputs("  ", stream);
    }
    std::cout << value;
    fputc('\n', stream);
}

static void PrintTree(FILE* stream, Node* node, const std::size_t level) {
    if (!node) {
        return;
    }

    PrintValueWithIndent(stream, node->value, level);
    PrintTree(stream, node->left, level + 1);
    PrintTree(stream, node->right, level + 1);
}

static std::int32_t GetRandomInteger(const std::size_t n) {
    std::random_device dev;
    std::mt19937 rng(dev());
    std::uniform_int_distribution<std::mt19937::result_type> dist(0, DIGITS.size());

    std::int32_t ans;
    for (std::size_t i = 1; i < n; ++i) {
        ans = ans * 10 + DIGITS[dist(rng)];
    }
    return ans;
}

static Node* RandomTree(NodePool* pool, const std::size_t level) {
    if (level == 0) return nullptr;
    Node* node = NodePoolAllocWithValue(pool, GetRandomInteger(6));
    node->left = RandomTree(pool, level - 1);
    node->right = RandomTree(pool, level - 1);
    return node;
}

// static void SaveNodePoolToFile(NodePool* pool, const char* filename) {
//     FILE* stream = fopen(filename, "wb");
//     if (stream == NULL) {
//         fprintf(stream, "Error! Could not write to file %s: %s\n", filename, strerror(errno));
//         exit(1);
//     }

//     const std::size_t n = fwrite(pool->nodes, sizeof(Node), pool->size, stream);
//     assert(n == pool->size);
//     if (ferror(stream)) {
//         fprintf(stream, "Error! Could not write to file: %s: %s\n", filename, strerror(errno));
//         exit(1);
//     }

//     fclose(stream);
// }

// void LoadNodePoolFromFile(NodePool* pool, const char* filename) {
//     FILE* stream = fopen(filename, "rb");
//     if (stream == NULL) {
//         fprintf(stream, "Error! Could not write to file %s: %s\n", filename, strerror(errno));
//         exit(1);        
//     }

//     fseek(stream, 0, SEEK_END);
//     const auto m = ftell(stream);
//     fseek(stream, 0, SEEK_SET);
//     assert(m % sizeof(Node) == 0);
//     pool->size = m / sizeof(Node);
//     assert(pool->size < NODE_POOL_CAP);

//     const std::size_t n = fread(pool->nodes, sizeof(Node), pool->size, stream);
//     assert(n == pool->size);
//     if (ferror(stream)) {
//         fprintf(stream, "Error! Could not write to file: %s: %s\n", filename, strerror(errno));
//         exit(1);
//     }

//     fclose(stream);
// }

int main() {
    Node* root = RandomTree(&globalNodePool, 6);
    PrintTree(stdout, root, 0);
    return 0;
}
