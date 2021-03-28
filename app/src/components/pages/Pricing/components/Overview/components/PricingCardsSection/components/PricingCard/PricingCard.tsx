import React from 'react';
import { v4 as uuidv4 } from 'uuid';
import { Link } from 'react-router-dom';
import { Price } from './components';
import { Paragraph, Button, Grid } from 'components/ui';
import styles from './PricingCard.module.css';

interface Props {
    title: string
    price: string
    features: Array<string>
}

export const PricingCard: React.FunctionComponent<Props> = ({ title, price, features }): JSX.Element => {
    return (
        <div className={styles.root}>
            <Paragraph size="xxl" align="center" marginTop="l">{title}</Paragraph>
            <Price price={price}/>
            <div className={styles.image}>Here will be an image</div>
            <Grid.Row marginTop="l">
                <Link to="/payment" className={styles.link}>
                    <Button design="primary" rounded>Get Started</Button>            
                </Link>
            </Grid.Row>
            <ul className={styles.list}>
                {features.map(feature => (
                    <li key={uuidv4()} className={styles.list_item}> {/* TODO random number for a key! */}
                        <Paragraph size="l">{feature}</Paragraph>
                    </li>
                ))}
            </ul>
        </div>
    )
}