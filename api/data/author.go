package data

import "time"

// Author defines the structure of the author api
type Author struct {
	ID               int      `json:"id"`
	CourseID         []int    `json:"courseid"`
	Email            string   `json:"email"`
	Twitter          string   `json:"twitter"`
	Facebook         string   `json:"facebook"`
	Instagram        string   `json:"instagram"`
	Location         string   `json:"location"`
	Fullname         string   `json:"fullname"`
	ShortDescription string   `json:"shortdescription"`
	Speciality       string   `json:"speciality"`
	Features         []string `json:"features"`
	CreatedOn        string   `json:"-"`
}

// Authors is a slice of authors
type Authors []*Author

// courseList is a hard coded list of courses / test data source
var authorsList = []*Author{
	{
		ID:               1,
		CourseID:         []int{1, 2, 3},
		Email:            "topless@datacademy.net",
		Twitter:          "twitter.com",
		Facebook:         "facebook.com",
		Instagram:        "instagram.com",
		Location:         "Vancouver, Canada",
		Fullname:         "Dmitry Anoshin",
		ShortDescription: "Founder and CEO of Datacademy",
		Speciality:       "Big data",
		Features: []string{"Data Engineer in Amazon, Alexa AI",
			"10+ years of experience in Analytics (Russia, Europe, Canada and USA)",
			"Organizer of Vancouver Tableau User Group, Snowflake Canada User Group, Amazon Tableau User Group and Amazon BI Tech Talks",
			"University of Victoria lecturer - Cloud Computing",
			"Author of 6 books about Analytics",
			"Develops consaltyng in North America - rockyourdata.cloud",
			"Speaker on conferences and meetups in Russia and North America",
			"Courses are being developed on the West Coast"},
		CreatedOn: time.Now().UTC().String(),
	},
	{
		ID:               2,
		CourseID:         []int{4, 5},
		Email:            "bottomless@datacademy.net",
		Twitter:          "twitter.com",
		Facebook:         "facebook.com",
		Instagram:        "instagram.com",
		Location:         "Moscow, Russia",
		Fullname:         "Roman Ponomarev",
		ShortDescription: "Co-founder and admin of the project",
		Speciality:       "Small data",
		Features: []string{"Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?"},
		CreatedOn: time.Now().UTC().String(),
	},
	{
		ID:               3,
		CourseID:         []int{1, 2, 3},
		Email:            "fullydressed@datacademy.net",
		Twitter:          "twitter.com",
		Facebook:         "facebook.com",
		Instagram:        "instagram.com",
		Location:         "Moscow, Russia",
		Fullname:         "Sergei Isaev",
		ShortDescription: "Just a random guy",
		Speciality:       "Doing nothing",
		Features:         []string{"Lorem ipsum dolor sit amet.", "Lorem ipsum dolor sit amet.", "Lorem ipsum dolor sit amet."},
		CreatedOn:        time.Now().UTC().String(),
	},
}

// GetAuthors returns list of authors
func GetAuthors() Authors {
	return authorsList
}

// findIndexByCourseID finds the index of a course in the DB
// returns -1 when no course isfound
func findIndexByAuthorsID(id int) int {
	for i, a := range authorsList {
		if a.ID == id {
			return i
		}
	}
	return -1
}

/* TODO: make this function available for all handlers */

// getNextId generates a new id for a product being inserted into db
func getNextAuthorID() int {
	lastItemInDB := authorsList[len(authorsList)-1]
	return lastItemInDB.ID + 1
}

// GetAuthorByID returns a single author which matches the id from the DB
// if an author is not found this function returns an AuthorNotFound error
func GetAuthorByID(id int) (*Author, error) {
	index := findIndexByAuthorsID(id)
	if index == -1 {
		return nil, ErrorAuthorNotFound
	}
	return authorsList[index], nil
}

/* TODO: currently if some filds are comming empty - it makes corresponding fields in data set empty as well - needs to be fixed */

// UpdateAuthor replaces an author in the DB with the given item.
// If an author with the given id does not exist in th DB
// this function returns AuthorNotFound error
func UpdateAuthor(a Author) error {
	index := findIndexByAuthorsID(a.ID)
	if index == -1 {
		return ErrorAuthorNotFound
	}
	authorsList[index] = &a
	return nil
}

// AddAuthor adds a new author in the DB
func AddAuthor(a Author) {
	a.ID = getNextAuthorID()
	authorsList = append(authorsList, &a)
}

// DeleteAuthor deletes the author from the DB
func DeleteAuthor(id int) error {
	index := findIndexByAuthorsID(id)
	if index == -1 {
		return ErrorAuthorNotFound
	}
	authorsList = append(authorsList[:index], authorsList[index+1])
	return nil
}
