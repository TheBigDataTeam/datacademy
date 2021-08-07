package courses

import (
	"database/sql"
	"fmt"

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

// ErrorCourseNotFound is an error raised when a course can not be found
var ErrorCourseNotFound = fmt.Errorf("Course not found")

/* TODO: all functions are not concurrently safe - add mutexes or something! */

// GetCourses returns a list of courses
func (r *Repo) GetCourses() ([]*Course, error) {
	courses := make([]*Course, 0, 10)
	rows, err := r.db.Query("SELECT id, title, description, theme, author, author_id, tech_stack, syllabus, duration, beneficiars, difficulty, created_on FROM courses")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := &Course{}
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Theme, &item.Author, &item.AuthorID, &item.TechStack,
			&item.Syllabus, &item.Duration, &item.Beneficiars, &item.Difficulty, &item.CreatedOn)
		if err != nil {
			return nil, err
		}
		courses = append(courses, item)
	}
	return courses, nil
}

// GetCourseByID returns a single course which matches the id from the DB
// if a course is not found this function returns a CourseNotFound error
func (r *Repo) GetCourseByID(id string) (*Course, error) {
	index := findIndexByCourseID(id)
	if index == -1 {
		return nil, ErrorCourseNotFound
	}
	return courseList[index], nil
}

// UpdateCourse replaces a course in the DB with the given item.
// If a course with the given id does not exist in th DB
// this function returns CourseNotFound error
/* TODO: currently if some filds are comming empty - it makes corresponding fields in data set empty as well - needs to be fixed */
func (r *Repo) UpdateCourse(c Course) error {
	index := findIndexByCourseID(c.ID)
	if index == -1 {
		return ErrorCourseNotFound
	}
	courseList[index] = &c
	return nil
}

// AddCourse adds new course to the DB
func (r *Repo) AddCourse(c Course) {
	c.ID = getNextCourseID() // TODO handle possible mistakes
	courseList = append(courseList, &c)
}

// DeleteCourse deletes a course from the DB
func (r *Repo) DeleteCourse(id string) error {
	index := findIndexByCourseID(id)
	if index == -1 {
		return ErrorCourseNotFound
	}
	courseList = append(courseList[:index], courseList[index+1])
	return nil
}

/* TODO: make this function available for all handlers */

// getNextId generates a new id for a product being inserted into db
func getNextCourseID() string {
	newID := util.RandString()
	return newID
}

// findIndexByCourseID finds the index of a course in the DB
// returns -1 when no course is found
func findIndexByCourseID(id string) int {
	for i, c := range courseList {
		if c.ID == id {
			return i
		}
	}
	return -1
}

// courseList is a hard coded list of courses / test data source
var courseList = []*Course{
	{
		ID:          "aaaaaaaaaa",
		Title:       "Big Data for Dummies",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
		Author:      "Dmitry Anoshin",
		/* AuthorID:    "qqqqq",
		TechStack: []string{"Excel", "SQL: Postgres/MySQL", "Amazon Redshift", "ETL Pentaho DI",
			"BigData Elastic Map Reduce (Hadoop), Hive, Presto, Athena, Spectrum", "BI Tableau"},
		Syllabus: []string{"Module 1: Roles of analytics and data engineer in an organization",
			"Module 2: Databases and SQL",
			"Module 3: Data visualization, dashboards and reporting - Business Intelligence",
			"Module 4: Data integration and data pipelines",
			"Module 4: Cloud Computing",
			"Module 6: Cloud data storage",
			"Module 7: Introduction to Apache Spark",
			"Module 8: Big Data problem solving with Hadoop and Spark",
			"Module 9: Introduction to Data Lake and how to make it with AWS",
			"Module 10: Data streaming task completion",
			"Module 11: Machine Learning tasks from data engineer's point of view",
			"Module 12: Data engineer's best practices"},
		Beneficiars: []string{"Analysts: Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
			"Market managers: Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
			"Engineers: Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
			"Entrepreneurs: Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
			"Financiers: Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
			"And of course - Novices"},
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(), */
	},
	{
		ID:          "bbbbbbbbb",
		Title:       "Small Data for Dummies",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
		Author:      "Dima Anoshin",
		/* AuthorID:    "qqqqq",
		TechStack: []string{"Excel", "SQL: Postgres/MySQL", "Amazon Redshift", "ETL Pentaho DI",
			"BigData Elastic Map Reduce (Hadoop), Hive, Presto, Athena, Spectrum", "BI Tableau"},
		Syllabus: []string{"Module 1: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 2: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 3: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 4: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 4: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 6: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 7: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 8: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 9: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 10: Lorem ipsum dolor sit amet consectetur adipisicing elit."},
		Beneficiars: []string{"Analysts", "Marketologs", "Engineers", "Entrepreneurs", "Newcomers"},
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(), */
	},
	{
		ID:          "cccccccc",
		Title:       "Very Big Data for Dummies",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
		Author:      "Dima Anoshin",
		/* AuthorID:    "qqqqq",
		TechStack: []string{"Excel", "SQL: Postgres/MySQL", "Amazon Redshift", "ETL Pentaho DI",
			"BigData Elastic Map Reduce (Hadoop), Hive, Presto, Athena, Spectrum", "BI Tableau"},
		Syllabus: []string{"Module 1: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 2: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 3: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 4: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 4: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 6: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 7: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 8: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 9: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 10: Lorem ipsum dolor sit amet consectetur adipisicing elit."},
		Beneficiars: []string{"Marketologs", "Engineers", "Entrepreneurs", "Newcomers"},
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(), */
	},
	{
		ID:          "dddddddddd",
		Title:       "So So Data for Dummies",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
		Author:      "Dima Anoshin",
		/* AuthorID:    "qqqqq",
		TechStack: []string{"Excel", "SQL: Postgres/MySQL", "Amazon Redshift", "ETL Pentaho DI",
			"BigData Elastic Map Reduce (Hadoop), Hive, Presto, Athena, Spectrum", "BI Tableau"},
		Syllabus: []string{"Module 1: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 2: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 3: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 4: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 4: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 6: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 7: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 8: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 9: Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			"Module 10: Lorem ipsum dolor sit amet consectetur adipisicing elit."},
		Beneficiars: []string{"Marketologs", "Engineers", "Entrepreneurs", "Newcomers"},
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(), */
	},
}
