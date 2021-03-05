import React from 'react';
import styles from './LandingPage.module.css';
import { LandingLayout } from 'components/layouts';
import { Footer, Header } from 'components/common';
import { Grid, Paragraph } from 'components/ui';
import { Form } from './components';

export const LandingPage: React.FunctionComponent = (): JSX.Element => {
	return (
		<LandingLayout header={<Header inverted/>} footer={<Footer />}>
			<div className={styles.root}>
				<Grid.Row>
					<Grid.Col>
						<Paragraph size="xxl" color="inverted">I am a Landing Page</Paragraph>
					</Grid.Col>
				</Grid.Row> 
				<Grid.Row>
					<Form />
				</Grid.Row>               
			</div>
		</LandingLayout>
	);
};
