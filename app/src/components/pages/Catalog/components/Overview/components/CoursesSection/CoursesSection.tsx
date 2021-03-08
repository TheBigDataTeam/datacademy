import React from 'react';
import { CourseCard} from "./components";
import { Grid } from 'components/ui';

export const CoursesSection: React.FunctionComponent = (): JSX.Element => {

    const listOfCourseCards = [
        {
            id: 1,
            title: 'Big Data',
            image: '../logo_datalearn.png',
            description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
            tech_stack: 'SQL',
            author: 'Dima Anoshin'
        },
        {
            id: 2,
            title: 'Big Data',
            image: '../logo_datalearn.png',
            description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
            tech_stack: 'SQL',
            author: 'Dima Anoshin'
        },
        {
            id: 3,
            title: 'Big Data',
            image: 'components/pages/catalog/logo_datalearn.png',
            description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
            tech_stack: 'SQL',
            author: 'Dima Anoshin'
        },
        {
            id: 4,
            title: 'Big Data',
            image: '../logo_datalearn.png',
            description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
            tech_stack: 'SQL',
            author: 'Dima Anoshin'
        }
    ];

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