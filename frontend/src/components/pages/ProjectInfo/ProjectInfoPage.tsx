import React from 'react'
import { PageLayout } from 'components/layouts'
import { Header, Footer } from 'components/common'
import { MainSection } from './components'
import { useDocTitle } from 'components/hooks'
import { TITLE_PREFIX } from 'constants/common'

export const ProjectInfoPage: React.FunctionComponent = (): JSX.Element => {

    useDocTitle(TITLE_PREFIX + 'Project')

    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            <MainSection />
        </PageLayout>
    )
}