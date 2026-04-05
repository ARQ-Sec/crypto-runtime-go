package manifests

type Record struct { ID string; Owner string; State string }

func Load() []Record { return []Record{{ID: "manifests-seed", Owner: "manifests-owner", State: "ACTIVE"}} }
