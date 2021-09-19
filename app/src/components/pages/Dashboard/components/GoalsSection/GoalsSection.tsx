import React from 'react'
import { Grid } from 'components/ui'

export const GoalsSection: React.FunctionComponent = (): JSX.Element => {
    return (
        <>
            <Grid.Row>
                <Grid.Col>
                    <h2>You have not set any goals yet</h2>
                </Grid.Col>
            </Grid.Row>
        </>
    )
}
