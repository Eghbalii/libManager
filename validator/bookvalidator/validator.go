package bookvalidator

type Repository interface {
	IsISBNUnique(isbn string) bool
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
