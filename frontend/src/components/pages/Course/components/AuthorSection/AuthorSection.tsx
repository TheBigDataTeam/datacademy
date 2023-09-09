import React from 'react'
//import { v4 as uuidv4 } from 'uuid'
import { Grid, Paragraph, Heading } from 'components/ui'
import styles from './AuthorSection.module.css'
import { Author } from 'models/Author'

interface Props {
    author: Author
}

export const AuthorSection: React.FunctionComponent<Props> = ({ author }): JSX.Element => {
    return (
        <>
            <Grid.Row marginBottom="none">
                <Grid.Col marginBottom="none">
                    <Heading>Your teacher: {author?.fullname}</Heading>
                </Grid.Col>
            </Grid.Row>
            <Grid.Row marginTop="none">
                <Grid.Col cols={12} colsSM={12} colsMD={12}>
                    <div className={styles.wrapper}>
                        {/* <ul className={styles.list}>
                            {author?.features.map((feature) => (
                                <li key={uuidv4()} className={styles.list_item}>
                                    <Paragraph size="ml">{feature}</Paragraph>
                                </li>
                            ))}
                        </ul> */}
                        <Paragraph>{author?.features}</Paragraph>
                        <div className={styles.img}/>
                    </div>
                </Grid.Col>
            </Grid.Row>
        </>
    )
}
