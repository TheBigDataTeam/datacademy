import React, { useEffect, useState } from 'react';
import { Header, Footer } from 'components/common';
import { PageLayout } from 'components/layouts';
import { Grid, Paragraph } from 'components/ui';
import { AuthorSection, SyllabusSection, BenefitsSection, TechStackSection } from './components';
import { Author, Course } from 'models';
import axios, { AxiosResponse } from 'axios';

export const CoursePage: React.FunctionComponent = (): JSX.Element => {

    const [course, setCourse] = useState<Course>(null);
    const [author, setAuthor] = useState<Author>(null);

    /* TODO URLs to be dynamic */

    useEffect(() => {
        const fetchData = async (): Promise<void> => {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const result: AxiosResponse<any> = await axios.get("http://localhost:3100/courses/1")
            setCourse(result.data)
        }
        fetchData();
    }, []);

    useEffect(() => {
        const fetchAuthor = async ():Promise<void> => {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const result: AxiosResponse<any> = await axios.get("http://localhost:3100/authors/1")
            setAuthor(result.data)
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
            </>
            : <h2>Loading...</h2>}
        </PageLayout>
    ) 
}