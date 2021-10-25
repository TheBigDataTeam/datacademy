import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { Header, Footer } from 'components/common'
import { PageLayout } from 'components/layouts'
import { Grid, Paragraph } from 'components/ui'
import { AuthorSection, SyllabusSection, BeneficiarsSection, TechStackSection, SubscribeSection } from './components'
import { Author, Course } from 'models'
import axios, { AxiosResponse } from 'axios'

type ParamsType = {
    id: string
}

export const CoursePage: React.FunctionComponent = (): JSX.Element => {

    const [course, setCourse] = useState<Course>()

    const [authorIdToFetch, setAuthorIdToFetch] = useState<string>()

    const [author, setAuthor] = useState<Author>()

    const params: ParamsType = useParams()

    useEffect(() => {
        const fetchData = async (): Promise<void> => {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const result: AxiosResponse<any> = await axios.get(`http://localhost:3100/courses/${params.id}`)
            setCourse(result.data)
            setAuthorIdToFetch(result.data.authorid)

        }
        fetchData()
    }, [params.id])

    useEffect(() => {
        const fetchAuthor = async ():Promise<void> => {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const result: AxiosResponse<any> = await axios.get(`http://localhost:3100/authors/${authorIdToFetch}`)
            setAuthor(result.data)
        }
        fetchAuthor()
    }, [authorIdToFetch])

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