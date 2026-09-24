# md5 Implemented in go

go implementation of the md5 checksum algorithm, converted from the
[C implementation at https://github.com/bahamas10/c-md5](https://github.com/bahamas10/c-md5)

## Usage

```
go build -o md5 cmd/
./md5 file.txt
./md5 < file.bin
echo 'hello' | ./md5
./md5 -debug file.txt
./md5 -memory=1024 file.txt
```

## Benchmark

```
❯ go test -bench=BenchmarkHash -benchtime=10s -cpuprofile=cpu.prof github.com/gdey/md5-go
goos: darwin
goarch: arm64
pkg: github.com/gdey/md5-go
cpu: Apple M2 Max
BenchmarkHash/empty.txt-12              77531754               142.6 ns/op
BenchmarkHash/hello.txt-12              84914096               142.0 ns/op
BenchmarkHash/random_4kb.bin-12         85146640               142.0 ns/op
PASS
ok      github.com/gdey/md5-go  35.435s
```

## TODO/Wants 

Reading from the stream is slower then the native implementation, tops out around ~265MiB, not sure if that can be improved.


```

❯ pv  Qwen2.5-Omni-3B-GGUF/Qwen2.5-Omni-3B-Q8_0.gguf | time ./md5sum -memory 1024
3.37GiB 0:00:07 [ 445MiB/s] [========================================================================================================================================================================>] 100%
6699017cba6687c49be285590d54b32a <<StandardIn>>
./md5sum -memory 1024  7.58s user 0.17s system 100% cpu 7.748 total

❯ pv  Qwen2.5-Omni-3B-GGUF/Qwen2.5-Omni-3B-Q8_0.gguf | time md5sum
3.37GiB 0:00:06 [ 548MiB/s] [========================================================================================================================================================================>] 100%
6699017cba6687c49be285590d54b32a  -
md5sum  5.95s user 0.22s system 97% cpu 6.299 total

```

License
-------

[MIT License](LICENSE)
