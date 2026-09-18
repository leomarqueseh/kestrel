package version

const AppVersion = "0.1.0"

type Info struct {
	Version string `json:"version"`
	Phase   string `json:"phase"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get() Info {
	return Info{
		Version: AppVersion,
		Phase:   "02-backend-architecture",
	}
}
