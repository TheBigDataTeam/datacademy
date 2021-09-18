import React, { useEffect, useState } from 'react'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Grid, Paragraph } from 'components/ui'
import { AuthorsList } from "./components"
import axios, { AxiosResponse } from 'axios'

export const AuthorsPage: React.FunctionComponent = (): JSX.Element => {

	const [authors, setAuthors] = useState(null);

	useEffect(() => {
		const fetchAuthors = async (): Promise<void> => {
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
			const results: AxiosResponse<any> = await axios.get("http://localhost:3100/authors");
			setAuthors(results.data);
		}
		fetchAuthors()
	}, [])

	return (
		<PageLayout header={<Header />} footer={<Footer />} topOffset>
			<Grid.Row>
				<Grid.Col>
					<Paragraph size='xxl'>Authors and founders of the project</Paragraph>
				</Grid.Col>
			</Grid.Row>
			<Grid.Row>
				<Grid.Col>
					<AuthorsList authors={authors}/>
				</Grid.Col>
			</Grid.Row>
		</PageLayout>
	);
};
