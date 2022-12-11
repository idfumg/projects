#include <iostream>
#include <optional>
#include <foo.hpp>
#include "websocketpp/server.hpp"
#include "websocketpp/config/asio_no_tls.hpp"

using namespace std;
using websocketpp::connection_hdl;

class WebsocketServer {
public:
    static bool init();
    static void run();
    static void stop();
 
    static bool sendClose(std::string id);
    static bool sendData(std::string id, std::string data);
         
private:
    static bool getWebsocket(const std::string &id, websocketpp::connection_hdl &hdl);
     
    inline static websocketpp::server<websocketpp::config::asio> server;
    inline static pthread_rwlock_t websocketsLock = PTHREAD_RWLOCK_INITIALIZER;
    inline static std::map<std::string, websocketpp::connection_hdl> websockets;
    // static LogStream ls;
    // inline static std::ostream os;
     
    // callbacks
    static bool on_validate(websocketpp::connection_hdl hdl);
    static void on_fail(websocketpp::connection_hdl hdl);
    static void on_close(websocketpp::connection_hdl hdl);
};


int main() {
    // static initialisations
    // ostream WebsocketServer::os(&ls);
        
    // namespace merging
    

    return 0;
}