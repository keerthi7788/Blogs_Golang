package modals

type LoginModal struct {
	Email    string `json:"username,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

func (l *LoginModal) Validate() map[string]string {
	if l.Password == "" || l.Email == "" {
		return map[string]string{"email or password": "cannot be empty"}

	}
	return nil
}
