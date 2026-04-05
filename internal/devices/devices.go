package devices

type Record struct { ID string; Owner string; State string }

func Load() []Record { return []Record{{ID: "devices-seed", Owner: "devices-owner", State: "ACTIVE"}} }
