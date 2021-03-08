import React from 'react';
import { PageLayout } from 'components/layouts';
import { Header, Footer } from 'components/common';
import { Overview } from './components';
import styles from './CatalogPage.module.css';

export const CatalogPage: React.FunctionComponent = (): JSX.Element => {

	return (
		<PageLayout header={<Header />} footer={<Footer />} topOffset>
			<h1>Courses available</h1>
			<Overview />
		</PageLayout>
	);
};
