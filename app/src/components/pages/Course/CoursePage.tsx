import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Header, Footer } from 'components/common';
import { PageLayout } from 'components/layouts';
import { Grid, Paragraph } from 'components/ui';
import { AuthorSection, SyllabusSection, BenefitsSection, TechStackSection, SubscribeSection } from './components';
import { Author, Course } from 'models';
import axios, { AxiosResponse } from 'axios';

type ParamsType = {
    id: string
}

export const CoursePage: React.FunctionComponent = (): JSX.Element => {

    const [course, setCourse] = useState<Course>(null);
    const [author, setAuthor] = useState<Author>(null);

    const params: ParamsType = useParams();

    useEffect(() => {
        const fetchData = async (): Promise<void> => {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const result: AxiosResponse<any> = await axios.get(`http://localhost:3100/courses/${params.id}`);
            setCourse(result.data);
        }
        fetchData();
    }, [params.id]);

    /* TODO URL to fetch author to be dynamic */

    useEffect(() => {
        const fetchAuthor = async ():Promise<void> => {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const result: AxiosResponse<any> = await axios.get("http://localhost:3100/authors/1");
            setAuthor(result.data);
        }
        fetchAuthor();
    }, []);

    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            {course ? 
            <>
                <Grid.Row>
                    <Grid.Col>
                        <Paragraph size="xxl">{course.title}</Paragraph>
                    </Grid.Col>
                </Grid.Row>
                <AuthorSection author={author}/>
                <BenefitsSection />
                <SyllabusSection syllabus={course.syllabus}/>
                <TechStackSection techstack={course.techstack}/>
                <SubscribeSection />
            </>
            : <h2>Loading...</h2>}
        </PageLayout>
    ) 
}