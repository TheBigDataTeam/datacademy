import { combineReducers } from 'redux'
import { userAuthReducer } from './user_auth/userAuthReducer'

export const rootReducer = combineReducers({
    userAuth: userAuthReducer
})

type RootReducerType = typeof rootReducer
export type AppStateType = ReturnType<RootReducerType>