import React from 'react'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Grid, Heading } from 'components/ui'
import { SubscribedSection, GoalsSection } from './components'
import { useSelector } from 'react-redux'
import { AppStateType } from 'redux/rootReducer'
import { User } from 'models'
import { useDocTitle } from 'components/hooks'
import { TITLE_PREFIX } from 'constants/common'

export const DashboardPage: React.FunctionComponent = (): JSX.Element => {

	const user: User = useSelector((state: AppStateType) => state.userAuth.user)
	
	useDocTitle(TITLE_PREFIX + user?.name + '\'s' + ' dashboard')
	
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
