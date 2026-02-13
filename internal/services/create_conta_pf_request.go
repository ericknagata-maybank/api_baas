package services

type CreateContaPfRequest struct {
	NaturalPerson struct {
		DocumentNumber   string `json:"documentNumber"`
		IdentityDocument string `json:"identityDocument"`
		IssueDate        string `json:"issueDate"`
		IssuingAgency    string `json:"issuingAgency"`
		IssuingState     string `json:"issuingState"`

		PassportType           *string `json:"passportType"`
		PassportIssuingCountry *string `json:"passportIssuingCountry"`
		PassportNationality    *string `json:"passportNationality"`
		PassportPlaceOfBirth   *string `json:"passportPlaceOfBirth"`
		PassportExpirationDate *string `json:"passportExpirationDate"`

		Name            string  `json:"name"`
		Idoc            *string `json:"idoc"`
		Gender          string  `json:"gender"`
		BirthDate       string  `json:"birthDate"`
		IdMaritalStatus string  `json:"idMaritalStatus"`
		Political       bool    `json:"political"`
		Rent            float64 `json:"rent"`
		NameMother      string  `json:"nameMother"`
	} `json:"naturalPerson"`

	Address struct {
		ZipCode       string  `json:"zipCode"`
		AddressNumber string  `json:"addressNumber"`
		Address       string  `json:"address"`
		State         string  `json:"state"`
		City          string  `json:"city"`
		Neighborhood  string  `json:"neighborhood"`
		Complement    *string `json:"complement"`
	} `json:"address"`

	Contact struct {
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	} `json:"contact"`

	UserAdmin struct {
		Password string `json:"password"`
	} `json:"userAdmin"`

	ConfirmationUrl string `json:"confirmationUrl"`
}
