import React from 'react'
import { Link } from 'react-router-dom'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Heading, Grid, Paragraph } from 'components/ui'
import { useDocTitle } from 'components/hooks'

export const Page404: React.FunctionComponent = (): JSX.Element => {

    useDocTitle('404 | Page not found')

    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            <Grid.Row>
                <Grid.Col>
                    <Heading size="xxl">404</Heading>
                    <Paragraph size="l">It seems that we have not found what you were looking for</Paragraph>
                    <Paragraph>You may go to <Link to="/">home page</Link> or check our <Link to="/courses">courses</Link></Paragraph>
                </Grid.Col>
            </Grid.Row>
        </PageLayout>
    )
}
