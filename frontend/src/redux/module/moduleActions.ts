import { ModuleActionTypes } from './moduleActionTypes'
import { Module } from 'models'

/* TODO Error type */
type Error = string

export type ActionChangePayload = {
    name: string
    value: string
}

export type ActionChange = {
    type: ModuleActionTypes.FETCH_ACTION_CHANGE
    payload: ActionChangePayload
}

export type AddModuleRequest = {
    type: ModuleActionTypes.FETCH_ADD_MODULE_REQUEST
}

export type AddModuleSuccess = {
    type: ModuleActionTypes.FETCH_ADD_MODULE_SUCCESS
    payload: Module
}

export type AddModuleFailure = {
    type: ModuleActionTypes.FETCH_ADD_MODULE_FAILURE
    payload: Error
}

export type ModuleAction = ActionChange | AddModuleRequest | AddModuleSuccess | AddModuleFailure

export const fetchModuleActionChange = (name: string, value: string): ActionChange => {
    return {
        type: ModuleActionTypes.FETCH_ACTION_CHANGE,
        payload: {
            name,
            value
        }
    }
}

export const fetchAddModuleRequest = (): AddModuleRequest => {
    return {
        type: ModuleActionTypes.FETCH_ADD_MODULE_REQUEST
    }
}

export const fetchAddModuleSuccess = (module: Module): AddModuleSuccess => {
    return {
        type: ModuleActionTypes.FETCH_ADD_MODULE_SUCCESS,
        payload: module
    }
}

export const fetchAddModuleFailure = (error: Error): AddModuleFailure => {
    return {
        type: ModuleActionTypes.FETCH_ADD_MODULE_FAILURE,
        payload: error
    }
}