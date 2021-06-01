# YQL Elastic

YQL Elastic ( aka Why QL ) is a golang library implementing its own query language to construct Bool based Elasticsearch queries.

It is based on the Lucene query language but doesn't implement duplicate operators like `AND`, `OR`, `NOT` ect...

It is useful when Elasticsearch's query syntax shows its limit ( eg [Nested documents](https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html) ).

## Features

- Parse Lucene like query to [Olivere's Elasticsearch Client](https://gopkg.in/olivere/elastic.v7) [Query](https://pkg.go.dev/gopkg.in/olivere/elastic.v7#Query) interface.
- Supports groups
- Supports field remapping
- Supports nested fields


## Lexical declarations :

- Any term or group of term following `+` is a `Must` condition.
- Any term or group of term following `-` is a `MustNot` condition.
- Any term or group of term without preceding operator is a `Should` condition.
- A group is query inside delimiting `()`
- Groups can have sub-groups
- A term can be a single word 
- A term can be a phrase between quote 
- A term can be a field, prefixed by `:` with or without a (single word or phrase) as value
- Any field having a `>` after its `:` is a greater range condition.
- Any field having a `<` after its `:` is a lower range condition.
- Any field having a `=` after its `:` is a precise match condition.
- Any field having a `~` after its `:` is a regex condition.


## Quirks

- Field groups (eg `field:(test1 -test2)`) is directly passed to SimpleQueryString. This means `AND` `OR` 
  and other lucene operators are active for this field.

## Build tester

You need the [Go tools](https://github.com/golang/tools#downloadinstall) installed to re-generate the `*_string.go`
files if necessary.

```shell
$ go generate
$ go build ./cmd/testyql
$ ./testyql "+(test test2 test3) wrong'test'  +'test phrase phrase' --((-test4  +test34:\"323\" test14 test453) -(a:test +b:3333)) (test:)"
{"bool":{"must":[{"bool":{"should":[{"multi_match":{"fields":["default_field"],"query":"test"}},{"multi_match":{"fields":["default_field"],"query":"test2"}},{"multi_match":{"fields":["default_field"],"query":"test3"}}]}},{"multi_match":{"fields":["default_field"],"query":"'test phrase phrase'"}}],"must_not":{"bool":{"must_not":{"bool":{"must":{"multi_match":{"fields":["b"],"query":"3333"}},"should":{"multi_match":{"fields":["a"],"query":"test"}}}},"should":{"bool":{"must":{"multi_match":{"fields":["test34"],"query":"\"323\""}},"must_not":{"multi_match":{"fields":["default_field"],"query":"test4"}},"should":[{"multi_match":{"fields":["default_field"],"query":"test14"}},{"multi_match":{"fields":["default_field"],"query":"test453"}}]}}}},"should":[{"multi_match":{"fields":["default_field"],"query":"wrong'test'"}},{"bool":{"should":{"multi_match":{"fields":["test"],"query":""}}}}]}}
```

## Examples :

### Valid queries

- `test`
- `+(test test2 test3) e'test'  +'super phrase' --((-test4  +test:"323" e qe) -(a:test +b:ewe)) (test:)`
- `+(text field1:(+test -test2) -field2:text2 +test.id:43`

### Usage

From `cmd/testysql/main.go`:

```golang
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gboddin/yql-elastic"
	"log"
	"os"
)

elasticQuery, err := yql_elastic.Parse(os.Args[1],
    yql_elastic.WithDefaultFields([]string{"events.hostname","events.summary"}),
    yql_elastic.WithNestedPaths([]string{"events"}),
    yql_elastic.WithFieldMapping(map[string]string{
        "host":"events.host",
        "port":"open_ports",
        "fingerprint":"fingerprints",
        "ssl":"events.ssl",
    }),
)
if err != nil {
    log.Fatal(err)
}
querySource, err := elasticQuery.Source()
if err != nil {
    log.Fatal(err)
}
querySourceJson, err := json.Marshal(querySource)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(querySourceJson))
```
