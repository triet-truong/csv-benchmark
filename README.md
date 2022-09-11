# CSV benchmark
A simple CLI application to benchmark two approaches of parsing large CSV file:
- gocsv: Read whole file, load whole content to memory before iterating every line. (`readCsvChanOld`) 
- csvutil: read file and decode line by line. (`readCsvFile`)

```shell
go test --bench=. --benchmem --count 3
```

Result screenshot:
![alt text](result.png)
