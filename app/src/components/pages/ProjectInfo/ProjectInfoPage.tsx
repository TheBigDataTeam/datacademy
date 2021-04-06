import React from 'react';
import { PageLayout } from 'components/layouts';
import { Header, Footer } from 'components/common';
import { MainSection } from './components';

export const ProjectInfoPage: React.FunctionComponent = (): JSX.Element => {
    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            <MainSection />
        </PageLayout>
    )
}