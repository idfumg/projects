uname_S := $(shell sh -c 'uname -s 2>/dev/null || echo not')

CXX=clang++
CC=clang++
LOCAL_CXX=clang++
LOCAL_CC=clang++

LD=$(CXX)
AR=ar

DEBUG=-ggdb
INCPATHS=
LIBPATHS=
CXXFLAGS=-O3 -Wall -Wextra -Werror -pedantic -pedantic-errors -std=c++20 $(DEBUG) $(INCPATHS)
LDFLAGS=-O3 $(DEBUG) $(LIBPATHS) -L.

SRCS=main.cpp main2.cpp
OBJS=$(SRCS:%.cpp=%.o)
BINS=$(OBJS:%.o=%)

OTHER_OBJS=
OTHER_HDRS=

all: $(BINS)

$(OBJS) $(OTHER_OBJS): %.o: %.cpp $(OTHER_HDRS)

$(patsubst %.o,%,$(filter step%,$(OBJS))): $(OTHER_OBJS)
$(BINS): %: %.o

clean:
	rm -rf *.o $(OBJS) $(BINS) $(OTHER_OBJS)
