import React, { useEffect, useState } from 'react'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Grid, Heading } from 'components/ui'
import { AuthorsList } from "./components"
import axios, { AxiosResponse } from 'axios'
import { BASE_URL, TITLE_PREFIX } from 'constants/common'
import { useDocTitle } from 'components/hooks'
import { Author } from 'models/Author'

export const AuthorsPage: React.FunctionComponent = (): JSX.Element => {

	const [authors, setAuthors] = useState<Author[] | null>(null);

	useEffect(() => {
		const fetchAuthors = async (): Promise<void> => {
			const results: AxiosResponse<Author[]> = await axios.get(BASE_URL+ "/api/authors")
			setAuthors(results.data)
		}
		fetchAuthors()
	}, [])

	useDocTitle(TITLE_PREFIX + 'Our Authors')

	return (
		<PageLayout header={<Header />} footer={<Footer />} topOffset>
			<Grid.Row>
				<Grid.Col>
					<Heading>Authors and founders of the project</Heading>
				</Grid.Col>
			</Grid.Row>
			<AuthorsList authors={authors}/>
		</PageLayout>
	);
};
