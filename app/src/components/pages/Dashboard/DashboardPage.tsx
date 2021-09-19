import React from 'react'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Grid, Heading } from 'components/ui'
import { SubscribedSection, GoalsSection } from './components'

export const DashboardPage: React.FunctionComponent = (): JSX.Element => {
    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
			<Grid.Row>
				<Grid.Col>
					<Heading>My Dashboard</Heading>
				</Grid.Col>
			</Grid.Row>
            <GoalsSection />
			<SubscribedSection />
		</PageLayout>
    )
}
