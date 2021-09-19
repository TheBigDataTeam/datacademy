import React from 'react';
import { Paragraph, Heading, Grid } from 'components/ui';
//import styles from './SyllabusSection.module.css';
//import { v4 as uuidv4 } from 'uuid';


interface Props {
    syllabus: string
}

export const SyllabusSection: React.FunctionComponent<Props> = ({ syllabus }): JSX.Element => {
    return (
        <Grid.Row>
            <Grid.Col>
                <Heading>Course content:</Heading>
                {/* <ul className={styles.list}>
                    {syllabus?.map((item) => (
                        <li key={uuidv4()} className={styles.list_item}>
                            <Paragraph size="ml">{item}</Paragraph>
                        </li>
                    ))}
                </ul> */}
                <Paragraph>{syllabus}</Paragraph>
            </Grid.Col>
        </Grid.Row>
    )
}
