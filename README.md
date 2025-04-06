# html2pug-go (experimental)

<img src="assets/go-pug-html.png" alt="Gopher and Pug working on HTML" style="width: 256px; height: 256px;" />

A utility library to map HTML to Pug equivalent, inspired by and partly ported from [html2jade](https://github.com/donpark/html2jade)

## Using the library

Examples can be found within ./examples/

## Running the tests

```bash
go test -v ./test/...
```
Note: some tests are failing as I fix the issues with the port, so I've added a flag to skip

## Issues

Work on html2pug-go is currently ongoing, so YMMV - feel free to raise issues or submit a PR (with supporting tests)

## Todo

- Fix unit tests (doSkip has been added to failing tests)
- Support rendering of partials
- Increase coverage of unit tests
- CLI app