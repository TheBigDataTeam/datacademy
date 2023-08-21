import React, { ChangeEventHandler, useCallback, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { Input, Button, Grid, Paragraph } from 'components/ui'
import { AuthLayout } from 'components/layouts'
import { SocialButtons } from "components/common"
import axios from 'axios'
import { useDispatch, useSelector } from 'react-redux'
import { fetchUserLogin } from 'redux/user_auth/userAuthActions'
import { AppStateType } from 'redux/rootReducer'
import { BASE_URL, TITLE_PREFIX, USER_HOMEPAGE_URL } from 'constants/common'
import { useDocTitle } from 'components/hooks'

export const SignUpPage: React.FunctionComponent = (): JSX.Element => {

  useDocTitle(TITLE_PREFIX + 'User registration')

  const [email, setEmail] = useState<string>("")
  const [name, setName] = useState<string>("")
  const [surname, setSurname] = useState<string>("")
  const [password, setPassword] = useState<string>("")

  const navigate = useNavigate()
  const dispatch = useDispatch()
  const isLoading = useSelector((state: AppStateType) => state.userAuth.isLoading)

  const handleChange = useCallback<ChangeEventHandler<HTMLInputElement>>(
    (event) => {
        if (event.target.name === "email") {
          setEmail(event.target.value)
        }
        if (event.target.name === "name") {
          setName(event.target.value)
        }
        if (event.target.name === "surname") {
        setSurname(event.target.value)
        }
        if (event.target.name === "password") {
        setPassword(event.target.value)
        }
    },
    []
  )

  const handleSubmit = useCallback<React.FormEventHandler<HTMLFormElement>>(
    async (event) => {
      event.preventDefault()

      try {
        await axios.post(BASE_URL+ "/api/auth/signup", {email, name, surname, password}, {withCredentials: true})
        navigate(USER_HOMEPAGE_URL)
        dispatch(fetchUserLogin())
      } catch (error) {
          console.error(error) /* TODO: handle error properly */
      }
    },
    [dispatch, email, name, surname, password]
  )

  return (
    <AuthLayout>
      <form onSubmit={handleSubmit}>
        <Grid.Row>
          <Input
              name='name'
              value={name}
              placeholder='Name'
              onChange={handleChange}
              autoFocus
          />
        </Grid.Row>
        <Grid.Row>
          <Input
              name='surname'
              value={surname}
              placeholder='Surname'
              onChange={handleChange}
          />
        </Grid.Row>
        <Grid.Row>
          <Input
            name='email'
            value={email}
            placeholder='Email'
            onChange={handleChange}
          />
        </Grid.Row>
        <Grid.Row>
          <Input
            type='password'
            name='password'
            value={password}
            placeholder='Password'
            onChange={handleChange}
          />
        </Grid.Row>
        <Grid.Row>
          <Button type='submit' fullWidth design='secondary' rounded disabled={isLoading} >Sign Up</Button>
        </Grid.Row>
        <Grid.Row>
          <SocialButtons />
        </Grid.Row>
        <Grid.Row>
          <Grid.Col>
            <Paragraph align="center" size="s">
              Already have an account? <Link to="/auth/login">Log In</Link>
            </Paragraph>
          </Grid.Col>
        </Grid.Row>
      </form>
    </AuthLayout>
  )
}
