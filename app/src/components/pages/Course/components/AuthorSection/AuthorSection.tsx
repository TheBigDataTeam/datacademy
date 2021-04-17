import React from 'react';

interface Props {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: any
}

export const AuthorSection: React.FunctionComponent<Props> = ({ data }): JSX.Element => {
    return (
        <div>
            <h2>{data.data.author}</h2>
        </div>
    )
}
