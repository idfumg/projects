package server

type Strategy interface {
	Next(replica []*Replica) (*Replica, error)
}

type Service struct {
	Name     string
	Replicas []*Replica
	Strategy Strategy
}

func (s *Service) Next() (*Replica, error) {
	return s.Strategy.Next(s.Replicas)
}
