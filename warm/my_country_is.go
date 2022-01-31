package warm

type User struct {
	ID        int
	Username  string
	CountryID int
}

type Country struct {
	ID   int
	Name string
}

type ResultMyCountry struct {
	Username    string
	CountryName string
}

// optimize using indexing method like elasticsearch
func SetMyCountry(users []User, countries []Country) []ResultMyCountry {

	// indexing country
	indexCountry := make(map[int]string)

	for i := range countries {
		indexCountry[countries[i].ID] = countries[i].Name

	}

	// set result , try find country id and name
	result := []ResultMyCountry{}

	for i := range users {
		resultI := ResultMyCountry{
			Username:    users[i].Username,
			CountryName: indexCountry[users[i].CountryID],
		}

		result = append(result, resultI)

	}

	return result

}
