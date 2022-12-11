#pragma once

#include <boost/beast/core.hpp>
#include <boost/beast/ssl.hpp>
#include <boost/beast/websocket.hpp>
#include <boost/beast/websocket/ssl.hpp>
#include <boost/asio/connect.hpp>
#include <boost/asio/ip/tcp.hpp>
#include <boost/asio/ssl/stream.hpp>

#include <string>
#include <optional>

namespace net = boost::asio;            // from <boost/asio.hpp>
namespace ssl = boost::asio::ssl;       // from <boost/asio/ssl.hpp>
namespace beast = boost::beast;         // from <boost/beast.hpp>
namespace websocket = beast::websocket; // from <boost/beast/websocket.hpp>
using tcp = boost::asio::ip::tcp;       // from <boost/asio/ip/tcp.hpp>

namespace ws {

struct Error : public std::string {
    template<class T>
    Error(T&& arg) : std::string(std::forward<T>(arg)) {

    }
};

class Response {
public:
    std::string text;
    std::optional<Error> error;

public:
    Response(const std::string& text, const std::optional<Error>& error);
};

class Client final {
private:
    // The io_context is required for all I/O
    net::io_context ioc{};
    tcp::resolver resolver{ioc};

    // The SSL context is required, and holds certificates
    // This holds the root certificate used for verification
    ssl::context ctx{ssl::context::tlsv12_client};

    websocket::stream<beast::ssl_stream<tcp::socket>> ws{ioc, ctx};

    std::string host;
    std::string port;

public:
    Client(const std::string& host, const std::string& port);
    ~Client();
    std::optional<Error> connect(const std::string& path = "/");
    Response sendRequest(const std::string& msg);
    Response waitResponse();
};

} // namespace ws