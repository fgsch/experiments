VERBOSE		=
CC		= $(shell brew --prefix llvm)/bin/clang
CFLAGS		= $(VERBOSE) \
		  --target=wasm32-unknown-wasi \
		  -O2 \
		  -Wl,--allow-undefined \
		  -Wl,--export-all \
		  -Wl,--import-memory \
		  -Wl,--no-entry \
		  -Wl,--shared \
		  -nostdlib
RUSTC		= rustc
RUSTC_FLAGS	= $(VERBOSE) \
		  --crate-type=lib \
		  --target=wasm32-wasi \
		  -Copt-level=2 \
		  -Crelocation-model=pic
TARGET		= main.wasm
SRCS_C		= add.c
SRCS_RS		= sub.rs
SRCS_GO		= mul.go
OBJS		= $(SRCS_C:.c=.o) $(SRCS_RS:.rs=.o)

.SUFFIXES: .rs

all: $(TARGET)

.c.o:
	$(CC) $(CFLAGS) -o $@ $?

.rs.o:
	$(RUSTC) $(RUSTC_FLAGS) -o $@ $?

$(TARGET): $(OBJS)
	tinygo build -no-debug -opt=2 -target=wasi -x -o $(TARGET)

check: $(TARGET)
	wasmtime run $(TARGET)

clean:
	rm -f $(OBJS) $(TARGET)

deps:
	brew install llvm tinygo
