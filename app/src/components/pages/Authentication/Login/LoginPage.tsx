import React, { ChangeEventHandler, useCallback, useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import { SocialButtons} from 'components/common';
import { Input, Button, Grid, Paragraph } from 'components/ui';
import { AuthLayout } from 'components/layouts';
import axios from 'axios';

export const LoginPage: React.FunctionComponent = (): JSX.Element => {

    const [email, setEmail] = useState<string>("")
    const [password, setPassword] = useState<string>("")
    //const [disabled, setDisabled] = useState<boolean>(false)

    const history = useHistory() /* TODO: types */

    const handleChange = useCallback<ChangeEventHandler<HTMLInputElement>>(
        (event) => {
            if (event.target.name === "email") {
                setEmail(event.target.value)
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
            //setDisabled(true)
            try {
                await axios.post("http://localhost:3100/api/user/login", {email, password})
                history.push("/courses") /* TODO: check that it works as intended */
            } catch (error) {
                console.log(error) /* TODO: handle errors properly */
            }
            //setDisabled(false)
        },
        [history, email, password]
    )

    return (
        <AuthLayout>
            <form onSubmit={handleSubmit}>
                <Grid.Row>
                    <Input
                        name='email'
                        value={email}
                        placeholder='Email'
                        onChange={handleChange}
                        autoFocus
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
                    <Grid.Col>
                        <Paragraph align="right" size="s">
                            <Link to="/auth/forget">Forgot password?</Link>
                        </Paragraph>
                    </Grid.Col>
                </Grid.Row>
                <Grid.Row>
                    <Button type='submit' fullWidth design='primary' rounded /* disabled={disabled} */>
                        Log In
                    </Button>
                </Grid.Row>
                <Grid.Row>
                    <SocialButtons />
                </Grid.Row>
                <Grid.Row>
                    <Grid.Col>
                        <Paragraph align="center" size="s">
                            Don&apos;t have an account? <Link to="/auth/signup">Sign Up</Link>
                        </Paragraph>
                    </Grid.Col>
                </Grid.Row>
            </form>
        </AuthLayout>
    );
};

