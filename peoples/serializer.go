package peoples

import "time"

type PeopleSerializer struct {
	// c *gin.Context
	People
}

type PeoplesSerializer struct {
	// c       *gin.Context
	Peoples []People
}

type PeopleResponse struct {
	NIK         string `json:"nik"`
	Name        string `json:"name"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Photo       string `json:"photo"`
	RT          string `json:"rt"`
	RW          string `json:"rw"`
	Subdistrict string `json:"subdistrict"`
	District    string `json:"district"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Age         int    `json:"age"`
}

func (s *PeopleSerializer) Response() PeopleResponse {
	dateParse, _ := time.Parse(time.RFC3339, s.People.Birthday)
	s.People.Birthday = dateParse.Format("2006-01-02")
	response := PeopleResponse{
		NIK:         s.People.NIK,
		Name:        s.People.Name,
		Birthday:    s.People.Birthday,
		Gender:      s.People.Gender,
		Address:     s.People.Address,
		Photo:       s.People.Photo,
		RT:          s.People.RT,
		RW:          s.People.RW,
		Subdistrict: s.People.Subdistrict,
		District:    s.People.District,
		City:        s.People.City,
		Province:    s.People.Province,
		Age:         s.People.Age,
	}

	return response
}

func (s *PeoplesSerializer) Response() []PeopleResponse {
	response := []PeopleResponse{}
	for _, people := range s.Peoples {
		serializer := PeopleSerializer{people}
		response = append(response, serializer.Response())
	}

	return response
}
