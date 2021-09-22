import { combineReducers } from 'redux'
import { userAuthReducer } from './user/userAuthReducer'

export const rootReducer = combineReducers({
    userAuth: userAuthReducer
})