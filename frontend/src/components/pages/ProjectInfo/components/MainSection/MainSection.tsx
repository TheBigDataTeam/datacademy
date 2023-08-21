import React from 'react'
import { Paragraph, Grid } from 'components/ui'
import styles from './MainSection.module.css'
import { listOfCitations } from './data/data'

export const MainSection: React.FunctionComponent = (): JSX.Element => {
    return (
        <>
            <Grid.Row marginBottom="xxl">
                <Grid.Col cols={6} colsMD={6} colsSM={12}>
                    <span className={styles.title}>Our Culture and Values</span>
                </Grid.Col>
                <Grid.Col cols={6} colsMD={6} colsSM={12}>
                    <div className={styles.cover}></div>
                </Grid.Col>
            </Grid.Row>
            {listOfCitations.map((citation) => (
                <Grid.Row key={citation.id}>
                    <Grid.Col cols={12} colsMD={12} colsSM={12}>
                        <Paragraph size="xl">
                            {citation.text}
                        </Paragraph>
                        <Paragraph align="center" size="l" marginBottom="xl">
                            © {citation.author}
                        </Paragraph>
                        <div className={styles.bottom_border}></div>
                    </Grid.Col> 
                </Grid.Row>
            ))}
        </>
    )
} 