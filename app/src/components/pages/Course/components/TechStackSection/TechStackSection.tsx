import React from 'react';
import { Paragraph } from 'components/ui';
import { v4 as uuidv4 } from 'uuid';
import { Grid, Header } from 'components/ui';

interface Props {
    techstack: Array<string>
}

export const TechStackSection: React.FunctionComponent<Props> = ({ techstack }): JSX.Element => {
    return (
        <Grid.Row>
            <Grid.Col>
                <Header>Here is what we will be using:</Header>
                <ul>
                    {techstack.map((stack) => (
                        <li key={uuidv4()}><Paragraph size="ml">{stack}</Paragraph></li>
                    ))}
                </ul>
            </Grid.Col>
        </Grid.Row>
    )
}
