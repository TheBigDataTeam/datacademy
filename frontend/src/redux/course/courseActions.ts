import { CourseActionTypes } from './courseActionTypes'
import { Course } from 'models'

/* TODO Error type */
type Error = string

export type ActionChangePayload = {
    name: string
    value: string
}

export type ActionChange = {
    type: CourseActionTypes.FETCH_ACTION_CHANGE
    payload: ActionChangePayload
}

export type AddCourseRequest = {
    type: CourseActionTypes.FETCH_ADD_COURSE_REQUEST
}

export type AddCourseSuccess = {
    type: CourseActionTypes.FETCH_ADD_COURSE_SUCCESS
    payload: Course
}

export type AddCourseFailure = {
    type: CourseActionTypes.FETCH_ADD_COURSE_FAILURE
    payload: Error
}

export type CourseAction = ActionChange | AddCourseRequest | AddCourseSuccess | AddCourseFailure

export const fetchCourseActionChange = (name: string, value: string): ActionChange => {
    return {
        type: CourseActionTypes.FETCH_ACTION_CHANGE,
        payload: {
            name,
            value
        }
    }
}

export const fetchAddCourseRequest = (): AddCourseRequest => {
    return {
        type: CourseActionTypes.FETCH_ADD_COURSE_REQUEST
    }
}

export const fetchAddCourseSuccess = (course: Course): AddCourseSuccess => {
    return {
        type: CourseActionTypes.FETCH_ADD_COURSE_SUCCESS,
        payload: course
    }
}

export const fetchAddCourseFailure = (error: Error): AddCourseFailure => {
    return {
        type: CourseActionTypes.FETCH_ADD_COURSE_FAILURE,
        payload: error
    }
}