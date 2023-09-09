import React, { useState } from 'react'
import { PricingCard } from './components'
import { listOfPricingCards } from './components/data/data'
import { Grid, Paragraph, Switch } from 'components/ui'
import styles from './PricingCardsSection.module.css'

export const PricingCardsSection: React.FunctionComponent = (): JSX.Element => {

    const [toggled, setToggled] = useState<boolean>(false)

    return (
        <>
            <Grid.Row>
                <Grid.Col align="center">
                    <div className={styles.wrapper}>
                        <Paragraph size="l">Annual</Paragraph>
                        <Switch rounded size="s" toggled={toggled} onToggle={() => setToggled(!toggled)} />
                        <Paragraph size="l">Monthly</Paragraph>
                    </div>
                </Grid.Col>
            </Grid.Row>
            <Grid.Row>
                {listOfPricingCards.map((card) => (
                    <Grid.Col key={card.id} cols={12} colsSM={6} colsMD={4}>
                        <PricingCard 
                            title={card.title}
                            price={toggled ? card.price : card.priceY}
                            features={card.features}
                            icon={card.icon}
                        />
                    </Grid.Col>
                ))}
            </Grid.Row>
        </>
    )
}