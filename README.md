# CGo Bridge 🌉

Type-safe Go/Rust FFI bridge.

## Features

- **Type Safety**: Compile-time checks
- **Memory Management**: Automatic GC integration
- **Error Propagation**: Go error ↔ Rust Result
- **Zero-copy**: Shared memory where possible

## Quick Start

```go
// Go side
result, err := rustlib.Process(input)
```

```rust
// Rust side
#[no_mangle]
pub extern "C" fn process(input: *const c_char) -> *mut c_char { ... }
```

## License

MIT