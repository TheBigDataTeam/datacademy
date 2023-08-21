import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import axios, { AxiosResponse } from 'axios'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { Author } from 'models'
import { Grid } from 'components/ui'
import { SocialSection, BioSection } from './components'
import { BASE_URL, TITLE_PREFIX } from 'constants/common'
import { useDocTitle } from 'components/hooks'

interface ParamsType {
    id?: string 
}

export const AuthorPage: React.FunctionComponent = (): JSX.Element => {

    const params: ParamsType = useParams()
    const [author, setAuthor] = useState<Author | null>(null)

    useEffect(() => {
        const fetchAuthor = async () => {
            const result: AxiosResponse<Author> = await axios.get(BASE_URL + `/api/authors/${params.id}`, {withCredentials: true})
            setAuthor(result.data)
        }
        fetchAuthor()
    }, [params.id])

    useDocTitle(TITLE_PREFIX + author?.fullname)

    return (
        <PageLayout header={<Header /> } footer={<Footer />} topOffset>
            {author ?
            <>
                <Grid.Row>
                    <Grid.Col>
                        <h2>{author.fullname}</h2>
                        <h3>{author.shortdescription}</h3>
                    </Grid.Col>
                </Grid.Row>
                <BioSection bio={author.bio}/>
                <Grid.Row marginBottom="xl" marginTop="xl">
                    <Grid.Col colsLG={12} colsSM={6} align="center">
                        <SocialSection author={author} />
                    </Grid.Col>
                </Grid.Row>
            </>
            : <h2>Loading...</h2>}  
        </PageLayout>
    )
}
