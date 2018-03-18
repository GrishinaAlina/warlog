# warlog

A captain's log management system. Boatswain recommends.

* [Docs](#docs)
* [Installation](#installation)
* [Usage](#usage)
* [Tests](#tests)
* [License](#license)

## Docs [![GoDoc][GoDocBadge]][GoDoc]

> Documentation is available only in Russian at this time.

See GoDoc to read code documentation or [project wiki][wiki] to read common documentation.

## Installation [![Travis][TravisBadge]][Travis]

```sh
go get -u github.com/ShestakovDA/warlog
cd $GOPATH/github.com/ShestakovDA/warlog
make build
```

[How to install][win_make_ru] `make` on Windows (RUSSIAN).

## Usage [![GoReportCard][GoReportCardBadge]][GoReportCard]

```sh
./warlog
```

## Tests [![TestCoverage][CodeCovBadge]][CodeCov]

```sh
make test
```

## License [![License][LicenseBadge]](./LICENSE)

The code licensed under [The Lil License][lil_license]. See LICENSE file read it.

[wiki]: https://github.com/ShestakovDA/warlog/wiki/
[win_make_ru]: https://github.com/ShestakovDA/warlog/wiki/%D0%A3%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0-make-%D0%B2-Windows
[lil_license]: http://lillicense.org/v1.html

[GoDoc]: https://godoc.org/github.com/ShestakovDA/warlog
[Travis]: https://travis-ci.org/ShestakovDA/warlog
[GoReportCard]: https://goreportcard.com/report/github.com/ShestakovDA/warlog
[CodeCov]: https://codecov.io/gh/ShestakovDA/warlog

[GoDocBadge]: https://godoc.org/github.com/ShestakovDA/warlog?status.svg
[TravisBadge]: https://travis-ci.org/ShestakovDA/warlog.svg?style=flat-square&&branch=develop
[GoReportCardBadge]: https://goreportcard.com/badge/github.com/ShestakovDA/warlog
[CodeCovBadge]: https://codecov.io/gh/ShestakovDA/warlog/branch/develop/graph/badge.svg
[LicenseBadge]: https://img.shields.io/badge/license-Lil-blue.svg
