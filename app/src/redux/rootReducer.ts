import { combineReducers } from 'redux'
import { userAuthReducer, authorReducer } from './indexReducers'

export const rootReducer = combineReducers({
    userAuth: userAuthReducer,
    author: authorReducer
})

type RootReducerType = typeof rootReducer
export type AppStateType = ReturnType<RootReducerType>