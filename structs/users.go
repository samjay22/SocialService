package structs

type UserProfile struct {
	JoinDate string `json:"join_date"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Bio      string `json:"bio"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Username string
}

func NewUserProfile(joinDate string, avatar string, name string, email string, phone string, address string, bio string, age int, gender string) *UserProfile {
	return &UserProfile{JoinDate: joinDate, Avatar: avatar, Name: name, Email: email, Phone: phone, Address: address, Bio: bio, Age: age, Gender: gender}
}
