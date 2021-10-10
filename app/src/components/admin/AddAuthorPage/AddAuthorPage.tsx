import React, { useCallback, useState } from 'react'
import { Footer, Header } from 'components/common'
import { PageLayout } from 'components/layouts'
import { Button, Grid } from 'components/ui'
import { AddAuthorForm, Success } from './components'
import axios from 'axios'

export const AddAuthorPage: React.FunctionComponent = (): JSX.Element => {

    const [ok, setOk] = useState(false)
    const handleClick = () => {
        setOk(ok => !ok)
    }

    const data =  {
      email: '',
      fullname: '',
      bio: '',
      location: '',
      facebook: '',
      instagram: '',
      twitter: '',
      shortdescription: '',
      features: [''],
      }

    const handleSubmit: React.FormEventHandler<HTMLFormElement> = useCallback(
        async (event) => {
        event.preventDefault();

        try {
            await axios.post('/events')
        } catch (error) {
            console.log(error) /* TODO: handle error properly */
        }
        },[]
    )

    const handleChange: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement> = useCallback((event) => {
        const { name, value } = event.target
        console.log(name, value)
    }, [])

  return (
    <PageLayout center={ok} topOffset header={<Header />} footer={<Footer />}>
      <Grid.Row>
        <Grid.Col align='center'>
          {ok ? (
            <Success />
          ) : (
            <AddAuthorForm
              data={data}
              errors={data}
              error={" "}
              isFetching={false}
              onChange={handleChange}
              onSubmit={handleSubmit}
            />
          )}
        </Grid.Col>
      </Grid.Row>
      <Button onClick={handleClick}>Submit</Button>
    </PageLayout>
  )
}
