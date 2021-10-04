import React, { ChangeEventHandler, useCallback, useState } from 'react'
import { Link, useHistory } from 'react-router-dom'
import { Input, Button, Grid, Paragraph } from 'components/ui'
import { AuthLayout } from 'components/layouts'
import { SocialButtons } from "components/common"
import axios from 'axios'
import { useDispatch } from 'react-redux'
import { fetchUserLogin } from 'redux/user_auth/userAuthActions'

export const SignUpPage: React.FunctionComponent = (): JSX.Element => {

  const [email, setEmail] = useState<string>("")
  const [name, setName] = useState<string>("")
  const [surname, setSurname] = useState<string>("")
  const [password, setPassword] = useState<string>("")
  //const [disabled, setDisabled] = useState<boolean>(false)

  const history = useHistory()
  const dispatch = useDispatch()

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
      event.preventDefault();

      //setDisabled(true)
      try {
        await axios.post("http://localhost:3100/api/auth/signup", {email, name, surname, password}, {withCredentials: true})
        history.push("/dashboard")
        dispatch(fetchUserLogin())
      } catch (error) {
          console.error(error) /* TODO: handle error properly */
      }
      //setDisabled(false)
    },
    [history, dispatch, email, name, surname, password]
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
          <Button type='submit' fullWidth design='secondary' rounded /* disabled={disabled} */>Sign Up</Button>
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
