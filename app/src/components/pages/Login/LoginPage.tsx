import React, { useCallback, useState } from 'react';
import { Link } from 'react-router-dom';
import { SocialButtons} from "./Components";
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
                        Log In
                    </Button>
                </Grid.Row>
                <Grid.Row>
                        <SocialButtons />
                </Grid.Row>
                <Grid.Row>
                    <Grid.Col>
                        <span>Don`&apost` have an account?</span><Link to="/signup">Sign Up</Link>
                    </Grid.Col>
                </Grid.Row>
            </form>
        </AuthLayout>
    );
};

