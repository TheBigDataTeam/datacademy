import React, { ChangeEventHandler, useCallback, useState } from 'react'
import { Link, useHistory } from 'react-router-dom'
import { SocialButtons} from 'components/common'
import { Input, Button, Grid, Paragraph } from 'components/ui'
import { AuthLayout } from 'components/layouts'
import { fetchUserLogin } from 'redux/user_auth/userAuthActions'
import { useDispatch, useSelector } from 'react-redux'
import axios from 'axios'
import { AppStateType } from 'redux/rootReducer'
import { BASE_URL } from 'constants/common'

export const LoginPage: React.FunctionComponent = (): JSX.Element => {

    const [email, setEmail] = useState<string>("")
    const [password, setPassword] = useState<string>("")

    const history = useHistory()
    const dispatch = useDispatch()
    const isLoaded = useSelector((state: AppStateType) => state.userAuth.isLoaded)
    const isLoading = useSelector((state: AppStateType) => state.userAuth.isLoading)

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
            try {
                await axios.post(BASE_URL + "api/auth/login", {email, password}, {withCredentials: true})
                history.push("/dashboard")
                dispatch(fetchUserLogin())
            } catch (error) {
                console.log(error) /* TODO: handle errors properly */
            }
        },
        [history, email, password, dispatch]
    )

    return (
        <AuthLayout>
            { !isLoaded ?
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
                        <Button type='submit' fullWidth design='primary' rounded disabled={isLoading} >
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
                :
                <>
                    <Grid.Row>
                        <Grid.Col>
                            <Paragraph align="center" size="l">You have already logged in</Paragraph>
                        </Grid.Col>
                    </Grid.Row>
                    <Grid.Row>
                        <Grid.Col align="center">
                            <Link to="/dashboard">
                                <Button design="primary">Start learning</Button>
                            </Link>
                        </Grid.Col>
                    </Grid.Row>
                </>
            }     
        </AuthLayout>
    )
}

