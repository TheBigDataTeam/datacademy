import React from 'react'
import { Grid } from 'components/ui'
import { useSelector } from 'react-redux'
import { User } from 'models/User'


export const SubscribedSection: React.FunctionComponent = (): JSX.Element => {

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const user: User = useSelector((state: any) => state.userAuth.user)

    return (
        <>
            <Grid.Row>
                <Grid.Col>
                    {user && 
                        <h2>{user.name}, you are not subscribed to any course yet</h2>
                    }
                </Grid.Col>
            </Grid.Row>
        </>
    )
}
