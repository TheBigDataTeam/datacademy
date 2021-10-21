import { AuthorActionTypes } from './authorActionTypes'
import { AuthorAction } from './authorActions'
import { Author } from 'models/Author'

export type initialStateType = {
    isLoading: boolean,
    author: Author,
    error: string,
    isLoaded: boolean
}

export const initialState: initialStateType = {
    isLoading: false,
    author: {
        email: '',
        fullname: '',
        bio: '',
        location: '',
        facebook: '',
        instagram: '',
        twitter: '',
        shortdescription: '',
        features: ''
    },
    error: '',
    isLoaded: false
}

export const authorReducer = (state = initialState, action: AuthorAction): initialStateType => {
    switch(action.type) {
        case AuthorActionTypes.FETCH_ACTION_CHANGE: {
            const { name, value } = action.payload
            return {
                ...state,
                author: {
                  ...state.author,
                  [name]: value,
                }
            }
        }
        case AuthorActionTypes.FETCH_ADD_AUTHOR_REQUEST: {
            return {
                ...state,
                isLoading: true
            }
        }
        case AuthorActionTypes.FETCH_ADD_AUTHOR_SUCCESS: {
            return {
                isLoading: false,
                author: action.payload,
                error: '',
                isLoaded: true
            }
        }
        case AuthorActionTypes.FETCH_ADD_AUTHOR_FAILURE: {
            return {
                ...state,
                isLoading: false,
                author: initialState.author,
                error: action.payload,
            }
        }
        default: {
            return state
        }
    }
}