import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Landing.module.css';
import { PageLayout } from '../../layouts';
import { Footer, Header } from '../../common';
import { Grid, Button } from '../../ui';

export const Landing: React.FunctionComponent = () => {
    return (
        <PageLayout header={<Header />} footer={<Footer />} topOffset>
            <div className={styles.root}>
                <Grid.Row>
                    <Grid.Col>
                        <h1>Landing Page</h1>
                    </Grid.Col>
                </Grid.Row> 
                <Grid.Row>
                    <Grid.Col>
                        <Button>
                            Signup
                        </Button>
                    </Grid.Col>
                </Grid.Row>               
            </div>
        </PageLayout>
    )
}
