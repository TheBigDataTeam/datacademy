import { UserAuthActionTypes } from './userAuthActionTypes'
import { BASE_URL } from 'constants/common'
import { User } from 'models/User'
import axios from 'axios'
import { Dispatch } from 'redux'

/* TODO Error type */
type Error = string

export type UserAuthRequest = {
    type: UserAuthActionTypes.FETCH_USER_AUTH_REQUEST,
}

export type UserAuthSuccess = {
    type: UserAuthActionTypes.FETCH_USER_AUTH_SUCCESS,
    payload: User
}

export type UserAuthFailure = {
    type: UserAuthActionTypes.FETCH_USER_AUTH_FAILURE,
    payload: Error
}

export type UserAuthActions = UserAuthRequest | UserAuthSuccess | UserAuthFailure

export const fetchUserAuthRequest = (): UserAuthRequest => {
    return {
        type: UserAuthActionTypes.FETCH_USER_AUTH_REQUEST
    }
}

export const fetchUserAuthSuccess = (user: User): UserAuthSuccess => {
    return {
        type: UserAuthActionTypes.FETCH_USER_AUTH_SUCCESS,
        payload: user
    }
}

export const fetchUserAuthFailure = (error: Error): UserAuthFailure => {
    return {
        type: UserAuthActionTypes.FETCH_USER_AUTH_FAILURE,
        payload: error
    }
}

export const fetchUserAuth = () => {
    return (dispatch: Dispatch<UserAuthActions>): void => {
        dispatch(fetchUserAuthRequest())
        axios.get<User>(BASE_URL + `api/auth/user`, {withCredentials: true}).then(response => {
            const user = response.data
            dispatch(fetchUserAuthSuccess(user))
        }).catch(error => {
            const errMsg = error.message
            dispatch(fetchUserAuthFailure(errMsg))
        })
    }
}