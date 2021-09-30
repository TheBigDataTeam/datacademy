import { UserAuthActionTypes } from './userAuthActionTypes'
import { UserAuthAction } from './userAuthActions'

export type initialStateType = {
    isLoading: boolean,
    userId: string,
    error: string,
    isLoaded: boolean
}

export const initialState: initialStateType = {
    isLoading: false,
    userId: '',
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
                userId: action.payload,
                error: '',
                isLoaded: true
            }
        }
        case UserAuthActionTypes.FETCH_USER_AUTH_FAILURE: {
            return {
                ...state,
                isLoading: false,
                userId: '',
                error: action.payload,
                isLoaded: false
            }
        }
        default: return state
    }
}