import React from 'react'
import { Buttons } from './components'
import { Paragraph, Heading } from 'components/ui'
import styles from './Form.module.css'

export const Form:React.FunctionComponent = (): JSX.Element => {
    return (
        <div className={styles.root}>
            <Heading size="xxl" weight="normal" color="inverted">Datalearn</Heading>
            <Paragraph size="l" color="inverted">West Coast Analytics from Dmitriy Anoshin</Paragraph>
            <Buttons />
        </div>
    )
}
