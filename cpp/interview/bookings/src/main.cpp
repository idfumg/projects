/*
1. add bookings
2. select bookings departing before a given time
3. select bookings visiting two airports sequentially. 
    (e.g. the search AMS→LHR gives you bookings with itineraries like HAM→AMS→LHR, AMS→LHR, AAL→AMS→LHR→JFK→SFO etc,
    but not LHR→AMS or AMS→CDG→LHR)
*/

#include <set>
#include <unordered_map>
#include <vector>
#include <iostream>

using PaxDepartureItinRaw = std::tuple<std::string, std::string, std::string>;
using Code = std::string;
using CodeList = std::vector<Code>;
using CodePair = std::pair<Code, Code>;
using Indexes = std::vector<std::size_t>;

static const PaxDepartureItinRaw raws[] = {
    {"Alice", "May-26 06:45 2020", "LHR->AMS"},
    {"Bruce", "Jun-04 11:04 2020", "GVA->AMS->LHR"},
    {"Cindy", "Jun-06 10:00 2020", "AAL->AMS->LHR->JFK->SFO"},
    {"Derek", "Jun-12 08:09 2020", "AMS->LHR"},
    {"Erica", "Jun-13 20:40 2020", "ATL->AMS->AAL"},
    {"Fred", "Jun-14 09:10 2020", "AMS->CDG->LHR"},
};

static time_t ParseTime(const std::string& departure) {
    struct tm result;
    if (strptime(departure.c_str(), "%b-%d %R %Y", &result) == NULL) {
        throw std::runtime_error("Wrong departure date: " + departure);
    }
    return timegm(&result);
}

static std::vector<std::string> ItinStringToCodesList(const std::string& itin) {
    CodeList codes;
    Code code;
    for (std::size_t i = 0; i < itin.size(); ++i) {
        if (i > 0 and itin[i - 1] == '-' and itin[i] == '>') {
            codes.push_back(code);
            code.clear();
        }
        else if (std::isalpha(itin[i])) {
            code += itin[i];
        }
    }
    codes.push_back(code);
    return codes;
}

class Booking final {
public:
    std::string pax;
    std::string departure;
    time_t date;
    std::string itin;
    CodeList codes;

public:
    bool operator<(const Booking& rhs) const noexcept {
        return date < rhs.date;
    }

public:
    static Booking parse(const PaxDepartureItinRaw& raw) {
        const auto& [pax, departure, itin] = raw;
        return {
            pax,
            departure,
            ParseTime(departure),
            itin,
            ItinStringToCodesList(itin),
        };
    }
};

std::ostream& operator<<(std::ostream& os, const Booking& booking) {
    return os << booking.pax << ' ' << booking.departure << ' ' << booking.itin;
}

struct PairHash { // https://stackoverflow.com/questions/32685540/why-cant-i-compile-an-unordered-map-with-a-pair-as-key
    template <class T1, class T2>
    std::size_t operator()(const std::pair<T1, T2>& p) const { // naive hash function
        return std::hash<T1>{}(p.first) ^ std::hash<T2>{}(p.second);
    }
};

using CodePairToIndexes = std::unordered_map<CodePair, std::vector<Booking>, PairHash>;

class Manager final {
private:
    std::set<Booking> bookings;
    CodePairToIndexes matched;

public:
    void addBooking(const PaxDepartureItinRaw& raw) {
        const auto booking = *bookings.insert(Booking::parse(raw)).first;
        for (std::size_t i = 1; i < booking.codes.size(); ++i) {
            matched[{booking.codes[i - 1], booking.codes[i]}].push_back(booking);
        }
    }

    std::vector<Booking> findBookingsByCodes(const std::string& itin) const { // O(1) - depends on a hash function
        const auto codes = ItinStringToCodesList(itin);
        return matched.at({codes[0], codes[1]});
    }

    std::set<Booking> findBookingsByDate(const std::string& date) const { // O(n) - worst case, O(logn) - otherwise
        const auto it = bookings.lower_bound(Booking{"", {}, ParseTime(date), {}, {}});
        return std::set<Booking>(bookings.begin(), it);
    }
};

int main(int argc, char** argv) {
    Manager manager;

    for (const auto& raw : raws) {
        manager.addBooking(raw);
    }

    for (const auto& booking : manager.findBookingsByCodes("AMS->LHR")) {
        std::cout << booking << std::endl;
    }
    std::cout << std::endl;

    for (const auto& booking : manager.findBookingsByCodes("ATL->AMS")) {
        std::cout << booking << std::endl;
    }
    std::cout << std::endl;

    for (const auto& booking : manager.findBookingsByDate("Jun-07 09:10 2020")) {
        std::cout << booking << std::endl;
    }
    std::cout << std::endl;

    return EXIT_SUCCESS;
}