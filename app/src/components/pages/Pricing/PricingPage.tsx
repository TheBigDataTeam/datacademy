import React from 'react';
import { PageLayout } from 'components/layouts';
import { Header, Footer } from 'components/common';
import { Grid, Paragraph } from 'components/ui';
import { Overview } from './components';

export const PricingPage: React.FunctionComponent = (): JSX.Element => {
    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            <Grid.Row>
                <Grid.Col>
                    <Paragraph size="xxl">Choose the right plan for you</Paragraph>
                </Grid.Col>
            </Grid.Row>
            <Overview />
        </PageLayout>
    )
}