import React from 'react';
import { PageLayout } from 'components/layouts';
import { Header, Footer } from 'components/common';
import { Overview } from './components';
import { Grid, Paragraph } from 'components/ui';

export const CatalogPage: React.FunctionComponent = (): JSX.Element => {

	return (
		<PageLayout header={<Header />} footer={<Footer />} topOffset>
			<Grid.Row>
				<Grid.Col>
					<Paragraph size="xxl">Courses available</Paragraph>
				</Grid.Col>
			</Grid.Row>
			<Overview />
		</PageLayout>
	);
};
