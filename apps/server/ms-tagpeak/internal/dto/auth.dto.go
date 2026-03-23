package dto

type CreateUserDto struct {
	Email         string  `json:"email" validate:"required,email"`
	Password      string  `json:"password" validate:"required,min=8"`
	FirstName     string  `json:"firstName" validate:"required"`
	LastName      string  `json:"lastName" validate:"required"`
	Country       *string `json:"country"`
	Currency      string  `json:"currency"`
	UtmParams     *string `json:"utmParams"`
	ReferralCode  *string `json:"referralCode"`
	ReferralClick *string `json:"referralClick"`
	IsShop        *bool   `json:"isShop"`
}

type CreateUserFromMigrationDto struct {
	IdentifyProvider *string `json:"identifyProvider"`
	FederatedUserId  *string `json:"federatedUserId"`
	Country          string  `json:"country"`
	Currency         string  `json:"currency"`
	Email            string  `json:"email"`
	FirstName        string  `json:"firstName"`
	LastName         string  `json:"lastName"`
	Password         string  `json:"password"`
	BirthDate        *string `json:"birthDate"`
	DisplayName      string  `json:"displayName"`
	ProfilePic       *string `json:"profilePic"`
	LegacyId         int64   `json:"legacyId"`
	RefPercent       string  `json:"refPercent"`
	UserPercent      string  `json:"userPercent"`
	BadgeType        string  `json:"badgeType"`
	RefCode          string  `json:"refCode"`
	Newsletter       string  `json:"newsletter"`
}

type UpdateUserDto struct {
	FirstName          *string            `json:"firstName,omitempty"`
	LastName           *string            `json:"lastName,omitempty"`
	Country            *string            `json:"country,omitempty"`
	Currency           *string            `json:"currency,omitempty"`
	Balance            *string            `json:"balance,omitempty"`
	Groups             *[]string          `json:"groups,omitempty"`
	DisplayName        *string            `json:"displayName,omitempty"`
	BirthDate          *string            `json:"birthDate,omitempty"`
	IsVerified         *string            `json:"isVerified,omitempty"`
	OnboardingFinished *string            `json:"onboardingFinished,omitempty"`
	UtmParams          *string            `json:"utmParams,omitempty"`
	Newsletter         *bool              `json:"newsletter,omitempty"`
	Source             *string            `json:"source,omitempty"`
	CurrencySelected   *bool              `json:"currencySelected,omitempty"`
	EmailExtras        *map[string]string `json:"emailExtras,omitempty"`
}

type ResetPasswordDto struct {
	Email string `json:"email" validate:"required,email"`
}

type EmailVerifiedDto struct {
	UserId     string `json:"userId,omitempty"`
	IsVerified string `json:"isVerified,omitempty"`
}

type UserDto struct {
	Uuid      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CurrencySetDto struct {
	CurrencyCode string `json:"currencyCode"`
}

type SocialProfileFinish struct {
	Currency      string  `json:"currency"`
	FirstName     *string `json:"firstName,omitempty"`
	LastName      *string `json:"lastName,omitempty"`
	ReferralCode  *string `json:"referralCode"`
	ReferralClick *string `json:"referralClick"`
	UtmParams     *string `json:"utmParams"`
}
