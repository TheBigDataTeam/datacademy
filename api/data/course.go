package data

import (
	"fmt"
	"time"
)

// Course defines the structure for an API
type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"`
	Author      string `json:"author" validate:"required,author"`
	TechStack   string `json:"techstack"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
}

// Courses is a slice of Course(s)
type Courses []*Course

// ErrorCourseNotFound is an error raised when a product can not be found
var ErrorCourseNotFound = fmt.Errorf("Course not found")

// GetCourses returns a list of courses
func GetCourses() Courses {
	return courseList
}

// GetCourseByID returns a single course which matches the id from the DB
// if a course is not found this function returns a CourseNotFound error
func GetCourseByID(id int) (*Course, error) {
	index := findIndexByCourseID(id)
	if id == -1 {
		return nil, ErrorCourseNotFound
	}
	return courseList[index], nil
}

// UpdateCourse replaces a course in the DB with the given item.
// If a course with the given id does not exist in th DB
// this function returns CourseNotFound error
/* TODO: currently if some filds are comming empty - it makes corresponding filds in data set empty as well - needs to be fixed */
func UpdateCourse(c Course) error {
	positionOfFoundCourse := findIndexByCourseID(c.ID)
	if positionOfFoundCourse == -1 {
		return ErrorCourseNotFound
	}
	courseList[positionOfFoundCourse] = &c
	return nil
}

// AddCourse adds new course to the DB
func AddCourse(c Course) {
	c.ID = getNextID()
	courseList = append(courseList, &c)
}

// DeleteCourse deletes a course from the DB
func DeleteCourse(id int) error {
	index := findIndexByCourseID(id)
	if index == -1 {
		return ErrorCourseNotFound
	}
	courseList = append(courseList[:index], courseList[index+1])
	return nil
}

/* TODO: make this function available for all handlers */

// getNextId generates a new id for a product being inserted into db
func getNextID() int {
	lastItemInDB := courseList[len(courseList)-1]
	return lastItemInDB.ID + 1
}

// findIndexByCourseID finds the index of a course in the DB
// returns -1 when no course isfound
func findIndexByCourseID(id int) int {
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
		ID:          1,
		Title:       "Big Data for Dummies",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
		Image:       "big-data.png",
		Author:      "Dima Anoshin",
		TechStack:   "SQL",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Title:       "Big Data for Dummies",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
		Image:       "big-data.png",
		Author:      "Dima Anoshin",
		TechStack:   "SQL",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
