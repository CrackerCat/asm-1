# {name}

{go:header}

An x86-64 assembler written in Go. It is used by the [Q programming language](https://github.com/akyoto/q) for machine code generation.

## Architectures

- [x] Linux x86-64 (ELF binaries)
- [ ] ...

## Examples

See [examples](https://github.com/akyoto/asm/tree/master/examples).

## Reference

### Registers

|  8B  |  4B  |  2B  |  1B  |
|------|------|------|------|
| rax  | eax  | ax   | al   |
| rcx  | ecx  | cx   | cl   |
| rdx  | edx  | dx   | dl   |
| rbx  | ebx  | bx   | bl   |
| rsi  | esi  | si   | sil  |
| rdi  | edi  | di   | dil  |
| rsp  | esp  | sp   | spl  |
| rbp  | ebp  | bp   | bpl  |
| r8   | r8d  | r8w  | r8b  |
| r9   | r9d  | r9w  | r9b  |
| r10  | r10d | r10w | r10b |
| r11  | r11d | r11w | r11b |
| r12  | r12d | r12w | r12b |
| r13  | r13d | r13w | r13b |
| r14  | r14d | r14w | r14b |
| r15  | r15d | r15w | r15b |

## Status

This project is currently work in progress. Contributions are welcome.

{go:footer}
