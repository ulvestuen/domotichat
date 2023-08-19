package status

type Status struct {
	Status string `json:"status"`
}

func GetStatus() *Status {
	return &Status{
		Status: "UP",
	}
}
