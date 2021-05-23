import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import axios, { AxiosResponse } from 'axios';
import { PageLayout } from 'components/layouts';
import { Header, Footer } from 'components/common';
import { Author } from 'models';
import { Grid } from 'components/ui';
import { SocialSection, BioSection } from './components';

interface ParamsType {
    id: string
}

export const AuthorPage: React.FunctionComponent = (): JSX.Element => {

    const params: ParamsType = useParams();
    const [author, setAuthor] = useState<Author>(null);

    useEffect(() => {
        const fetchAuthor = async () => {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const result: AxiosResponse<any> = await axios.get(`http://localhost:3100/authors/${params.id}`);
            setAuthor(result.data);
        }
        fetchAuthor();
    }, [params.id]);

    return (
        <PageLayout header={<Header /> } footer={<Footer />} topOffset>
            {author ?
            <>
                <Grid.Row>
                    <Grid.Col>
                        <h2>{author.shortdescription}</h2>
                    </Grid.Col>
                </Grid.Row>
                <Grid.Row>
                    <Grid.Col>
                        <BioSection />
                    </Grid.Col>
                </Grid.Row>
                <Grid.Row>
                    <Grid.Col colsLG={12} colsSM={6}>
                        <SocialSection author={author} />
                    </Grid.Col>
                </Grid.Row>
            </>
            : <h2>Loading...</h2>}  
        </PageLayout>
    )
}
