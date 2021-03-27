import React from 'react';
import { Paragraph } from 'components/ui';
import { Icon } from 'components/common';
import styles from './Price.module.css';

interface Props {
    price: string
}

export const Price: React.FunctionComponent<Props> = ({ price }): JSX.Element => {
    return (
        <div className={styles.root}>
            <Icon type="dollar"size="l"/>
            <Paragraph align="center" size="xl">{price}</Paragraph>
        </div>
    )
}