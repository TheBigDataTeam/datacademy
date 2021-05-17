import React from 'react';
import { Link } from 'react-router-dom';
import { Paragraph } from 'components/ui';
import styles from './AuthorsCard.module.css';
import { Author } from 'models';

interface Props {
    author: Author
}

export const AuthorsCard: React.FunctionComponent<Props> = ({ author }): JSX.Element => {
    return (
        <div className={styles.root}>
            <div className={styles.img}/>
            <div className={styles.text_block}>
                <Paragraph size="l">{author.fullname}</Paragraph>
                <Paragraph size="l">{author.location}</Paragraph>
                <Paragraph size="l">{author.shortdescription}</Paragraph>
                <Paragraph size="l" marginBottom="l">Author of: {/* {props.authorOf} */}</Paragraph>
                <Paragraph size="l">More Info about {author.fullname} <Link to={`/authors/${author.id}`}>here</Link></Paragraph>
            </div>
        </div>
    )
}