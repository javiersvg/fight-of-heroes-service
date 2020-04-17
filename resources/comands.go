package resources

import (
	"encoding/json"
	"net/http"

	"github.com/javiersvg/fight-of-heroes-service/aggregates"
)

type RequestCommand struct {
	Name string
	Data struct {
		ID     string
		Heroes []string
	}
}

func CommandResource(fights *aggregates.Fights) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var rc RequestCommand

			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			err := dec.Decode(&rc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			switch rc.Name {
			case "CreateFight":
				fights.CreateFight(rc.Data.Heroes)
				w.WriteHeader(http.StatusOK)
			case "UpdateFight":
				fights.UpdateFights(rc.Data.ID, rc.Data.Heroes)
				w.WriteHeader(http.StatusOK)
			case "DeleteFight":
				fights.DeleteFight(rc.Data.ID)
				w.WriteHeader(http.StatusOK)
			default:
				http.Error(w, "Unnable to create requested command: "+rc.Name, http.StatusBadRequest)
			}

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
