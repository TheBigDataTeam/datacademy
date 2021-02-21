import React, { useCallback, useState } from 'react';
import { Link } from 'react-router-dom';
import { Input, Button, Grid, Paragraph } from '../../ui';
import { AuthLayout } from '../../layouts';

interface FormData {
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
};

export const SignUpPage: React.FunctionComponent = (): JSX.Element => {
  const [formData, setFormData] = useState<FormData>(initialFormData);

  const handleChange = () => {
    console.log('TODO')
  }

  const handleSubmit = useCallback(
    (event) => {
      event.preventDefault();
    },
    []
  );

  return (
    <AuthLayout>
      <form onSubmit={handleSubmit}>
        <Grid.Row>
          <Input
              name='name'
              value={formData.name}
              placeholder='Name'
              onChange={handleChange}
          />
        </Grid.Row>
        <Grid.Row>
          <Input
              name='surname'
              value={formData.surname}
              placeholder='Surname'
              onChange={handleChange}
          />
        </Grid.Row>
        <Grid.Row>
          <Input
            name='email'
            value={formData.email}
            placeholder='Email'
            onChange={handleChange}
          />
        </Grid.Row>
        <Grid.Row>
          <Input
            type='password'
            name='password'
            value={formData.password}
            placeholder='Password'
            onChange={handleChange}
          />
        </Grid.Row>
        <Grid.Row>
          <Button type='submit' fullWidth design='secondary' rounded>
            SignUp
          </Button>
        </Grid.Row>
        <Grid.Row marginBottom='none' marginTop='xl'>
          <Grid.Col cols={12}>
            <Paragraph align="center">
                Already have an account?
            </Paragraph>
          </Grid.Col>
        </Grid.Row>
        <Grid.Row>
          <Grid.Col cols={12}>
            <Paragraph align="center" marginTop='none'>
              <Link to="/login">Go to Login Page</Link>
            </Paragraph>
          </Grid.Col>
        </Grid.Row>
      </form>
    </AuthLayout>
  );
};
