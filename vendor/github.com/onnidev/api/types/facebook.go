package types

// FacebookLoginRequest sdkjfn
type FacebookLoginRequest struct {
	AccessToken string `json:"accessToken" bson:"accessToken"`
}

// FacebookSignUpRequest sdkjfn
type FacebookSignUpRequest struct {
	FirstName      string     `bson:"firstName" json:"firstName"`
	LastName       string     `bson:"lastName" json:"lastName"`
	Mail           string     `json:"mail" bson:"mail"`
	BirthDate      *Timestamp `json:"birthdate" bson:"birthdate"`
	AccessToken    string     `json:"accessToken" bson:"accessToken"`
	Phone          string     `json:"phone" bson:"phone"`
	UserName       string     `json:"username" bson:"username"`
	DocumentNumber string     `json:"documentNumber" bson:"documentNumber"`
}

// FacebookValidation sdkjnf
type FacebookValidation struct {
	Data ValidationData `json:"data"`
}

// ValidationData sdkjnf
type ValidationData struct {
	AppID       string   `json:"app_id"`
	Type        string   `json:"type"`
	Application string   `json:"application"`
	ExpiresAt   int      `json:"expires_at"`
	IsValid     bool     `json:"is_valid"`
	IssuedAt    int      `json:"issued_at"`
	Scope       []string `json:"scopes"`
	UserID      string   `json:"user_id"`
	Metadata    Metadata `json:"metadata"`
}

// Metadata sdkjfn
type Metadata struct {
	SSO string `json:"sso"`
}

// FacebookShit sdkjfnksdjfn
type FacebookShit struct {
	Birthday  string
	ID        string
	Email     string
	FirstName string
	LastName  string
	Photos    FBPhoto
}

// FBPhoto jknfsdf
type FBPhoto struct {
	Data []FBPhotoData
}

// FBPhotoData sdkjfn
type FBPhotoData struct {
	ID          string
	CreatedTime string
	Name        string
}
