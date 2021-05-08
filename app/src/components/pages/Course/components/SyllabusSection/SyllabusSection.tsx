import React from 'react';
import { Paragraph, Header } from 'components/ui';

interface Props {
    syllabus: Array<string>
}

export const SyllabusSection: React.FunctionComponent<Props> = ({syllabus}): JSX.Element => {
    return (
        <div>
            <Header>Course content:</Header>
            <ul>
                {syllabus?.map((item) => {
                    <li>
                        <Paragraph>{item}</Paragraph>
                    </li>
                })}
            </ul>       
        </div>
    )
}
