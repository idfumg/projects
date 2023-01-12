#include <iostream>

// Implementation with the primitive operations
class Device {
public:
    virtual ~Device() {}
    void enable() {}
    void disable() {}
    void setVolume(int value) {}
};

class Tv : public Device {};
class Radio : public Device {};

// Abstraction with the complex operations on the primitive operations from the implementation
class RemoteControl {
protected:
    Device* device = {};
public:
    virtual ~RemoteControl() {}
    RemoteControl(Device* device) : device(device) {}
    void volumeDown() {}
    void volumeUp() {}
};

class AdvancedRemoteControl : public RemoteControl {
public:
    AdvancedRemoteControl(Device* device) : RemoteControl(device) {}
    void mute() { device->setVolume(0); }
};

int main() {
    RemoteControl* remote = new RemoteControl(new Tv());
    remote->volumeDown();
    AdvancedRemoteControl* advancedRemote = new AdvancedRemoteControl(new Radio());
    advancedRemote->mute();
    return 0;
}