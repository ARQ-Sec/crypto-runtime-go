package bootstrap

type Record struct { ID string; Owner string; State string }

func Load() []Record { return []Record{{ID: "bootstrap-seed", Owner: "bootstrap-owner", State: "ACTIVE"}} }
