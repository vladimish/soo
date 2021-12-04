package requests

type FindUser struct {
	VerifyRegister VerifyRegister `json:"verify_register"`
	SearchQuery    string         `json:"search_query" validate:"ascii,gte=4"`
	ResultsAmount  int            `json:"results_amount" validate:"gt=0,lte=25"`
}
