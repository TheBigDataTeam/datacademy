import React from 'react'
import { useNavigate } from 'react-router-dom'
import { Button } from 'components/ui'
import styles from './Buttons.module.css'
import { fetchUserLogout } from 'redux/user_auth/userAuthActions'
import { useDispatch, useSelector } from 'react-redux'
import { AppStateType } from 'redux/rootReducer'

export const Buttons: React.FunctionComponent = (): JSX.Element => {

    const dispatch = useDispatch()
    const navigate = useNavigate()
    const isLoading = useSelector((state: AppStateType) => state.userAuth.isLoading)

    const handleLogoutClick = () => {
        try {
            dispatch(fetchUserLogout())
            navigate("/")
        } catch (error) {
            console.log(error)
        }
    }

    const handleCancelClick = () => {
        navigate(-1)
    }

    return (
        <div className={styles.root}>
            <Button onClick={handleLogoutClick} design="caution" disabled={isLoading}>Logout</Button>
            <Button onClick={handleCancelClick} design="primary">Cancel</Button>
        </div>
    )
}
