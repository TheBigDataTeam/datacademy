import React from 'react'
import styles from './SocialButtons.module.css'
import { Button } from 'components/ui'
import { Icon } from 'components/common'

export const SocialButtons: React.FunctionComponent = (): JSX.Element => {
    return (
        <div className={styles.root}>
            <Button icon={<Icon type='google'/>} size='m'/>
            <Button icon={<Icon type='facebook' color/>} size='m'/>
            <Button icon={<Icon type='github'/>} size='m'/>
            <Button icon={<Icon type='linkedin'/>} size='m'/>
        </div>
    )
}