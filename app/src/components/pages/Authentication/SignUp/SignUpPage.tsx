import React, { useCallback, useState } from 'react';
import { Link } from 'react-router-dom';
import { Input, Button, Grid, Paragraph } from 'components/ui';
import { AuthLayout } from 'components/layouts';
import { SocialButtons } from "components/common";
import axios from 'axios';

/* interface FormData {
  name: string,
  surname: string,
  email: string;
  password: string;
}

const initialFormData: FormData = {
  name: '',
  surname: '',
  email: '',
  password: '',
}; */

export const SignUpPage: React.FunctionComponent = (): JSX.Element => {
  //const [formData, setFormData] = useState<FormData>(initialFormData);
  const [email, setEmail] = useState<string>()
  const [name, setName] = useState<string>()
  const [surname, setSurname] = useState<string>()
  const [password, setPassword] = useState<string>()


  const handleSubmit = useCallback(
    async (event) => {
      event.preventDefault();
      try {
        await axios.post("http://localhost:3100/api/user/signup", {email, name, surname, password})
      } catch (error) {
        console.error(error) /* TODO: handle error properly */
      }
    },
    [email, name, surname, password]
  );

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
          <Button type='submit' fullWidth design='secondary' rounded>Sign Up</Button>
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
  );
};
