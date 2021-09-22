import { FETCH_USER_REQUEST, FETCH_USER_SUCCESS, FETCH_USER_FAILURE } from './userAuthTypes'
import { BASE_URL } from 'constants/common'
import axios from 'axios'

type User = string
type Error = string

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

export const fetchUser = () => {
    return (dispatch: any): any => {
        dispatch(fetchUserRequest)
        axios.get(BASE_URL)
    }
}