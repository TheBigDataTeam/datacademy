/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
import { FETCH_USER_REQUEST, FETCH_USER_SUCCESS, FETCH_USER_FAILURE } from './userAuthTypes'
import { User } from 'models'
//import { UserRequest, UserSuccess, UserFailure } from './userAuthActions'

export type initialStateType = {
    isLoading: boolean,
    user: User | null,
    error: string
}

export const initialState: initialStateType = {
    isLoading: false,
    user: null,
    error: '',
}

export const userAuthReducer = (state = initialState, action: any): initialStateType => {
    switch(action.type) {
        case FETCH_USER_REQUEST:
            return {
                ...state,
                isLoading: true
            }
        case FETCH_USER_SUCCESS: {
            return {
                isLoading: false,
                user: action.payload,
                error: ''
            }
        }
        case FETCH_USER_FAILURE: {
            return {
                ...state,
                isLoading: false,
                error: action.payload
            }
        }
        default: return state
    }
}