package main

type user struct {
	ID       int    // Perlu diberikan penjelasan tentang maksud dan penggunaan ID
	Username string // Seharusnya tipe data string
	Password string // Seharusnya tipe data string
}

type userService struct {
	users []user // Lebih deskriptif
}

func (u *userService) getAllUsers() []user {
	return u.users
}

func (u *userService) getUserByID(id int) *user {
	for _, r := range u.users {
		if id == r.ID {
			return &r // Mengembalikan pointer ke user yang ditemukan
		}
	}
	return nil // Mengembalikan nil jika tidak ada hasil yang cocok
}
