package enrollment

type Record struct { ID string; Owner string; State string }

func Load() []Record { return []Record{{ID: "enrollment-seed", Owner: "enrollment-owner", State: "ACTIVE"}} }
