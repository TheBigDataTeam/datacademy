import React from 'react'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Overview } from './components'
import { Grid, Heading } from 'components/ui'
import { TITLE_PREFIX } from 'constants/common'
import { useDocTitle } from 'components/hooks'

export const CatalogPage: React.FunctionComponent = (): JSX.Element => {

	useDocTitle(TITLE_PREFIX + 'Catalog')

	return (
		<PageLayout header={<Header />} footer={<Footer />} topOffset>
			<Grid.Row>
				<Grid.Col>
					<Heading>Courses available</Heading>
				</Grid.Col>
			</Grid.Row>
			<Overview />
		</PageLayout>
	)
}
