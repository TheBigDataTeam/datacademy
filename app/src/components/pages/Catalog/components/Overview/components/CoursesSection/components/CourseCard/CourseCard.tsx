import React from 'react';
import styles from './CourseCard.module.css';
import { Paragraph } from 'components/ui';

interface Props {
    title: string
    description: string
    author: string
}

export const CourseCard: React.FunctionComponent<Props> = ({ title, description, author }): JSX.Element => {
    return (
        <div className={styles.root}>
            <Paragraph align="center" size="xl">{title}</Paragraph>
            <div className={styles.img}>picture</div>
            <Paragraph marginTop="xl">{description}</Paragraph>
            <Paragraph marginBottom="xl">Author: {author}</Paragraph>
        </div>

    )
}