import React from 'react'
import { Paragraph, Grid } from 'components/ui'
import styles from './BioSection.module.css'

interface Props {
    bio: string
}

export const BioSection: React.FunctionComponent<Props> = ({ bio }): JSX.Element => {
    return (
        <Grid.Row>
            <Grid.Col>
                <div className={styles.wrapper}>
                    <div className={styles.img}></div>
                    <Paragraph size="ml" marginTop="none">{bio}</Paragraph>
                </div>
            </Grid.Col>
        </Grid.Row>
    )
}