import { ModuleActionTypes } from './moduleActionTypes'
import { ModuleAction } from './moduleActions'
import { Module } from 'models'

export type initialStateType = {
    isLoading: boolean,
    module: Module,
    error: string,
    isLoaded: boolean
}

export const initialState: initialStateType = {
    isLoading: false,
    module: {
        courseId: '',
        title: '',
        body: '',
        badge: '',
    },
    error: '',
    isLoaded: false
}

export const moduleReducer = (state = initialState, action: ModuleAction): initialStateType => {
    switch(action.type) {
        case ModuleActionTypes.FETCH_ACTION_CHANGE: {
            const { name, value } = action.payload
            return {
                ...state,
                module: {
                  ...state.module,
                  [name]: value,
                }
            }
        }
        case ModuleActionTypes.FETCH_ADD_MODULE_REQUEST: {
            return {
                ...state,
                isLoading: true
            }
        }
        case ModuleActionTypes.FETCH_ADD_MODULE_SUCCESS: {
            return {
                isLoading: false,
                module: action.payload,
                error: '',
                isLoaded: true
            }
        }
        case ModuleActionTypes.FETCH_ADD_MODULE_FAILURE: {
            return {
                ...state,
                isLoading: false,
                module: initialState.module,
                error: action.payload,
            }
        }
        default: {
            return state
        }
    }
}