package policies

type Record struct { ID string; Owner string; State string }

func Load() []Record { return []Record{{ID: "policies-seed", Owner: "policies-owner", State: "ACTIVE"}} }
