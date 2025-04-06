# html2pug-go (experimental)

<img src="assets/go-pug-html.png" alt="Gopher and Pug working on HTML" style="width: 256px; height: 256px;" />

A utility library to map HTML to Pug equivalent, inspired by and partly ported from [html2jade](https://github.com/donpark/html2jade)

## Command line usage

Outputs to stdout if input is URL

```bash
html2pug-go http://twitter.com
html2pug-go http://twitter.com > twitter.jade
```

Outputs to file if input is file

```bash
html2pug-go mywebpage.html # outputs mywebpage.jade
html2pug-go public/*.html  # converts all .html files to .jade
```

Convert HTML from stdin

```bash
cat mywebpage.html | html2pug-go -
```

To generate Scalate compatible output:

```bash
html2pug-go --scalate http://twitter.com
html2pug-go --scalate http://twitter.com > twitter.jade
html2pug-go --scalate mywebpage.html
html2pug-go --scalate public/*.html
```

### Command-line Options

- `-d`, `--double` - use double quotes for attributes
- `-o`, `--outdir <dir>` - path to output generated pug file(s) to
- `-n`, `--nspaces <n>` - the number of spaces to indent generated files with. Default is 2 spaces
- `-t`, `--tabs` - use tabs instead of spaces
- `--donotencode` - do not html encode characters. This is useful for template files which may contain expressions like {{username}}
- `--bodyless` - do not output enveloping html and body tags
- `--numeric` - use numeric character entities
- `-s`, `--scalate` - generate Scalate variant of pug syntax
- `--noattrcomma` - omit attribute separating commas
- `--noemptypipe` - omit lines with only pipe ('|') printable character
