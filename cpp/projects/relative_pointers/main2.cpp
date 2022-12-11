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

#define abs2rel32(rloc, aptr) \
    rloc = ((int32_t)((uint8_t*)(aptr) - (uint8_t*)&(rloc)))^0x80000000

#define rel2abs32(Type, rloc) \
    (Type*)((uint8_t*)&rloc + (rloc^0x80000000))

static const auto NODE_POOL_CAP = 1024;
static const auto DIGITS = "123456789";
static const auto DIGITS_LEN = strlen(DIGITS);
static const auto FILENAME = "data.bin";

struct Node {
    std::int32_t value;
    std::int32_t left;
    std::int32_t right;
};

struct NodePool {
    size_t size{};
    Node nodes[NODE_POOL_CAP];
};

static NodePool globalNodePool = {0, {}};

[[maybe_unused]] static Node* NodePoolAlloc(NodePool* pool) {
    assert(pool->size < NODE_POOL_CAP);
    Node* result = &pool->nodes[pool->size++];
    memset(result, 0, sizeof(Node));
    return result;
}

[[maybe_unused]] static Node* NodePoolAllocWithValue(NodePool* pool, const std::int32_t value) {
    Node* root = NodePoolAlloc(pool);
    root->value = value;
    return root;
}

[[maybe_unused]] static void PrintNode(FILE* stream, Node* node, const std::size_t level) {
    for (std::size_t i = 0; i < level; ++i) {
        fputs("  ", stream);
    }
    std::cout << node->value;
    fputc('\n', stream);
}

[[maybe_unused]] static std::int32_t GetRandomInteger(const std::size_t n) {
    std::random_device dev;
    std::mt19937 rng(dev());
    std::uniform_int_distribution<std::mt19937::result_type> dist(0, DIGITS_LEN);

    std::int32_t ans;
    for (std::size_t i = 1; i < n; ++i) {
        ans = ans * 10 + DIGITS[dist(rng)];
    }
    return ans;
}

[[maybe_unused]] static void PrintTree(FILE* stream, Node* node, const std::size_t level) {
    PrintNode(stream, node, level);
    if (node->left) PrintTree(stream, rel2abs32(Node, node->left), level + 1);
    if (node->right) PrintTree(stream, rel2abs32(Node, node->right), level + 1);
}

[[maybe_unused]] static Node* RandomTree(NodePool* pool, const std::size_t level) {
    Node* node = NodePoolAllocWithValue(pool, GetRandomInteger(6));
    if (level > 1) abs2rel32(node->left, RandomTree(pool, level - 1));
    if (level > 1) abs2rel32(node->right, RandomTree(pool, level - 1));
    return node;
}

[[maybe_unused]] static void SaveNodePoolToFile(NodePool* pool, const char* filename) {
    FILE* stream = fopen(filename, "wb");
    if (stream == NULL) {
        fprintf(stream, "Error! Could not write to file %s: %s\n", filename, strerror(errno));
        exit(1);
    }

    const std::size_t n = fwrite(pool->nodes, sizeof(Node), pool->size, stream);
    assert(n == pool->size);
    if (ferror(stream)) {
        fprintf(stream, "Error! Could not write to file: %s: %s\n", filename, strerror(errno));
        exit(1);
    }

    fclose(stream);
}

[[maybe_unused]] static std::int64_t GetFileLength(FILE* stream) {
    fseek(stream, 0, SEEK_END);
    const auto m = ftell(stream);
    fseek(stream, 0, SEEK_SET);
    return m;
}

[[maybe_unused]] static void LoadNodePoolFromFile(NodePool* pool, const char* filename) {
    FILE* stream = fopen(filename, "rb");
    if (stream == NULL) {
        fprintf(stream, "Error! Could not write to file %s: %s\n", filename, strerror(errno));
        exit(1);        
    }

    const auto len = GetFileLength(stream);
    assert(len % sizeof(Node) == 0);
    pool->size = len / sizeof(Node);
    assert(pool->size < NODE_POOL_CAP);

    const std::size_t n = fread(pool->nodes, sizeof(Node), pool->size, stream);
    assert(n == pool->size);
    if (ferror(stream)) {
        fprintf(stream, "Error! Could not write to file: %s: %s\n", filename, strerror(errno));
        exit(1);
    }

    fclose(stream);
}

int main1() {
    Node* root = RandomTree(&globalNodePool, 6);
    PrintTree(stdout, root, 0);
    SaveNodePoolToFile(&globalNodePool, FILENAME);
    return 0;
}

int main() {
    LoadNodePoolFromFile(&globalNodePool, FILENAME);
    PrintTree(stdout, &globalNodePool.nodes[0], 0);
    return 0;
}
