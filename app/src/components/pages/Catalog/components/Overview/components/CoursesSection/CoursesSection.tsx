import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { CourseCard} from "./components";
import { Grid } from 'components/ui';
import { Course } from 'models';
import axios, { AxiosResponse } from 'axios';

export const CoursesSection: React.FunctionComponent = (): JSX.Element => {

    const [courses, setCourses] = useState<Array<Course>>(null);

    useEffect(() => {
        const fetchCourses = async (): Promise<void> => {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const results: AxiosResponse<any> = await axios.get("http://localhost:3100/courses");
            setCourses(results.data);
        };
        fetchCourses();
    }, []);

    return (
        <Grid.Row>
            {courses ? courses.map(course => (
                <Grid.Col key={course.id} cols={12} colsSM={6} colsMD={4} marginBottom='xl'>
                    <Link to={`/courses/${course.id}`}>
                        <CourseCard
                            title={course.title}
                            description={course.description}
                            author={course.author}
                        />
                    </Link>
                </Grid.Col>
            )) : <h2>Loading ...</h2>}
        </Grid.Row>
    )
}