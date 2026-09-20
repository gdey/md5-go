# Various Tools use to help with optimizing the code:

## Profiling for memory and alloc 


```
 go test -bench=. -benchmem
```


## Profiling for cpu 

generate the profile:


```
go test -bench=BenchmarkHash -benchtime=10s -cpuprofile=cpu.prof github.com/gdey/md5-go
```

explore it:

```
 go tool pprof cpu.prof
```

or 

```
 go tool pprof -http=:8080 cpu.prof
```

## disassemble

```
go build -gcflags="-S" > block_asm.txt 2>&1

```
