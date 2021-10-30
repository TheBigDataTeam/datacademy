import { CourseActionTypes } from './courseActionTypes'
import { CourseAction } from './courseActions'
import { Course } from 'models'

export type initialStateType = {
    isLoading: boolean,
    course: Course,
    error: string,
    isLoaded: boolean
}

export const initialState: initialStateType = {
    isLoading: false,
    course: {
        title: '',
        author: '',
        description: '',
        techstack: '',
        moduleQuantity: '',
        workshopQuantity: ''
    },
    error: '',
    isLoaded: false
}

export const courseReducer = (state = initialState, action: CourseAction): initialStateType => {
    switch(action.type) {
        case CourseActionTypes.FETCH_ACTION_CHANGE: {
            const { name, value } = action.payload
            return {
                ...state,
                course: {
                  ...state.course,
                  [name]: value,
                }
            }
        }
        case CourseActionTypes.FETCH_ADD_COURSE_REQUEST: {
            return {
                ...state,
                isLoading: true
            }
        }
        case CourseActionTypes.FETCH_ADD_COURSE_SUCCESS: {
            return {
                isLoading: false,
                course: action.payload,
                error: '',
                isLoaded: true
            }
        }
        case CourseActionTypes.FETCH_ADD_COURSE_FAILURE: {
            return {
                ...state,
                isLoading: false,
                course: initialState.course,
                error: action.payload,
            }
        }
        default: {
            return state
        }
    }
}