import React from 'react'
import { PricingCardsSection } from './components'
import { Grid } from 'components/ui'

export const Overview: React.FunctionComponent = (): JSX.Element => {
    return (
        <Grid.Row>
            <Grid.Col>
                <PricingCardsSection />
            </Grid.Col>
        </Grid.Row>
    )
}