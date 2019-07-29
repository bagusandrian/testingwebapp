package users

type (
	Users struct {
		Username  string `json:"username", db:"username"`
		UserID    string `json:"user_id", db:"user_id"`
		Name      string `json:"name", db:"name"`
		Password  string `json:"password", db:"password"`
		LastLogin string `json:"last_login", db:"last_login"`
		BirthDate string `json:"birth_date", db:"birth_date"`
		Address   string `json:"address", db:"address"`
		Gender    int    `json:"gender", db:"gender"`
		RoleID    int    `json:"role_id", db:"role_id"`
	}

	ResponseJSON struct {
		Data []Users `json:"data"`
	}
)
