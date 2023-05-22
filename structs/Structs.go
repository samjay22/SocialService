package structs

//Internal usage
type UserProfile struct {
	ID       int64
	JoinDate string
	Name     string
	Age      int
	Role     string

	//Friends list one to many based on ID. Less storage this way.
	Friends []int64

	//Posts list one to many based on ID. Less storage this way.
	Posts []int64
}

type Reaction struct {
	ID   int64
	Type string
}

type Comment struct {
	ID      int64
	Author  int64
	Content string

	//Reactions list one to many based on ID. Less storage this way.
	Reactions []int64
}

type Post struct {
	ID      int64
	Author  int64
	Content string

	//Reactions list one to many based on ID. Less storage this way.
	Reactions []int64
	//Reactions list one to many based on ID. Less storage this way.
	Comments []int64
}

type Session struct {
	ID string
}

//requests

type AddFriendRequest struct {
	UserId   int64 `json:"userId"`
	FriendId int64 `json:"friendId"`
}

type GetUserRequest struct {
	UserId int64 `json:"userId"`
}

type CreateUserProfileRequest struct {
	ID       int64  `json:"id"`
	JoinDate string `json:"join_date"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	SessionID string `json:"sessionId"`
}

type VerifySessionRequest struct {
	SessionID string `json:"sessionId"`
}

//responses

func NewUserProfile(id int64, joinDate string, name string, age int) *UserProfile {
	return &UserProfile{
		ID:       id,
		JoinDate: joinDate,
		Name:     name,
		Age:      age,
		Role:     "Standard",
		Friends:  []int64{},
		Posts:    []int64{},
	}
}
