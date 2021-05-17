import React from 'react';
import { AuthorsCard} from './components';
import { Grid } from 'components/ui';
import { Author } from 'models';

interface Props {
    authors: Array<Author>
}

export const AuthorsList: React.FunctionComponent<Props> = ({ authors }): JSX.Element => {

    return (
        <Grid.Row>
            { authors ? authors.map((author) => (
                <Grid.Col key={author.id} cols={12} colsSM={12} colsMD={12} marginBottom='xl'>
                    <AuthorsCard author={author} />
                </Grid.Col>
            )) : <h2>Loading...</h2> }
        </Grid.Row>
    )
}