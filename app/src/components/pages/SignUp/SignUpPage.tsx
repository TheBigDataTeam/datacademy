import React, { useCallback, useState } from 'react';
import { Input, Button, Grid, Paragraph } from '../../ui';
import { AuthLayout } from '../../layouts';

interface FormData {
  email: string;
  password: string;
}

const initialFormData: FormData = {
  email: '',
  password: '',
};

export const SignUpPage: React.FunctionComponent = () => {
  const [formData, setFormData] = useState<FormData>(initialFormData);

  const handleChange = useCallback(
    ({ target }) => {
      setFormData((prevFormData) => ({
        ...prevFormData,
        [target.name]: target.value,
      }));
    },
    []
  );

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
          <Button type='submit' fullWidth design='secondary'>
            SignUp
          </Button>
        </Grid.Row>
        <Grid.Row>
          <Grid.Col cols={12}>
            <Paragraph align="center">
                Already have an account?
            </Paragraph>
          </Grid.Col>
        </Grid.Row>
        <Grid.Row>
          <Button fullWidth design='default'>
            Go to Login Page
          </Button>
        </Grid.Row>
      </form>
    </AuthLayout>
  );
};
