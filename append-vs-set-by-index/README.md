```shell
$ go test -bench .
```

````
BenchmarkCopyUsingAppendWithoutAlloc-8             39045             28258 ns/op          259314 B/op         16 allocs/op
BenchmarkCopyUsingAppend-8                        141244              8473 ns/op           65537 B/op          1 allocs/op
BenchmarkCopyByIndex-8                            145719              8166 ns/op           65537 B/op          1 allocs/op
````