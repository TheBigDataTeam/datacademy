import React, { useCallback } from 'react'
import { Footer, Header } from 'components/common'
import { PageLayout } from 'components/layouts'
import { Grid } from 'components/ui'
import { AddCourseForm, Success } from './components'
import { BASE_URL } from 'constants/common'
import axios from 'axios'
import { useDispatch, useSelector } from 'react-redux'
import { fetchAddCourseFailure, fetchAddCourseRequest, fetchAddCourseSuccess, fetchCourseActionChange } from 'redux/course/courseActions'
import { AppStateType } from 'redux/rootReducer'

export const AddCoursePage: React.FunctionComponent = (): JSX.Element => {

    const dispatch = useDispatch()

    const course = useSelector((state: AppStateType) => state.course.course)
    const isLoaded = useSelector((state: AppStateType) => state.course.isLoaded)
    const isLoading = useSelector((state: AppStateType) => state.course.isLoading)

    const data =  {
        title: '',
        author: '',
        description: '',
        techstack: '',
        moduleQuantity: '',
        workshopQuantity: ''
    }

    const handleSubmit: React.FormEventHandler<HTMLFormElement> = useCallback(
        async (event) => {
          event.preventDefault()
          dispatch(fetchAddCourseRequest())

        try {
          await axios.post(BASE_URL + '/api/admin/add/course', {course}, {withCredentials: true})
          dispatch(fetchAddCourseSuccess(course))
        } catch (error) {
            console.log(error) /* TODO: handle error properly */
            dispatch(fetchAddCourseFailure(error))
        }
        },[dispatch, course]
    )

    const handleChange: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement> = useCallback((event) => {
        const { name, value } = event.target
        dispatch(fetchCourseActionChange(name, value))
    }, [dispatch])

  return (
    <PageLayout center={isLoaded} topOffset header={<Header />} footer={<Footer />}>
      <Grid.Row>
        <Grid.Col align='center'>
          { isLoaded ? (
            <Success />
          ) : (
            <AddCourseForm
              data={course}
              errors={data}
              error={" "}
              isLoading={isLoading}
              onChange={handleChange}
              onSubmit={handleSubmit}
            />
          )}
        </Grid.Col>
      </Grid.Row>
    </PageLayout>
  )
}
