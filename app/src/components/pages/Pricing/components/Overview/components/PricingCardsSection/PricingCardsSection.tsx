import React from 'react';
import { PricingCard } from './components';
import { listOfPricingCards } from 'redux/store';
import { Grid } from 'components/ui';

export const PricingCardsSection: React.FunctionComponent = (): JSX.Element => {
    return (
        <Grid.Row>
            {listOfPricingCards.map((card) => (
                <Grid.Col key={card.id} cols={12} colsSM={6} colsMD={4}>
                    <PricingCard 
                        title={card.title}
                        price={card.price}
                        features={card.features}
                    />
                </Grid.Col>
            ))}
        </Grid.Row>
    )
}