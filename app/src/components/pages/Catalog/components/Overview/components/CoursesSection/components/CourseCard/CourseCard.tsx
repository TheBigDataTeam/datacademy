import React from 'react';
import styles from './CourseCard.module.css';
import { Paragraph } from 'components/ui';

interface Props {
    title: string
    description: string
    tech_stack: string
    author: string
}

export const CourseCard: React.FunctionComponent<Props> = ({ title, description,tech_stack, author }): JSX.Element => {
    return (
        <div className={styles.root}>
            <Paragraph align="center" size="xl">{title}</Paragraph>
            <div className={styles.img}>picture</div>
            <Paragraph>{description}</Paragraph>
            <Paragraph>Tech stack: {tech_stack}</Paragraph>
            <Paragraph>Author: {author}</Paragraph>
        </div>

    )
}