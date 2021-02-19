import React from 'react';
import { Buttons } from '../Buttons';
import { Paragraph } from '../../../ui';
import styles from './Form.module.css';

export const Form:React.FunctionComponent = () => {
    return (
        <div className={styles.root}>
            <Paragraph size="xxl" color="inverted">DataLearn</Paragraph>
            <Paragraph size="l" color="inverted">West Coast Analytics from Dmitriy Anoshin</Paragraph>
            <Buttons />
        </div>
    )
}
