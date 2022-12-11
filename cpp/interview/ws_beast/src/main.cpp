#include "Client.hpp"
#include <iostream>
#include <nlohmann/json.hpp>

using json = nlohmann::json;

int main(int argc, char** argv) {
    static const auto host = "ws.binaryws.com";
    static const auto port = "443";
    static const auto path = "/websockets/v3?app_id=1089";
    static const auto token = "q38JqJXrS5Ax84F";

    try {
        ws::Client client{host, port};

        const auto error = client.connect(path);
        if (error) {
            std::cout << error.value() << std::endl;
            return EXIT_FAILURE;
        }

        const auto msg = json{{"authorize", token}}.dump();
        std::cout << "Authorization request message: " << msg << std::endl;
        const auto response = client.sendRequest(msg);
        if (response.error) {
            std::cout << "Authorization error: " << response.error.value() << std::endl;
            return EXIT_FAILURE;
        }
        
        while (true) {
            const auto response = client.waitResponse();
            if (response.error) {
                std::cout << "Wait response error: " << response.error.value() << std::endl;
                return EXIT_FAILURE;
            }

            const auto root = json::parse(response.text);
            if (not root.contains("msg_type")) {
                std::cout << "Error! Wrong response: " << root.dump(2) << std::endl;
                return EXIT_FAILURE;
            }

            if (root["msg_type"] == "authorize") {
                const auto msg = json::parse(R"({
                    "buy": 1,
                    "subscribe": 1,
                    "price": 10,
                    "parameters": { 
                        "amount": 10, 
                        "basis": "stake", 
                        "contract_type": "CALL",
                        "currency": "USD", 
                        "duration": 1, 
                        "duration_unit": "m", 
                        "symbol": "R_10"
                    }})");

                const auto response = client.sendRequest(msg.dump());
                if (response.error) {
                    std::cout << "Send request error: " << response.error.value() << std::endl;
                    return EXIT_FAILURE;
                }
            }
            else if (root["msg_type"] == "buy") {
                if (root["buy"].contains("contract_id")) {
                    std::cout << "Contract Id: " << root["buy"]["contract_id"] << std::endl;
                }
                if (root["buy"].contains("longcode")) {
                    std::cout << "Details: " << root["buy"]["longcode"] << std::endl;
                }
            }
            else if (root["msg_type"] == "proposal_open_contract") {
                const bool isSold = root["proposal_open_contract"].contains("is_sold")
                    ? root["proposal_open_contract"]["is_sold"].get<int>() == 1
                    : false;

                if (isSold) {
                    const std::string status = root["proposal_open_contract"].contains("status")
                        ? root["proposal_open_contract"]["status"]
                        : "";
                    const double profit = root["proposal_open_contract"].contains("profit")
                        ? root["proposal_open_contract"]["profit"].get<double>()
                        : 0.0;

                    std::cout << "Contract: " << status << std::endl;
                    std::cout << "Profit: " << profit << std::endl;

                    break;
                }
                else {
                    const double currentSpot = root["proposal_open_contract"].contains("current_spot")
                        ? root["proposal_open_contract"]["current_spot"].get<double>()
                        : 0.0;
                    const double entrySpot = root["proposal_open_contract"].contains("entry_tick")
                        ? root["proposal_open_contract"]["entry_tick"].get<double>()
                        : 0.0;

                    std::cout << "Entry spot: " << entrySpot << std::endl;
                    std::cout << "Current spot: " << currentSpot << std::endl;
                    std::cout << "Difference: " << (currentSpot - entrySpot) << std::endl;
                }
            }
        }
    }
    catch (const std::exception& e) {
        std::cerr << "Unexpected error: " << e.what() << std::endl;
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}