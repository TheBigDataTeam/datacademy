package authors

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Serj1c/datalearn/api/pkg/util"
)

// Repo represents DB
type Repo struct {
	db *sql.DB
}

// NewRepo returns an instance of a Repo
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// ErrorAuthorNotFound is an error raised when a author can not be found
var ErrorAuthorNotFound = fmt.Errorf("Author not found")

// GetAuthors returns list of authors
func (r *Repo) GetAuthors() ([]*Author, error) {
	/* authors := make([]*Author, 0, 10)
	rows, err := r.db.Query("SELECT id, courseid, email, twitter, facebook, instagram, location, fullname, bio, shortdescription, speciality, features FROM authors")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		author := &Author{}
		err := rows.Scan(&author.ID, &author.CourseID, &author.Email, &author.Twitter, &author.Facebook, &author.Instagram, &author.Location,
			&author.Fullname, &author.Bio, &author.ShortDescription, &author.Speciality, &author.Features)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	} */
	return authorsList, nil
}

// findIndexByAuthorID finds the index of a course in the DB
// returns -1 when no author is found
func findIndexByAuthorsID(id string) int {
	for i, a := range authorsList {
		if a.ID == id {
			return i
		}
	}
	return -1
}

/* TODO: make this function available for all handlers */

// getNextId generates a new id for a product being inserted into db
func getNextAuthorID() string {
	newID := util.RandString()
	return newID
}

