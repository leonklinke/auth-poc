package middleware

type Session struct {
	Id        string
	UserID    string
	CompanyID string
	Role      Role
}

type Role string

var (
	AdminRole  Role = "admin"
	ClientRole Role = "client"
	MemberRole Role = "member"
)

func (r Role) IsAdmin() bool {
	return r == AdminRole
}

func getSessionInfo(token string) (Session, error) {
	//mock session
	if token == "tokenClient" {
		return Session{
			Id:        "1",
			UserID:    "81de5fe8-eea1-11ed-a05b-0242ac120003",
			CompanyID: "81de5fe8-eea1-11ed-a05b-0242ac120004",
			Role:      ClientRole,
		}, nil
	}
	return Session{
		Id:        "1",
		UserID:    "81de5fe8-eea1-11ed-a05b-0242ac120003",
		CompanyID: "",
		Role:      MemberRole,
	}, nil
}
