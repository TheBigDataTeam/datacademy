import React from 'react'
import { Link } from 'react-router-dom'
import { Button, Text } from 'components/ui'
import { REGISTRATION_URL, LOGIN_URL } from 'constants/urls'
import styles from './Buttons.module.css'

export const Buttons: React.FunctionComponent = (): JSX.Element => {
    return (
        <div className={styles.root}>
            <Link to={REGISTRATION_URL} className={styles.root_link}>
                <Button design='secondary' size="m">
                    <Text size="l">Sign Up</Text>
                </Button>
            </Link>
            <Link to={LOGIN_URL} className={styles.root_link}>
                <Button design='primary' size="m">
                    <Text size="l" color="inverted">Sign In</Text>
                </Button>
            </Link>
        </div>
    )
}
