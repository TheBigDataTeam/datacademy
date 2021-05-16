import React from 'react';
import styles from './BeneficiarsSection.module.css';
import { Grid, Header, Paragraph } from 'components/ui';
import { v4 as uuidv4 } from 'uuid';

interface Props {
    beneficiars: Array<string>
}

export const BeneficiarsSection: React.FunctionComponent<Props> = ({ beneficiars }): JSX.Element => {
    return (
        <Grid.Row>
            <Grid.Col>
                <Header>Who will benefit from taking this course</Header>
                <ul className={styles.list}>
                    {beneficiars.map((person) => (
                        <li key={uuidv4()} className={styles.list_item}>
                            <Paragraph size="ml">{person}</Paragraph>
                        </li>
                    ))}
                </ul>
            </Grid.Col>
        </Grid.Row>
    )
}
