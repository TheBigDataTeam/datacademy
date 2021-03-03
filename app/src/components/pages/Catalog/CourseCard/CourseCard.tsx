import React from 'react';
import styles from './CourseCard.module.css';
import { Button} from "components/ui";

interface Props {
    title: string
    image: string
    description: string
    tech_stack: string
    author: string
}

export const CourseCard: React.FunctionComponent<Props> = ({ title, image, description,tech_stack, author }): JSX.Element => {
    return (
        <div className={styles.root}>
            <h2>{title}</h2>
            <p>{description}</p>
            <p>{tech_stack}</p>
            <p>{author}</p>
            <Button design='primary' type='button' />
        </div>

    )
}