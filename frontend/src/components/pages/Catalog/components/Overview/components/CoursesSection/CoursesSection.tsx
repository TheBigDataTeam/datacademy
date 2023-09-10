import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { CourseCard } from './components'
import { Spinner } from 'components/common'
import { Grid } from 'components/ui'
import { Course } from 'models'
import { BASE_URL, COURSES_URL, API_COURSES_URL } from 'constants/urls'
import axios, { AxiosResponse } from 'axios'

export const CoursesSection: React.FunctionComponent = (): JSX.Element => {

    const [courses, setCourses] = useState<Course[] | null>(null)

    useEffect(() => {
        const fetchCourses = async () => {
            const results: AxiosResponse<Course[]> = await axios.get(BASE_URL + API_COURSES_URL)
            setCourses(results.data)
        }
        fetchCourses()
    }, [])

    return (
        <Grid.Row>
            {courses ? courses.map(course => (
                <Grid.Col key={course.id} cols={12} colsSM={6} colsMD={4} marginBottom='xl'>
                    <Link to={`${COURSES_URL}/${course.id}`}>
                        <CourseCard
                            title={course.title}
                            description={course.description}
                            author={course.author}
                        />
                    </Link>
                </Grid.Col>
            )) : <Spinner witdth={100} />}
        </Grid.Row>
    )
}