import React from 'react';
import { Paragraph, Grid } from 'components/ui';
import { Icon } from 'components/common';
import styles from './MainSection.module.css';

export const MainSection: React.FunctionComponent = (): JSX.Element => {
    return (
        <>
            <Grid.Row>
                <Grid.Col cols={6} colsMD={6} colsSM={12}>
                    <Paragraph size="xxl" align="center">Our Culture and what we value</Paragraph>
                </Grid.Col>
                <Grid.Col cols={6} colsMD={6} colsSM={12}>
                    <div className={styles.cover}></div>
                </Grid.Col>
            </Grid.Row>
            <Grid.Row>
                <Paragraph size="xl" marginBottom="xl">
                    Lorem ipsum dolor sit amet consectetur, adipisicing elit. 
                    Eveniet ipsum libero dolores eius ducimus ab nam vitae maiores voluptatem harum.
                </Paragraph>
                <Paragraph>
                    <Icon type="copyright" size="m"/>
                </Paragraph>
            </Grid.Row>
            <Grid.Row>
                <Paragraph size="xl" marginBottom="xl">
                    Lorem ipsum dolor sit amet consectetur, adipisicing elit. 
                    Eveniet ipsum libero dolores eius ducimus ab nam vitae maiores voluptatem harum.
                </Paragraph>
            </Grid.Row>
            <Grid.Row>
                <Paragraph size="xl" marginBottom="xl">
                    Lorem ipsum dolor sit amet consectetur, adipisicing elit. 
                    Eveniet ipsum libero dolores eius ducimus ab nam vitae maiores voluptatem harum.
                </Paragraph>
            </Grid.Row>
            <Grid.Row>
                <Paragraph size="xl" marginBottom="xl">
                    Lorem ipsum dolor sit amet consectetur, adipisicing elit. 
                    Eveniet ipsum libero dolores eius ducimus ab nam vitae maiores voluptatem harum.
                </Paragraph>
            </Grid.Row>
        </>
    )
} 