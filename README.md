# go-evmasm

Package **evmasm** provides tools for writing an assembler that turns assembly language into bytecodes for the Ethereum Virtual Machine (EVM), for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-evmasm

[![GoDoc](https://godoc.org/github.com/reiver/go-evmasm?status.svg)](https://godoc.org/github.com/reiver/go-evmasm)


## Example

```go
assemblyCode := `push1 3 push1 2 add push1 1 mul`

var evmByteCodes bytes.Buffer

err := evmasm.Assemble(&evmByteCodes, assemblyCode)
```

The buffer evmByteCodes would then contain the equivalent of:
```go
[]byte{
	0x60, 0x03, // push1 3
	0x60, 0x02, // push1 2
	0x01,       // add
	0x60, 0x01, // push1 1
	0x02,       // mul
}
```
