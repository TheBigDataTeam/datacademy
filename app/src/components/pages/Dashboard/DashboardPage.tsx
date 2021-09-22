import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Grid, Heading } from 'components/ui'
import { SubscribedSection, GoalsSection } from './components'
import { fetchUser } from 'redux/user/userAuthActions'

export const DashboardPage: React.FunctionComponent = (): JSX.Element => {
	const dispatch = useDispatch()
	const userID = 'x3lwbutigsfbl5qupzj5d0w3' /* hardcoded only for testing */
	useEffect(() => {
		try {
			dispatch(fetchUser(userID))
		} catch (error) {
			console.log(error)
		}
	}, [userID, dispatch])

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
