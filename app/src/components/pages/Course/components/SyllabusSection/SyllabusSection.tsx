import React from 'react';
//import { Paragraph } from 'components/ui';

interface Props {
    syllabus: Array<string>
}

export const SyllabusSection: React.FunctionComponent<Props> = (): JSX.Element => {
    return (
        <div>
            {/* <ul>
                {syllabus.map((item => {
                    <li>
                        <Paragraph>{item}</Paragraph>
                    </li>
                }))}
            </ul>  */}         
        </div>
    )
}
