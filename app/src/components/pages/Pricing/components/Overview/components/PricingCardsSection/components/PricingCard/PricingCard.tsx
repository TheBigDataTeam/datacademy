import React from 'react';
import { Paragraph, Button } from 'components/ui'
import styles from './PricingCard.module.css';

interface Props {
    title: string
    price: string
    features: Array<string>
}

export const PricingCard: React.FunctionComponent<Props> = ({ title, price, features }): JSX.Element => {
    return (
        <div className={styles.root}>
            <Paragraph size="xl" align="center" marginTop="l">{title}</Paragraph>
            <Paragraph align="center" size="l">{price}</Paragraph>
            <ul>
                {features.map(feature => (
                    <li key={1} className={styles.list_item}> {/* TODO random number for a key! */}
                        <Paragraph size="l">{feature}</Paragraph>
                    </li>
                ))}
            </ul>
            <Button design="primary" rounded>Get Started</Button>
        </div>
    )
}