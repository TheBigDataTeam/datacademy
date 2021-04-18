import React, { useEffect, useState } from 'react';
import { Header, Footer } from 'components/common';
import { PageLayout } from 'components/layouts';
import { Grid, Paragraph } from 'components/ui';
import { AuthorSection, SyllabusSection, BenefitsSection, TechStackSection } from './components';
import axios from 'axios';

export const CoursePage: React.FunctionComponent = (): JSX.Element => {

    /* TODO TypeScript */

    const [data, setData] = useState(null);

    useEffect(() => {
        const fetchData = async (): Promise<void> => {
            const result = await axios.get("http://localhost:3100/courses/1")
            setData(result)
        }
        fetchData();
    }, [])

    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            {data ? 
            <>
                <Grid.Row>
                    <Grid.Col>
                        <Paragraph size="xxl">{data.data.title}</Paragraph>
                    </Grid.Col>
                </Grid.Row>
                <AuthorSection data={data}/>
                <SyllabusSection syllabus={data.data.syllabus}/>
                <BenefitsSection />
                <TechStackSection techstack={data.data.techstack}/>
            </>
            : <h2>Loading</h2>}
        </PageLayout>
    ) 
}