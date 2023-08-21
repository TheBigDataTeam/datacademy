import React from 'react'
import { Grid } from 'components/ui'
import { useSelector } from 'react-redux'
import { User } from 'models/User'
import { AppStateType } from 'redux/rootReducer'


export const SubscribedSection: React.FunctionComponent = (): JSX.Element => {

    const user: User = useSelector((state: AppStateType) => state.userAuth.user)

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
