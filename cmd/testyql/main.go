package main

import (
	"encoding/json"
	"fmt"
	"github.com/LeakIX/yql-elastic"
	"log"
	"os"
)


func main() {
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
}