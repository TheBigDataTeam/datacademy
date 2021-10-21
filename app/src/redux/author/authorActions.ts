import { AuthorActionTypes } from './authorActionTypes'
import { Author } from 'models/Author'

/* TODO Error type */
type Error = string

export type ActionChangePayload = {
    name: string
    value: string
}

export type ActionChange = {
    type: AuthorActionTypes.FETCH_ACTION_CHANGE
    payload: ActionChangePayload
}

export type AddAuthorRequest = {
    type: AuthorActionTypes.FETCH_ADD_AUTHOR_REQUEST
}

export type AddAuthorSuccess = {
    type: AuthorActionTypes.FETCH_ADD_AUTHOR_SUCCESS
    payload: Author
}

export type AddAuthorFailure = {
    type: AuthorActionTypes.FETCH_ADD_AUTHOR_FAILURE
    payload: Error
}

export type AuthorAction = ActionChange | AddAuthorRequest | AddAuthorSuccess | AddAuthorFailure

export const fetchAuthorActionChange = (name: string, value: string): ActionChange => {
    return {
        type: AuthorActionTypes.FETCH_ACTION_CHANGE,
        payload: {
            name,
            value
        }
    }
}

export const fetchAddAuthorRequest = (): AddAuthorRequest => {
    return {
        type: AuthorActionTypes.FETCH_ADD_AUTHOR_REQUEST
    }
}

export const fetchAddAuthorSuccess = (author: Author): AddAuthorSuccess => {
    return {
        type: AuthorActionTypes.FETCH_ADD_AUTHOR_SUCCESS,
        payload: author
    }
}

export const fetchAddAuthorFailure = (error: Error): AddAuthorFailure => {
    return {
        type: AuthorActionTypes.FETCH_ADD_AUTHOR_FAILURE,
        payload: error
    }
}