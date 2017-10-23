# covbadger

`covbadger` generates code coverage badges from Cobertura-compatible XML 
coverage reports.

### Compatibility

Several "standard" coverage libraries for various languages have Cobertura output built in:

* JavaScript supports Cobertura with [istanbul](https://istanbul.js.org/)
* Python supports Cobertura with [nosetest](http://nose.readthedocs.io/en/latest/) and [pytest](https://docs.pytest.org/en/latest/)
* Scala supports Cobertura with [scoverage](https://github.com/scoverage/scalac-scoverage-plugin)

There are also several coverage format conversion tools:

* Go has [gocov-xml](https://github.com/AlekSi/gocov-xml)
* Erlang has [covertool](https://github.com/idubrov/covertool)
* .NET has [OpenCoverToCoberturaConverter](https://github.com/danielpalme/OpenCoverToCoberturaConverter)
* JaCoCo has [cover2cover](https://github.com/rix0rrr/cover2cover)
* Finally, LCOV reports can be converted using [lcov-to-cobertura-xml](https://github.com/eriwen/lcov-to-cobertura-xml)

## License

MIT
