import React from 'react';
import { Paragraph } from 'components/ui';
//import { v4 as uuidv4 } from 'uuid';
import { Grid, Header } from 'components/ui';
//import styles from './TechStackSection.module.css';

interface Props {
    techstack: string
}

export const TechStackSection: React.FunctionComponent<Props> = ({ techstack }): JSX.Element => {
    return (
        <Grid.Row>
            <Grid.Col>
                <Header>Here is what we will be using:</Header>
                {/* <ul className={styles.list}>
                    {techstack.map((stack) => (
                        <li key={uuidv4()} className={styles.list_item}><Paragraph size="ml">{stack}</Paragraph></li>
                    ))}
                </ul> */}
                <Paragraph>{techstack}</Paragraph>
            </Grid.Col>
        </Grid.Row>
    )
}