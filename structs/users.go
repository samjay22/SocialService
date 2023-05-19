package structs

type UserProfile struct {
	ID       string `json:"id"`
	JoinDate string `json:"join_date"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Age      int    `json:"age"`
}

func NewUserProfile(id string, joinDate string, name string, phone string, age int) *UserProfile {
	return &UserProfile{
		ID:       id,
		JoinDate: joinDate,
		Name:     name,
		Phone:    phone,
		Age:      age,
	}
}
