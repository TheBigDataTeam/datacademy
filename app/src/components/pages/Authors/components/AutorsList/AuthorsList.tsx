import React from 'react';
import { AuthorsCard} from './components';
import { listOfAuthors } from 'redux/store';
import { Grid } from 'components/ui';

export const AuthorsList: React.FunctionComponent = (): JSX.Element => {
    return (
        <Grid.Row>
            { listOfAuthors.map((author) => (
                <Grid.Col key={author.id} cols={12} colsSM={12} colsMD={12} marginBottom='xl'>
                    <AuthorsCard
                        firstName={author.firstName}
                        fullName={author.fullName}
                        location={author.location}
                        photo={author.photo}
                        description={author.description}
                        authorOf={author.authorOf}
                    />
                </Grid.Col>
            ))}
        </Grid.Row>
    )
}