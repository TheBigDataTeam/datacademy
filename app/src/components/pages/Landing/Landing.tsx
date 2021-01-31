import React from 'react';
import styles from './Landing.module.css';
import { PageLayout } from '../../layouts';
import { Footer, Header } from '../../common';

export const Landing = () => {
    return (
        <PageLayout header={<Header />} footer={<Footer />}>
            <div className={styles.root}>
                <h1>Landing Page</h1>
            </div>
        </PageLayout>
    )
}
