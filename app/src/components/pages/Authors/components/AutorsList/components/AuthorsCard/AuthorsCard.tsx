import React from 'react';
import { Link } from 'react-router-dom';
import { Paragraph } from 'components/ui';
import styles from './AuthorsCard.module.css';

interface Props {
    fullName: string,
    firstName: string,
    location: string,
    photo: string,
    description: string,
    authorOf: string[],
}

export const AuthorsCard: React.FunctionComponent<Props> = (props): JSX.Element => {
    return (
        <div className={styles.root}>
            <div className={styles.img}/>
            <div className={styles.text_block}>
                <Paragraph size="l">{props.fullName}</Paragraph>
                <Paragraph size="l">{props.location}</Paragraph>
                <Paragraph size="l">{props.description}</Paragraph>
                <Paragraph size="l" marginBottom="l">Author of: {props.authorOf}</Paragraph>
                <Paragraph size="l">More Info about {props.firstName} <Link to="/authors">here</Link></Paragraph>
            </div>
        </div>
    )
}