import React from 'react'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Grid, Heading } from 'components/ui'
import { Overview } from './components'
import { useDocTitle } from 'components/hooks'
import { TITLE_PREFIX } from 'constants/common'

export const PricingPage: React.FunctionComponent = (): JSX.Element => {

    useDocTitle(TITLE_PREFIX + 'Pricing')

    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            <Grid.Row>
                <Grid.Col>
                    <Heading>Choose the right plan for you</Heading>
                </Grid.Col>
            </Grid.Row>
            <Overview />
        </PageLayout>
    )
}