import React from 'react';
import { PageLayout } from '../../layouts';
import { Header, Footer } from '../../common'

export const Catalog: React.FunctionComponent = (): JSX.Element => {
	return (
		<PageLayout header={<Header />} footer={<Footer />} topOffset>
			<h2>All courses are here!</h2>
		</PageLayout>
	);
};
