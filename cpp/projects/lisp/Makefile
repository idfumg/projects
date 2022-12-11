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

SRCS=step0_repl.cpp step1_read_print.cpp step2_eval.cpp step3_env.cpp step4_if_fn_do.cpp
OBJS=$(SRCS:%.cpp=%.o)
BINS=$(OBJS:%.o=%)

OTHER_OBJS=readline.o   tokenizer.o   reader.o   printer.o   env.o   exception.o   fn.o   hashmap.o   integer.o   list.o   symbol.o   value.o   vector.o   true.o   false.o   nil.o   core.o   string.o   keyword.o
OTHER_HDRS=readline.hpp tokenizer.hpp reader.hpp printer.hpp env.hpp exception.cpp fn.hpp hashmap.hpp integer.hpp list.hpp symbol.hpp value.hpp vector.hpp true.hpp false.hpp nil.hpp core.hpp string.hpp keyword.hpp

all: $(BINS)

$(OBJS) $(OTHER_OBJS): %.o: %.cpp $(OTHER_HDRS)

$(patsubst %.o,%,$(filter step%,$(OBJS))): $(OTHER_OBJS)
$(BINS): %: %.o

clean:
	rm -rf *.o $(OBJS) $(BINS) $(OTHER_OBJS)
