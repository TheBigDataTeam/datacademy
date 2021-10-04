import React from 'react'
import { useHistory } from 'react-router-dom'
import { Button } from 'components/ui'
import styles from './Buttons.module.css'
import { fetchUserLogout } from 'redux/user_auth/userAuthActions'
import { useDispatch } from 'react-redux'

export const Buttons: React.FunctionComponent = (): JSX.Element => {

    const dispatch = useDispatch()
    const history = useHistory()

    const handleLogoutClick = () => {
        try {
            dispatch(fetchUserLogout())
            history.push("/")
        } catch (error) {
            console.log(error)
        }
    }

    const handleCancelClick = () => {
        history.goBack()
    }

    return (
        <div className={styles.root}>
            <Button onClick={handleLogoutClick}>Logout</Button>
            <Button onClick={handleCancelClick}>Cancel</Button>
        </div>
    )
}
