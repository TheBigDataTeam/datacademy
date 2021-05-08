import React, { useEffect, useState } from 'react';
import { Header, Footer } from 'components/common';
import { PageLayout } from 'components/layouts';
import { Grid, Paragraph } from 'components/ui';
import { AuthorSection, SyllabusSection, BenefitsSection, TechStackSection } from './components';
import axios from 'axios';

export const CoursePage: React.FunctionComponent = (): JSX.Element => {

    /* TODO TypeScript */

    const [courses, setCourses] = useState(null);
    const [author, setAuthor] = useState(null);

    /* TODO URLs to be dynamic */

    useEffect(() => {
        const fetchData = async (): Promise<void> => {
            const result = await axios.get("http://localhost:3100/courses/1")
            setCourses(result)
        }
        fetchData();
    }, []);

    useEffect(() => {
        const fetchAuthor = async ():Promise<void> => {
            const result = await axios.get("http://localhost:3100/authors/1")
            setAuthor(result)
        }
        fetchAuthor();
    }, []);

    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            {courses ? 
            <>
                <Grid.Row>
                    <Grid.Col>
                        <Paragraph size="xxl">{courses.data.title}</Paragraph>
                    </Grid.Col>
                </Grid.Row>
                <AuthorSection author={author}/>
                <SyllabusSection syllabus={courses.data.syllabus}/>
                <BenefitsSection />
                <TechStackSection techstack={courses.data.techstack}/>
            </>
            : <h2>Loading</h2>}
        </PageLayout>
    ) 
}