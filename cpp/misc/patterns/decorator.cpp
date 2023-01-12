#include <iostream>

using namespace std;

class DataSource {
public:
    virtual ~DataSource() {}
    virtual string readData() = 0;
    virtual void writeData(const string& data) = 0;
};

class FileDataSource : public DataSource {
public:
    FileDataSource(const string& filename) : DataSource() {}
    virtual string readData() override { return ""; }
    virtual void writeData(const std::string& data) override {  };
};

class DataSourceDecorator : public DataSource {
protected:
    DataSource* source = {};
public:
    virtual ~DataSourceDecorator() {}
    DataSourceDecorator(DataSource* source) : source(source) {}
    virtual string readData() override { return source->readData(); }
    virtual void writeData(const string& data) override { source->writeData(data); }
};

class EncryptionDecorator : public DataSourceDecorator {
public:
    EncryptionDecorator(DataSource* source) : DataSourceDecorator(source) {}
    virtual string readData() override {
        const auto data = source->readData();
        // decrypt data
        return data;
    }
    virtual void writeData(const string& data) override {
        // encrypt data
        source->writeData(data);
    }
};

class CompressionDecorator : public DataSourceDecorator {
public:
    CompressionDecorator(DataSource* source) : DataSourceDecorator(source) {}
    virtual string readData() override {
        const auto data = source->readData();
        // decompredd data
        return data;
    }
    virtual void writeData(const string& data) override {
        // compress data
        source->writeData(data);
    }
};

int main() {
    DataSource* source = new FileDataSource("/tmp/data");
    if (true /*isEncryptionEnabled*/) {
        source = new EncryptionDecorator(source);
    }
    if (true /*isCompressionEnabled*/) {
        source = new CompressionDecorator(source);
    }
    const auto data = source->readData();
    source->writeData(data);
    return 0;
}