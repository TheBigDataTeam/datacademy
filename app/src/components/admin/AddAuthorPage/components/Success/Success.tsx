import React from 'react'
import { Link } from 'react-router-dom'
import { Heading, Paragraph } from 'components/ui'
import { HOMEPAGE_URL } from 'constants/common'
import {ReactComponent as Icon} from './resources/icon.svg'
import styles from './Success.module.css'

export const Success: React.FunctionComponent = (): JSX.Element => (
  <div className={styles.root}>
    <Icon className={styles.icon} />
    <Heading align='center'>Событие добавлено</Heading>
    <Paragraph align='center' marginBottom='xl'>
      Спасибо, после прохождения модерации оно появиться на сайте.
    </Paragraph>
    <Paragraph align='center'>
      <Link to={HOMEPAGE_URL}>На главную</Link>
    </Paragraph>
  </div>
)