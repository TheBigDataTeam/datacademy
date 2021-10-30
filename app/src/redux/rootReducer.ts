import { combineReducers } from 'redux'
import { userAuthReducer, authorReducer, courseReducer } from './indexReducers'

export const rootReducer = combineReducers({
    userAuth: userAuthReducer,
    author: authorReducer,
    course: courseReducer
})

type RootReducerType = typeof rootReducer

export type AppStateType = ReturnType<RootReducerType>