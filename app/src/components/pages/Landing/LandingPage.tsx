import React from 'react'
import { Link } from 'react-router-dom'
import styles from './LandingPage.module.css'
import { LandingLayout } from 'components/layouts'
import { Footer, Header } from 'components/common'
import { Grid, Paragraph, Button } from 'components/ui'
import { Form } from './components'
import { useSelector } from 'react-redux'
import { AppStateType } from 'redux/rootReducer'

export const LandingPage: React.FunctionComponent = (): JSX.Element => {

	const isLoaded = useSelector((state: AppStateType) => state.userAuth.isLoaded)

	return (
		<LandingLayout header={<Header inverted/>} footer={<Footer />}>
			<div className={styles.root}>
				<Grid.Row>
					<Grid.Col>
						<Paragraph size="xl" color="inverted">Data Engineering / Data Science</Paragraph>
					</Grid.Col>
				</Grid.Row> 
				<Grid.Row>
					{ !isLoaded ? <Form /> : 
						<Link to="/dashboard">
							<Button design="primary">Start learning!</Button>
						</Link> 
					}
				</Grid.Row>               
			</div>
		</LandingLayout>
	)
}