// GetAuthorByID returns a single author which matches the id from the DB
// if an author is not found this function returns an AuthorNotFound error
func (r *Repo) GetAuthorByID(id string) (*Author, error) {
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
func (r *Repo) UpdateAuthor(a Author) error {
	index := findIndexByAuthorsID(a.ID)
	if index == -1 {
		return ErrorAuthorNotFound
	}
	authorsList[index] = &a
	return nil
}

// AddAuthor adds a new author in the DB
func (r *Repo) AddAuthor(a Author) {
	a.ID = getNextAuthorID()
	authorsList = append(authorsList, &a)
}

// DeleteAuthor deletes the author from the DB
func (r *Repo) DeleteAuthor(id string) error {
	index := findIndexByAuthorsID(id)
	if index == -1 {
		return ErrorAuthorNotFound
	}
	authorsList = append(authorsList[:index], authorsList[index+1])
	return nil
}

// courseList is a hard coded list of courses / test data source
var authorsList = []*Author{
	{
		ID:        "qqqqq",
		CourseID:  "1",
		Email:     "topless@datalearn.biz",
		Twitter:   "twitter.com",
		Facebook:  "facebook.com",
		Instagram: "instagram.com",
		Location:  "Vancouver, Canada",
		Fullname:  "Dmitry Anoshin",
		Bio: `Меня зовут Дмитрий Аношин и я работаю в Amazon c 2016 года. С июня 2019 я присоединился к Alexa AI - Natural Understanding, 
		это команда которая создает data products для всего Amazon, кто взаимодействует с Alexa. Задача - повышение качества работы Alexa.
		До этого я 3 года поддерживал Amazon Marketplace (Книги и коллекционные предметы). Благодаря участию во всех communities в Amazon 
		(Веду Amazon Tableau User Group (2k+) и BI Tech Talk (100+ teams globally)), я смог присоединиться к одной из самых крутых команд 
		с точки зрения Data Engineering, которая находится в Бостоне и работать удаленно с острова Ванкувер.
		Я написал несколько книг по аналитике про BI инструменты и облачную аналитику (https://www.amazon.com/Dmitry-Anoshin/e/B01A5PVT2M)
		Я прекрасно понимаю, что работать на дядю вечно не получиться, поэтому с единомышленниками мы создаем консалтинг Rock Your Data 
		(https://rockyourdata.cloud/) (примером для меня служит Slalom, McKinsey и их дочернее предприятие QuantumBlack 
		(https://www.quantumblack.com/)). Фокус на Cloud Analytics and Analytics Modernization.	
		Сейчас создаю 3 курса (информация о них вы можете найти на страницах данного сайта). Они будут бесплатные, и будет отдельный 
		чат в Slack для каждого из них, где участники смогут помогать друг другу. Такой вот community driven.`,
		ShortDescription: "Founder and CEO of Datacademy",
		Speciality:       "Big data",
		Features: `#Data Engineer in Amazon, Alexa AI,
			#10+ years of experience in Analytics (Russia, Europe, Canada and USA),
			#Organizer of Vancouver Tableau User Group, Snowflake Canada User Group, Amazon Tableau User Group and Amazon BI Tech Talks,
			#University of Victoria lecturer - Cloud Computing,
			#Author of 6 books about Analytics,
			#Develops consaltyng in North America - rockyourdata.cloud,
			#Speaker on conferences and meetups in Russia and North America,
			#Courses are being developed on the West Coast`,
		CreatedOn: time.Now().UTC().String(),
	},
	{
		ID:               "www",
		CourseID:         "2",
		Email:            "bottomless@datalearn.biz",
		Twitter:          "twitter.com",
		Facebook:         "facebook.com",
		Instagram:        "instagram.com",
		Location:         "Moscow, Russia",
		Fullname:         "Roman Ponomarev",
		Bio:              "Lorem ipsum dolor sit amet consectetur adipisicing elit. Nostrum numquam quaerat magnam unde necessitatibus cum, nesciunt error earum optio molestias, laborum aliquam illum sint quo architecto minus magni culpa? Aspernatur explicabo mollitia, quae atque quidem error aperiam, perferendis nostrum recusandae, culpa molestiae. Ab assumenda saepe amet, optio, maxime incidunt commodi corporis totam recusandae provident corrupti! Enim sit suscipit optio voluptates recusandae quos at. Perspiciatis ut inventore corporis nostrum sit quas aliquam omnis natus neque quidem, delectus qui commodi dolorem quae dolore cumque, quisquam pariatur unde. Cum quam, autem provident deserunt explicabo unde earum blanditiis tenetur quis aspernatur quisquam. Aliquam, ducimus!",
		ShortDescription: "Co-founder and admin of the project",
		Speciality:       "Small data",
		Features: `#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?
					#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?,
			#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?,
			#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?,
			#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?,
			#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?`,
		CreatedOn: time.Now().UTC().String(),
	},
	{
		ID:               "eee",
		CourseID:         "3",
		Email:            "fullydressed@datalearn.biz",
		Twitter:          "twitter.com",
		Facebook:         "facebook.com",
		Instagram:        "instagram.com",
		Location:         "Moscow, Russia",
		Fullname:         "Sergei Isaev",
		Bio:              "Lorem ipsum dolor sit amet consectetur adipisicing elit. Nostrum numquam quaerat magnam unde necessitatibus cum, nesciunt error earum optio molestias, laborum aliquam illum sint quo architecto minus magni culpa? Aspernatur explicabo mollitia, quae atque quidem error aperiam, perferendis nostrum recusandae, culpa molestiae. Ab assumenda saepe amet, optio, maxime incidunt commodi corporis totam recusandae provident corrupti! Enim sit suscipit optio voluptates recusandae quos at. Perspiciatis ut inventore corporis nostrum sit quas aliquam omnis natus neque quidem, delectus qui commodi dolorem quae dolore cumque, quisquam pariatur unde. Cum quam, autem provident deserunt explicabo unde earum blanditiis tenetur quis aspernatur quisquam. Aliquam, ducimus!",
		ShortDescription: "Just a random guy",
		Speciality:       "Doing nothing",
		Features:         "#Lorem ipsum dolor sit amet. #Lorem ipsum dolor sit amet. #Lorem ipsum dolor sit amet.",
		CreatedOn:        time.Now().UTC().String(),
	},
}
