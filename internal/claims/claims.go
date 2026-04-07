package claims

type Record struct { ID string; Owner string; State string }

func Load() []Record { return []Record{{ID: "claims-seed", Owner: "claims-owner", State: "ACTIVE"}} }
