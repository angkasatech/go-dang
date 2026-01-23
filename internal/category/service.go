package category

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() ([]Category, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id int) (*Category, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(c *Category) error {
	return s.repo.Create(c)
}

func (s *Service) Update(id int, c *Category) error {
	return s.repo.Update(id, c)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
