import React from 'react'
import { Grid } from 'components/ui'

export const SubscribedSection: React.FunctionComponent = (): JSX.Element => {
    return (
        <>
            <Grid.Row>
                <Grid.Col>
                    <h2>You are not subscribed to anything</h2>
                </Grid.Col>
            </Grid.Row>
        </>
    )
}
