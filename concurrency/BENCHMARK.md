# benchmark report

## method one

```log
goos: darwin
goarch: amd64
pkg: github.com/fernandoocampo/justforfun/concurrency
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkProcessDataOne-16    	  160680	      7370 ns/op	     930 B/op	      17 allocs/op
PASS
ok  	github.com/fernandoocampo/justforfun/concurrency	1.628s
```

## method two

```log
goos: darwin
goarch: amd64
pkg: github.com/fernandoocampo/justforfun/concurrency
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkProcessDataTwo-16    	  171375	      7351 ns/op	     834 B/op	      16 allocs/op
PASS
ok  	github.com/fernandoocampo/justforfun/concurrency	1.654s
```

## method three

```log
goos: darwin
goarch: amd64
pkg: github.com/fernandoocampo/justforfun/concurrency
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkProcessDataThree-16    	  138313	      8870 ns/op	    1074 B/op	      16 allocs/op
PASS
ok  	github.com/fernandoocampo/justforfun/concurrency	1.677s
```

## method four

```log
goos: darwin
goarch: amd64
pkg: github.com/fernandoocampo/justforfun/concurrency
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkProcessDataFour-16    	   96165	     12055 ns/op	    1394 B/op	      18 allocs/op
PASS
ok  	github.com/fernandoocampo/justforfun/concurrency	1.451s
```

## method five

```log
goos: darwin
goarch: amd64
pkg: github.com/fernandoocampo/justforfun/concurrency
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkProcessDataFive-16    	   69081	     15783 ns/op	    1524 B/op	      24 allocs/op
PASS
ok  	github.com/fernandoocampo/justforfun/concurrency	1.581s
```

## method six

```log
goos: darwin
goarch: amd64
pkg: github.com/fernandoocampo/justforfun/concurrency
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkProcessDataSix-16    	   97870	     12225 ns/op	    1507 B/op	      19 allocs/op
PASS
ok  	github.com/fernandoocampo/justforfun/concurrency	1.527s
```

## method seven

```log
goos: darwin
goarch: amd64
pkg: github.com/fernandoocampo/justforfun/concurrency
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkProcessDataSeven-16    	   78636	     15891 ns/op	    2196 B/op	      28 allocs/op
PASS
ok  	github.com/fernandoocampo/justforfun/concurrency	1.624s
```

## method eight

```log
goos: darwin
goarch: amd64
pkg: github.com/fernandoocampo/justforfun/concurrency
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkProcessDataEight-16    	  140268	      8189 ns/op	     800 B/op	       9 allocs/op
PASS
ok  	github.com/fernandoocampo/justforfun/concurrency	1.457s
```