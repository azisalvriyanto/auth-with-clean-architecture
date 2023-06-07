package user

type UseCase struct {
	repository *Repository
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		repository: r,
	}
}

func (u UseCase) ShowAll() ([]User, error) {
	return u.repository.ShowAll()
}

func (u UseCase) Create(user *User) error {
	return u.repository.Create(user)
}

func (u UseCase) Show(ID string) (*User, error) {
	return u.repository.Show(ID)
}

func (u UseCase) Update(ID string, user User) (*User, error) {
	return u.repository.Update(ID, user)
}

func (u UseCase) Destroy(ID string) (*User, error) {
	return u.repository.Destroy(ID)
}
