package resources

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/javiersvg/fight-of-heroes-service/aggregates"
)

var re = regexp.MustCompile(`/query/(.*)?`)

func QueriesResource(fights *aggregates.Fights) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			id := re.FindSubmatch([]byte(r.URL.EscapedPath()))
			if len(id[1]) > 0 {
				json.NewEncoder(w).Encode(fights.GetFight(string(id[1])))
			} else {
				json.NewEncoder(w).Encode(fights.GetFights())
			}
		}
	}
}
