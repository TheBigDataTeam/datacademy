import React from 'react'
import { Grid } from 'components/ui'
import { User } from 'models/User'
import { useSelector } from 'react-redux'

export const GoalsSection: React.FunctionComponent = (): JSX.Element => {

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const user: User = useSelector((state: any) => state.userAuth.user)

    return (
        <>
            <Grid.Row>
                <Grid.Col>
                    {user && 
                        <h2>{user.name}, you have not set any goals yet</h2>
                    }
                </Grid.Col>
            </Grid.Row>
        </>
    )
}
