import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { Header, Footer } from 'components/common'
import { PageLayout } from 'components/layouts'
import { Grid, Paragraph } from 'components/ui'
import { AuthorSection, SyllabusSection, BeneficiarsSection, TechStackSection, SubscribeSection } from './components'
import { Author, Course } from 'models'
import { BASE_URL } from 'constants/common'
import axios, { AxiosResponse } from 'axios'

type ParamsType = {
    id?: string
}

export const CoursePage: React.FunctionComponent = (): JSX.Element => {

    const [course, setCourse] = useState<Course | null>(null)

    const [authorNameToFetch, setAuthorNameToFetch] = useState<string>()

    const [author, setAuthor] = useState<Author | null>(null)

    const params: ParamsType = useParams()

    useEffect(() => {
        const fetchData = async () => {
            const result: AxiosResponse<Course> = await axios.get(BASE_URL + `/api/courses/${params.id}`, {withCredentials: true})
            setCourse(result.data)
            setAuthorNameToFetch(result.data.author)

        }
        fetchData()
    }, [params.id])

    useEffect(() => {
        const fetchAuthor = async () => {
            const result: AxiosResponse<Author> = await axios.get(BASE_URL + `/api/authors/name/${authorNameToFetch}`, {withCredentials: true})
            setAuthor(result.data)
        }
        fetchAuthor()
    }, [authorNameToFetch])

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
                <BeneficiarsSection beneficiars={'beneficiars'}/> {/* TODO */}
                <SyllabusSection syllabus={'syllabus'}/> {/* TODO */}
                <TechStackSection techstack={course.techstack}/>
                <SubscribeSection />
            </>
            : <h2>Loading...</h2>}
        </PageLayout>
    ) 
}