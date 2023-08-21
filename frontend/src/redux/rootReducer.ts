import { combineReducers } from 'redux'
import { userAuthReducer, authorReducer, courseReducer, moduleReducer } from './indexReducers'

export const rootReducer = combineReducers({
    userAuth: userAuthReducer,
    author: authorReducer,
    course: courseReducer,
    module: moduleReducer,
})

type RootReducerType = typeof rootReducer

export type AppStateType = ReturnType<RootReducerType>