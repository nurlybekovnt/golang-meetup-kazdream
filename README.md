
### Run benchmarks
```bash
go test -bench . -benchmem -memprofile mem.out
```

### See profile
```bash
go tool pprof mem.out
```
