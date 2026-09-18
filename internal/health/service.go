package health

import "time"

type Status struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Time    string `json:"time"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Check() Status {
	return Status{
		Status:  "ok",
		Service: "kestrel-api",
		Time:    time.Now().UTC().Format(time.RFC3339),
	}
}
