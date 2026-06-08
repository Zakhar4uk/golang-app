package users_posgres_repository

type UserModel struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}
