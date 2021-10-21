import React, { useCallback } from 'react'
import { Footer, Header } from 'components/common'
import { PageLayout } from 'components/layouts'
import { Grid } from 'components/ui'
import { AddAuthorForm, Success } from './components'
import { BASE_URL } from 'constants/common'
import axios from 'axios'
import { useDispatch, useSelector } from 'react-redux'
import { fetchAddAuthorFailure, fetchAddAuthorRequest, fetchAddAuthorSuccess, fetchAuthorActionChange } from 'redux/author/authorActions'
import { AppStateType } from 'redux/rootReducer'

export const AddAuthorPage: React.FunctionComponent = (): JSX.Element => {

    const dispatch = useDispatch()

    const author = useSelector((state: AppStateType) => state.author.author)
    const isLoaded = useSelector((state: AppStateType) => state.author.isLoaded)
    const isLoading = useSelector((state: AppStateType) => state.author.isLoading)

    const data =  {
      email: '',
      fullname: '',
      bio: '',
      location: '',
      facebook: '',
      instagram: '',
      twitter: '',
      shortdescription: '',
      features: '',
    }

    const handleSubmit: React.FormEventHandler<HTMLFormElement> = useCallback(
        async (event) => {
          event.preventDefault()
          dispatch(fetchAddAuthorRequest())

        try {
          await axios.post(BASE_URL + 'api/admin/add/author', {author}, {withCredentials: true})
          dispatch(fetchAddAuthorSuccess(author))
        } catch (error) {
            console.log(error) /* TODO: handle error properly */
            dispatch(fetchAddAuthorFailure(error))
        }
        },[dispatch, author]
    )

    const handleChange: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement> = useCallback((event) => {
        const { name, value } = event.target
        dispatch(fetchAuthorActionChange(name, value))
    }, [dispatch])

  return (
    <PageLayout center={isLoaded} topOffset header={<Header />} footer={<Footer />}>
      <Grid.Row>
        <Grid.Col align='center'>
          { isLoaded ? (
            <Success />
          ) : (
            <AddAuthorForm
              data={author}
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
