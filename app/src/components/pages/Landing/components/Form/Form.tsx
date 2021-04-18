import React from 'react';
import { Buttons } from './components';
import { Paragraph, Header } from 'components/ui';
import styles from './Form.module.css';

export const Form:React.FunctionComponent = () => {
    return (
        <div className={styles.root}>
            <Header size="xxl" weight="normal" color="inverted">Datacademy</Header>
            <Paragraph size="l" color="inverted">West Coast Analytics from Dmitriy Anoshin</Paragraph>
            <Buttons />
        </div>
    )
}
