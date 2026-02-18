.PHONY: build test fmt lint clean run

BINARY_NAME=openkit
RUST_MANIFEST=rust-cli/Cargo.toml

all: build

build:
	cargo build --release --manifest-path $(RUST_MANIFEST)
	cp rust-cli/target/release/$(BINARY_NAME) ./$(BINARY_NAME)

test:
	cargo test --manifest-path $(RUST_MANIFEST)

fmt:
	cargo fmt --manifest-path $(RUST_MANIFEST) --all

lint:
	cargo clippy --manifest-path $(RUST_MANIFEST) --all-targets -- -D warnings

run:
	cargo run --manifest-path $(RUST_MANIFEST) -- --help

clean:
	rm -f ./$(BINARY_NAME)
	cargo clean --manifest-path $(RUST_MANIFEST)
