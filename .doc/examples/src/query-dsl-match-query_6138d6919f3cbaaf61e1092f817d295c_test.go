// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/match-query.asciidoc#L175>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "match" : {
//             "message" : {
//                 "query" : "this is a test",
//                 "operator" : "and"
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_match_query_6138d6919f3cbaaf61e1092f817d295c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6138d6919f3cbaaf61e1092f817d295c[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "message": {
		        "query": "this is a test",
		        "operator": "and"
		      }
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:6138d6919f3cbaaf61e1092f817d295c[]
}
