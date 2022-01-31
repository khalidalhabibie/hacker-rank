package warm

import "testing"

func TestSetMyCountry(t *testing.T) {

	users := []User{
		User{
			ID:        1,
			Username:  "user 1",
			CountryID: 10,
		},
		User{
			ID:        2,
			Username:  "user 2",
			CountryID: 50,
		},
		User{
			ID:        3,
			Username:  "user 3",
			CountryID: 40,
		},
		User{
			ID:        4,
			Username:  "user 4",
			CountryID: 20,
		},
		User{
			ID:        5,
			Username:  "user 5",
			CountryID: 30,
		},
	}

	countries := []Country{
		Country{
			ID:   10,
			Name: "country 10",
		},
		Country{
			ID:   20,
			Name: "country 20",
		},
		Country{
			ID:   30,
			Name: "country 30",
		},
		Country{
			ID:   40,
			Name: "country 40",
		},
		Country{
			ID:   50,
			Name: "country 50",
		},
	}

	got := SetMyCountry(users, countries)
	resultIndex5 := ResultMyCountry{Username: "user 5", CountryName: "country 30"}
	if got[4] != resultIndex5 {
		t.Errorf("idexing 5, we want %v but we got %v", got[4], resultIndex5)

	}

}
