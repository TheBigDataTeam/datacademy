import React from 'react';
import { CourseCard} from "./components";
import { Grid } from 'components/ui';
import { listOfCourseCards } from 'redux/store';

export const CoursesSection: React.FunctionComponent = (): JSX.Element => {

    return (
        <Grid.Row>
            { listOfCourseCards.map(item => (
                <Grid.Col key={item.id} cols={12} colsSM={6} colsMD={4} marginBottom='xl'>
                <CourseCard
                    title={item.title}
                    image={item.image}
                    description={item.description}
                    tech_stack={item.tech_stack}
                    author={item.author}
                />
                </Grid.Col>
            ))}
        </Grid.Row>
    )
}