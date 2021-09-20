package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LeakIX/yql-elastic"
	"log"
	"os"
)

func main() {
	hasAuth := "no"
	elasticQuery, err := yql_elastic.Parse(os.Args[1],
		yql_elastic.WithDefaultFields([]string{"events.hostname", "events.summary"}),
		yql_elastic.WithNestedPaths([]string{"events"}),
		yql_elastic.WithFieldMapping(map[string]string{
			"host":        "events.host",
			"port":        "open_ports",
			"fingerprint": "fingerprints",
			"ssl":         "events.ssl",
		}),
		yql_elastic.WithFieldCallBack("plugin", func(text string) (string, error) {
			if hasAuth != "yes" {
				return "$!#@#", errors.New("plugin field not allowed")
			}
			return text, nil
		}),
	)
	if err != nil {
		if err, isFieldError := err.(yql_elastic.FieldCallBackError); isFieldError {
			log.Fatalf("field error %s : %s", err.FieldName, err.FieldValue)
		}
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
	log.Println("done")
}
