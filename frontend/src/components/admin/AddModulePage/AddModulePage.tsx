import React, { useCallback } from 'react'
import { Footer, Header } from 'components/common'
import { PageLayout } from 'components/layouts'
import { Grid } from 'components/ui'
import { AddModuleForm, Success } from './components'
import { BASE_URL } from 'constants/common'
import axios from 'axios'
import { useDispatch, useSelector } from 'react-redux'
import { fetchAddModuleFailure, fetchAddModuleRequest, fetchAddModuleSuccess, fetchModuleActionChange } from 'redux/module/moduleActions'
import { AppStateType } from 'redux/rootReducer'

export const AddModulePage: React.FunctionComponent = (): JSX.Element => {

    const dispatch = useDispatch()

    const module = useSelector((state: AppStateType) => state.module.module)
    const isLoaded = useSelector((state: AppStateType) => state.module.isLoaded)
    const isLoading = useSelector((state: AppStateType) => state.module.isLoading)

    const data =  {
        title: '',
        body: '',
        badge: '',
    }

    const handleSubmit: React.FormEventHandler<HTMLFormElement> = useCallback(
        async (event) => {
          event.preventDefault()
          dispatch(fetchAddModuleRequest())

        try {
          await axios.post(BASE_URL + '/api/admin/add/module', {module}, {withCredentials: true})
          dispatch(fetchAddModuleSuccess(module))
        } catch (error) {
            console.log(error) /* TODO: handle error properly */
            dispatch(fetchAddModuleFailure(error))
        }
        },[dispatch, module]
    )

    const handleChange: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement> = useCallback((event) => {
        const { name, value } = event.target
        dispatch(fetchModuleActionChange(name, value))
    }, [dispatch])

  return (
    <PageLayout center={isLoaded} topOffset header={<Header />} footer={<Footer />}>
      <Grid.Row>
        <Grid.Col align='center'>
          { isLoaded ? (
            <Success />
          ) : (
            <AddModuleForm
              data={module}
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
