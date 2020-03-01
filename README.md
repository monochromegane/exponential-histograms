# Exponential histograms [![Actions Status](https://github.com/monochromegane/exponential-histograms/workflows/Go/badge.svg)](https://github.com/monochromegane/exponential-histograms/actions)


Exponential histograms is a data structure for sliding windows. It is from `Maintaining Stream Statistics over Sliding Windows, M.Datar, A.Gionis, P.Indyk, R.Motwani; ACM-SIAM, 2002`.

See also
- https://www.dbs.ifi.lmu.de/Lehre/BigData-Management&Analytics/WS15-16/Chapter-5_Stream_Processing_part2.pdf
- http://www.facom.ufu.br/~elaine/disc/MFCD/histograms.pdf

## Usage

Exponential histograms requires windowSize and epsilon for parameter.
The epsilon controls timing of merging the oldest two buckets into one bucket double the size.

Exponential histograms estimates count value in the window.
The absolute error in the value is at most C/2, where C is the size of the last bucket.

### Bits

```go
	hist := New(10, 0.5)
	stream := []uint{1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0}
	for _, x := range stream {
		hist.Add(x)
	}

	count := hist.Count()
	fmt.Println(count) // 6.0 (Actual: 8.0)
```

### Positive integers

```go
	hist := New(200, 0.01)
	for i := 1; i <= 200; i++ {
		hist.Add(uint(i))
	}

	count := hist.Count()
	fmt.Println(count) // 19972.0 (Actual: 20100.0)
```

## Installation

```sh
$ go get github.com/monochromegane/exponential-histograms
```

## License

[MIT](https://github.com/monochromegane/exponential-histograms/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)
