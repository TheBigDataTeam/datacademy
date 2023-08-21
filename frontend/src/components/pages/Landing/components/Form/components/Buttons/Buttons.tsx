import React from 'react'
import { Link } from 'react-router-dom'
import { Button, Text } from 'components/ui'
import styles from './Buttons.module.css'

export const Buttons:React.FunctionComponent = (): JSX.Element => {
    return (
        <div className={styles.root}>
            <Link to="/auth/signup" className={styles.root_link}>
                <Button design='secondary' size="m">
                    <Text size="l">Sign Up</Text>
                </Button>
            </Link>
            <Link to="/auth/login" className={styles.root_link}>
                <Button design='primary' size="m">
                    <Text size="l" color="inverted">Sign In</Text>
                </Button>
            </Link>
        </div>
    )
}
