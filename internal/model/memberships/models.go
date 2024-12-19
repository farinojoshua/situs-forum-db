package memberships

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RefreshTokenRequest struct {
		Token string `json:"token"`
	}
)

type (
	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	RefreshResponse struct {
		AccessToken string `json:"access_token"`
	}
)

type (
	UserModel struct {
		UserID    int64     `db:"user_id"`
		Email     string    `db:"email"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}

	RefreshTokenModel struct {
		RefreshTokenID int64     `db:"refresh_token_id"`
		UserID         int64     `db:"user_id"`
		RefreshToken   string    `db:"refresh_token"`
		ExpiredAt      time.Time `db:"expired_at"`
		CreatedAt      time.Time `db:"created_at"`
		UpdatedAt      time.Time `db:"updated_at"`
		CreatedBy      string    `db:"created_by"`
		UpdatedBy      string    `db:"updated_by"`
	}
)
