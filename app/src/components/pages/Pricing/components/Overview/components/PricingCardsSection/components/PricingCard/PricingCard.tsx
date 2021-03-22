import React from 'react';
import styles from './PricingCard.module.css';

interface Props {
    title: string
    price: string
    features: Array<string>
}

export const PricingCard: React.FunctionComponent<Props> = ({ title, price, features }): JSX.Element => {
    return (
        <div className={styles.root}>
            <h2>{title}</h2>
            <h4>{price}</h4>
            <ul>
                {features.map(feature => (
                    <li key={1} className={styles.list_item}>{feature}</li>
                ))}
            </ul>
        </div>
    )
}