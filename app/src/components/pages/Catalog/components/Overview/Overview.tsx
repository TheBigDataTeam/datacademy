import React from 'react';
import { Grid } from 'components/ui';
import { CoursesSection } from "./components";

export const Overview: React.FunctionComponent = (): JSX.Element => {
    return (
        <Grid.Row>
            <Grid.Col>
                <CoursesSection />
            </Grid.Col>
        </Grid.Row>
    )
}