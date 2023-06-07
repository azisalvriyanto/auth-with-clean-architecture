package customer

type UseCase struct {
	repository *Repository
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		repository: r,
	}
}

func (u UseCase) ShowAll() ([]Customer, error) {
	return u.repository.ShowAll()
}

func (u UseCase) Create(customer *Customer) error {
	return u.repository.Create(customer)
}

func (u UseCase) Show(ID string) (*Customer, error) {
	return u.repository.Show(ID)
}

func (u UseCase) Update(ID string, customer Customer) (*Customer, error) {
	return u.repository.Update(ID, customer)
}

func (u UseCase) Destroy(ID string) (*Customer, error) {
	return u.repository.Destroy(ID)
}
