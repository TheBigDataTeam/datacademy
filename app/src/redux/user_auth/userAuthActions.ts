import { UserAuthActionTypes } from './userAuthActionTypes'
import { BASE_URL } from 'constants/common'
import { User } from 'models/User'
import axios from 'axios'
import { Dispatch } from 'redux'

/* TODO Error type */
type Error = string

export type UserLoginRequest = {
    type: UserAuthActionTypes.FETCH_USER_LOGIN_REQUEST,
}

export type UserLoginSuccess = {
    type: UserAuthActionTypes.FETCH_USER_LOGIN_SUCCESS,
    payload: User
}

export type UserLoginFailure = {
    type: UserAuthActionTypes.FETCH_USER_LOGIN_FAILURE,
    payload: Error
}

export type UserLogoutRequest = {
    type: UserAuthActionTypes.FETCH_USER_LOGOUT_REQUEST
}

export type UserLogoutSuccess = {
    type: UserAuthActionTypes.FETCH_USER_LOGOUT_SUCCESS
}

export type UserLogoutFailure = {
    type: UserAuthActionTypes.FETCH_USER_LOGOUT_FAILURE,
    payload: Error
}

export type UserAuthActions = UserLoginRequest | UserLoginSuccess | UserLoginFailure | UserLogoutRequest | UserLogoutSuccess | UserLogoutFailure

export const fetchUserLoginRequest = (): UserLoginRequest => {
    return {
        type: UserAuthActionTypes.FETCH_USER_LOGIN_REQUEST
    }
}

export const fetchUserLoginSuccess = (user: User): UserLoginSuccess => {
    return {
        type: UserAuthActionTypes.FETCH_USER_LOGIN_SUCCESS,
        payload: user
    }
}

export const fetchUserLoginFailure = (error: Error): UserLoginFailure => {
    return {
        type: UserAuthActionTypes.FETCH_USER_LOGIN_FAILURE,
        payload: error
    }
}

export const fetchUserLogoutRequest = (): UserLogoutRequest => {
    return {
        type: UserAuthActionTypes.FETCH_USER_LOGOUT_REQUEST
    }
}

export const fetchUserLogoutSuccess = (): UserLogoutSuccess => {
    return {
        type: UserAuthActionTypes.FETCH_USER_LOGOUT_SUCCESS
    }
}

export const fetchUserLogoutFailure = (error: Error): UserLogoutFailure => {
    return {
        type: UserAuthActionTypes.FETCH_USER_LOGOUT_FAILURE,
        payload: error
    }
}

export const fetchUserLogin = () => {
    return (dispatch: Dispatch<UserAuthActions>): void => {
        dispatch(fetchUserLoginRequest())
        axios.get<User>(BASE_URL + `api/auth/user`, {withCredentials: true})
        .then(response => {
            const user = response.data
            dispatch(fetchUserLoginSuccess(user))})
        .catch(error => {
            const errMsg = error.message
            dispatch(fetchUserLoginFailure(errMsg))
        })
    }
}

export const fetchUserLogout = () => {
    return (dispatch: Dispatch<UserAuthActions>): void => {
        dispatch(fetchUserLoginRequest())
        axios.get(BASE_URL + `api/auth/logout`, {withCredentials: true})
        .then(response => {
            dispatch(fetchUserLogoutSuccess())
            console.log(response)
        })
        .catch(error => {
            const errMsg = error.message
            dispatch(fetchUserLogoutFailure(errMsg))
        })
    }
}