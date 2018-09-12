# fuzzing

## Automated testing with go-fuzz

Install:

```bash
$ go get -u github.com/dvyukov/go-fuzz/...
$ go get -d github.com/dvyukov/go-fuzz-corpus
```

Basic:

```bash
$ go-fuzz-build github.com/alextanhongpin/go-fuzz
$ mkdir -p workdir/corpus
$ echo -ne '\x0cHello World!' > workdir/corpus/example
$ go-fuzz -bin=bug-fuzz.zip -workdir=workdir
```

## With google's gofuzz

```bash
$ go get github.com/google/gofuzz
```
