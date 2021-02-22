import React, { useCallback, useState } from 'react';
import { Link } from 'react-router-dom';
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

export const LoginPage: React.FunctionComponent = (): JSX.Element => {
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
                    <Button type='submit' fullWidth design='primary' rounded>
                        Login
                    </Button>
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

