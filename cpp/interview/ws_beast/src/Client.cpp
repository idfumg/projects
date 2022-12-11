#include "Client.hpp"

#include "root_certificates.hpp"

namespace http = beast::http;           // from <boost/beast/http.hpp>

namespace ws {

Response::Response(const std::string& text, const std::optional<Error>& error) :
    text{text}, error{error}
{

}

Client::Client(const std::string& host, const std::string& port) : 
    ioc{}, 
    resolver{ioc}, 
    ctx{ssl::context::tlsv12_client},
    ws{ioc, ctx},
    host{host},
    port{port} 
{
    load_root_certificates(ctx);
}

Client::~Client() {
    if (ws.is_open()) {
        boost::beast::error_code ec;
        ws.close(websocket::close_code::normal, ec);
        // if the remote host has already closed his side of connection => do nothing with that type of errors
    }
}

std::optional<Error> Client::connect(const std::string& path) {
    boost::beast::error_code ec;

    const auto results = resolver.resolve(host, port, ec);
    if (ec) {
        return Error{ec.message()};
    }

    const auto ep = net::connect(get_lowest_layer(ws), results, ec);
    if (ec) {
        return Error{ec.message()};
    }

    if (! SSL_set_tlsext_host_name(ws.next_layer().native_handle(), host.c_str())) {
        throw beast::system_error(
            beast::error_code(
                static_cast<int>(::ERR_get_error()),
                net::error::get_ssl_category()),
            "Failed to set SNI Hostname");
    }

    const std::string url = host + ':' + std::to_string(ep.port());

    ws.next_layer().handshake(ssl::stream_base::client, ec);
    if (ec) {
        return Error{ec.message()};
    }

    ws.set_option(websocket::stream_base::decorator(
        [](websocket::request_type& req)
        {
            req.set(http::field::user_agent,
                std::string(BOOST_BEAST_VERSION_STRING) +
                    " websocket-client-coro");
        }));

    ws.handshake(url, path, ec);
    if (ec) {
        return Error{ec.message()};
    }

    return {};
}

Response Client::sendRequest(const std::string& msg) {
    boost::beast::error_code ec;

    ws.write(net::buffer(std::string(msg)), ec);
    if (ec) {
        return {"", ec.message()};
    }

    return {"", {}};
}

Response Client::waitResponse() {
    boost::beast::error_code ec;

    beast::flat_buffer buffer;
    ws.read(buffer, ec);
    if (ec) {
        return {"", ec.message()};
    }

    return {
        std::string(
            static_cast<const char*>(buffer.data().data()),
            static_cast<const char*>(buffer.data().data()) + buffer.data().size()),
        {}
    };
}

} // namespace ws