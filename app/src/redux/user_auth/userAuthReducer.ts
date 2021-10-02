import { UserAuthActionTypes } from './userAuthActionTypes'
import { UserAuthAction } from './userAuthActions'
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

export const userAuthReducer = (state = initialState, action: UserAuthAction): initialStateType => {
    switch(action.type) {
        case UserAuthActionTypes.FETCH_USER_AUTH_REQUEST:
            return {
                ...state,
                isLoading: true
            }
        case UserAuthActionTypes.FETCH_USER_AUTH_SUCCESS: {
            return {
                isLoading: false,
                user: action.payload,
                error: '',
                isLoaded: true
            }
        }
        case UserAuthActionTypes.FETCH_USER_AUTH_FAILURE: {
            return {
                ...state,
                isLoading: false,
                user: null,
                error: action.payload,
                isLoaded: false
            }
        }
        default: return state
    }
}