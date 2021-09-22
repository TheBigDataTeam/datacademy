/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
import { FETCH_USER_REQUEST, FETCH_USER_SUCCESS, FETCH_USER_FAILURE } from './userAuthTypes'
import { BASE_URL } from 'constants/common'
import { User } from 'models'
import axios from 'axios'

type Error = any

export type UserRequest = {
    type: typeof FETCH_USER_REQUEST,
}

export type UserSuccess = {
    type: typeof FETCH_USER_SUCCESS,
    payload: User
}

export type UserFailure = {
    type: typeof FETCH_USER_FAILURE,
    payload: Error
}

export const fetchUserRequest = (): UserRequest => {
    return {
        type: FETCH_USER_REQUEST
    }
}

export const fetchUserSuccess = (user: User): UserSuccess => {
    return {
        type: FETCH_USER_SUCCESS,
        payload: user
    }
}

export const fetchUserFailure = (error: Error): UserFailure => {
    return {
        type: FETCH_USER_FAILURE,
        payload: error
    }
}

export const fetchUser = (userID: string) => {
    return (dispatch: any): any => {
        dispatch(fetchUserRequest)
        axios.get(BASE_URL + `api/users?user=${userID}`).then(response => {
            console.log(response.data)
            const user = response.data
            dispatch(fetchUserSuccess(user))
        }).catch(error => {
            const errMsg = error.message
            dispatch(fetchUserFailure(errMsg))
        })
    }
}