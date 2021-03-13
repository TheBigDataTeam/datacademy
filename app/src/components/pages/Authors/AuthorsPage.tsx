import React from 'react';
import { PageLayout } from 'components/layouts';
import { Header, Footer } from 'components/common';
import { Grid, Paragraph } from 'components/ui';
import { AuthorsList } from "./components";

export const AuthorsPage: React.FunctionComponent = (): JSX.Element => {

	return (
		<PageLayout header={<Header />} footer={<Footer />} topOffset>
			<Grid.Row>
				<Grid.Col>
					<Paragraph size='xxl'>Authors and founders of the project</Paragraph>
				</Grid.Col>
			</Grid.Row>
			<Grid.Row>
				<Grid.Col>
					<AuthorsList />
				</Grid.Col>
			</Grid.Row>
		</PageLayout>
	);
};
