package authors

import (
	"errors"
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Repo represents a data base
type Repo struct {
	mongo      *mgo.Session
	collection *mgo.Collection
}

// NewRepo returns an instance of a Repo
func NewRepo(mongo *mgo.Session, collection *mgo.Collection) *Repo {
	return &Repo{
		mongo:      mongo,
		collection: collection,
	}
}

var (
	// ErrNoRecord is returned when no record in database is found
	ErrNoRecord = errors.New("No author record found")
	// ErrorAuthorAlreadyExists is returned when one tries to add user which already exists
	ErrorAuthorAlreadyExists = errors.New("Author already exists")
	// ErrorBadRequest is returned when one tries to create a user with a wrong data
	ErrorBadRequest = errors.New("Error inserting info into db")
)

// AddAuthor adds a new author in the DB
func (r *Repo) AddAuthor(a Author) error {
	a.ID = bson.NewObjectId()
	err := r.collection.Insert(a)
	if err != nil {
		return ErrorBadRequest
	}
	return nil
	/* to be continued */
}

// GetAuthors returns list of authors
func (r *Repo) GetAuthors() ([]*Author, error) {

	return authorsList, nil
}

// GetAuthorByID returns a single author which matches the id from the DB
func (r *Repo) GetAuthorByID(id string) (*Author, error) {

	return nil, nil
}

// UpdateAuthor replaces an author in the DB with the given item.
func (r *Repo) UpdateAuthor(a Author) error {

	return nil
}

// DeleteAuthor deletes the author from the DB
func (r *Repo) DeleteAuthor(id string) error {

	return nil
}

// courseList is a hard coded list of courses / test data source
var authorsList = []*Author{
	{
		ID:        "qqqqq",
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
		Features: []string{"#Data Engineer in Amazon, Alexa AI",
			"#10+ years of experience in Analytics (Russia, Europe, Canada and USA)",
			"#Organizer of Vancouver Tableau User Group, Snowflake Canada User Group, Amazon Tableau User Group and Amazon BI Tech Talks",
			"#University of Victoria lecturer - Cloud Computing",
			"#Author of 6 books about Analytics",
			"#Develops consaltyng in North America - rockyourdata.cloud",
			"#Speaker on conferences and meetups in Russia and North America",
			"#Courses are being developed on the West Coast"},
		CreatedOn: time.Now().UTC().String(),
	},
	{
		ID:               "www",
		Email:            "bottomless@datalearn.biz",
		Twitter:          "twitter.com",
		Facebook:         "facebook.com",
		Instagram:        "instagram.com",
		Location:         "Moscow, Russia",
		Fullname:         "Roman Ponomarev",
		Bio:              "Lorem ipsum dolor sit amet consectetur adipisicing elit. Nostrum numquam quaerat magnam unde necessitatibus cum, nesciunt error earum optio molestias, laborum aliquam illum sint quo architecto minus magni culpa? Aspernatur explicabo mollitia, quae atque quidem error aperiam, perferendis nostrum recusandae, culpa molestiae. Ab assumenda saepe amet, optio, maxime incidunt commodi corporis totam recusandae provident corrupti! Enim sit suscipit optio voluptates recusandae quos at. Perspiciatis ut inventore corporis nostrum sit quas aliquam omnis natus neque quidem, delectus qui commodi dolorem quae dolore cumque, quisquam pariatur unde. Cum quam, autem provident deserunt explicabo unde earum blanditiis tenetur quis aspernatur quisquam. Aliquam, ducimus!",
		ShortDescription: "Co-founder and admin of the project",
		Features: []string{"#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?",
			"#Lorem ipsum dolor sit amet consectetur adipisicing elit. Non, ipsum?}"},
		CreatedOn: time.Now().UTC().String(),
	},
	{
		ID:               "eee",
		Email:            "fullydressed@datalearn.biz",
		Twitter:          "twitter.com",
		Facebook:         "facebook.com",
		Instagram:        "instagram.com",
		Location:         "Moscow, Russia",
		Fullname:         "Sergei Isaev",
		Bio:              "Lorem ipsum dolor sit amet consectetur adipisicing elit. Nostrum numquam quaerat magnam unde necessitatibus cum, nesciunt error earum optio molestias, laborum aliquam illum sint quo architecto minus magni culpa? Aspernatur explicabo mollitia, quae atque quidem error aperiam, perferendis nostrum recusandae, culpa molestiae. Ab assumenda saepe amet, optio, maxime incidunt commodi corporis totam recusandae provident corrupti! Enim sit suscipit optio voluptates recusandae quos at. Perspiciatis ut inventore corporis nostrum sit quas aliquam omnis natus neque quidem, delectus qui commodi dolorem quae dolore cumque, quisquam pariatur unde. Cum quam, autem provident deserunt explicabo unde earum blanditiis tenetur quis aspernatur quisquam. Aliquam, ducimus!",
		ShortDescription: "Just a random guy",
		Features:         []string{"#Lorem ipsum dolor sit amet.", "#Lorem ipsum dolor sit amet. #Lorem ipsum dolor sit amet."},
		CreatedOn:        time.Now().UTC().String(),
	},
}
