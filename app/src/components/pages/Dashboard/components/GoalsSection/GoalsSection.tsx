import React from 'react'
import { Grid } from 'components/ui'
import { User } from 'models/User'
import { useSelector } from 'react-redux'
import { AppStateType } from 'redux/rootReducer'

export const GoalsSection: React.FunctionComponent = (): JSX.Element => {

    const user: User = useSelector((state: AppStateType) => state.userAuth.user)

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
