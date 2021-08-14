import React, { useCallback, useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import { Input, Button, Grid, Paragraph } from 'components/ui';
import { AuthLayout } from 'components/layouts';
import { SocialButtons } from "components/common";
import axios from 'axios';

export const SignUpPage: React.FunctionComponent = (): JSX.Element => {

  const [email, setEmail] = useState<string>()
  const [name, setName] = useState<string>()
  const [surname, setSurname] = useState<string>()
  const [password, setPassword] = useState<string>()
  //const [disabled, setDisabled] = useState<boolean>(false)

  const history = useHistory()

  const handleSubmit = useCallback<React.FormEventHandler<HTMLFormElement>>(
    async (event) => {
      event.preventDefault();

      /* if (!email || !name || !surname || !password) {
        setFormErrors({
          email: !email ? 'Enter your email address' : undefined,
          name: !name ? 'Enter your name' : undefined,
          surname: !surname ? 'Enter your surname' : undefined,
          password: !password ? 'Come up with a strong password' : undefined,
        })
      } */
      //setDisabled(true)
      try {
        await axios.post("http://localhost:3100/api/user/signup", {email, name, surname, password})
      } catch (error) {
        console.error(error) /* TODO: handle error properly */
      }
      //setDisabled(false)
      history.push("/courses")
    },
    [history, email, name, surname, password]
  )

  return (
    <AuthLayout>
      <form onSubmit={handleSubmit}>
        <Grid.Row>
          <Input
              name='name'
              value={name}
              placeholder='Name'
              onChange={(event) => setName(event.target.value)}
              autoFocus
          />
        </Grid.Row>
        <Grid.Row>
          <Input
              name='surname'
              value={surname}
              placeholder='Surname'
              onChange={(event) => setSurname(event.target.value)}
          />
        </Grid.Row>
        <Grid.Row>
          <Input
            name='email'
            value={email}
            placeholder='Email'
            onChange={(event) => setEmail(event.target.value)}
          />
        </Grid.Row>
        <Grid.Row>
          <Input
            type='password'
            name='password'
            value={password}
            placeholder='Password'
            onChange={(event) => setPassword(event.target.value)}
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
