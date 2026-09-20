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
go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/gdey/md5-go
cpu: Apple M2 Max
BenchmarkHash/empty.txt-12               1446352               829.0 ns/op             0 B/op          0 allocs/op
BenchmarkHash/hello.txt-12                786171              1273 ns/op               0 B/op          0 allocs/op
BenchmarkHash/random_4kb.bin-12            79603             15114 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/gdey/md5-go  3.720s

```

## TODO/Wants 

Reading from the stream is slower then the native implementation, tops out around ~265MiB, not sure if that can be improved.

```
> pv Qwen2.5-Omni-3B-Q8_0.gguf | ./md5 -memory 1024
3.37GiB 0:00:12 [ 265MiB/s] [====================================================================================================================>] 100%
6699017cba6687c49be285590d54b32a <<StandardIn>>

> pv Qwen2.5-Omni-3B-Q8_0.gguf | md5sum
3.37GiB 0:00:06 [ 544MiB/s] [====================================================================================================================>] 100%
6699017cba6687c49be285590d54b32a  -

```

License
-------

[MIT License](LICENSE)
