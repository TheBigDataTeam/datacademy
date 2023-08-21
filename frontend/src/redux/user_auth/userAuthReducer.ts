import { UserAuthActionTypes } from './userAuthActionTypes'
import { UserAuthActions } from './userAuthActions'
import { User } from 'models/User'

export type initialStateType = {
    isLoading: boolean,
    user: User | null,
    error: string,
    isLoaded: boolean
}

export const initialState: initialStateType = {
    isLoading: false,
    user: null,
    error: '',
    isLoaded: false
}

export const userAuthReducer = (state = initialState, action: UserAuthActions): initialStateType => {
    switch(action.type) {
        case UserAuthActionTypes.FETCH_USER_LOGIN_REQUEST:
            return {
                ...state,
                isLoading: true
            }
        case UserAuthActionTypes.FETCH_USER_LOGIN_SUCCESS: {
            return {
                isLoading: false,
                user: action.payload,
                error: '',
                isLoaded: true
            }
        }
        case UserAuthActionTypes.FETCH_USER_LOGIN_FAILURE: {
            return {
                ...state,
                isLoading: false,
                user: null,
                error: action.payload,
                isLoaded: false
            }
        }
        case UserAuthActionTypes.FETCH_USER_LOGOUT_REQUEST: {
            return {
                ...state,
                isLoading: true
            }
        }
        case UserAuthActionTypes.FETCH_USER_LOGOUT_SUCCESS: {
            return {
                isLoading: false,
                user: null,
                error: '',
                isLoaded: false
            }
        }
        case UserAuthActionTypes.FETCH_USER_LOGOUT_FAILURE: {
            return {
                ...state,
                isLoading: false,
                error: action.payload,
            }
        }
        default: return state
    }
}