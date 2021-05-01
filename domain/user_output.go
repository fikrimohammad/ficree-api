package domain

// UserCompactOutput represents compact form of user output
type UserCompactOutput struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Title          string `json:"title"`
	ProfilePicture string `json:"profile_picture"`
}

func (u *User) AsUserCompactOutput() *UserCompactOutput {
	return &UserCompactOutput{
		ID:             u.ID,
		Name:           u.Name,
		Email:          u.Email,
		Title:          u.Title,
		ProfilePicture: u.ProfilePicture,
	}
}

// UserDetailOutput represents detail form of user output
type UserDetailOutput struct {
	UserCompactOutput
	PhoneNumber string `json:"phone_number"`
	GithubURL   string `json:"github_url"`
	LinkedinURL string `json:"linkedin_url"`
	TwitterURL  string `json:"twitter_url"`
	Summary     string `json:"summary"`
}

func (u *User) AsUserDetailOutput() *UserDetailOutput {
	return &UserDetailOutput{
		UserCompactOutput: UserCompactOutput{
			ID:             u.ID,
			Name:           u.Name,
			Email:          u.Email,
			Title:          u.Title,
			ProfilePicture: u.ProfilePicture,
		},
		PhoneNumber: u.PhoneNumber,
		GithubURL:   u.GithubURL,
		LinkedinURL: u.LinkedinURL,
		TwitterURL:  u.TwitterURL,
		Summary:     u.Summary,
	}
}
