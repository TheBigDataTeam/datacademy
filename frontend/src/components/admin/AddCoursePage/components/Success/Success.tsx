import React from 'react'
import { Link } from 'react-router-dom'
import { Heading, Paragraph } from 'components/ui'
import { ADMIN_HOMEPAGE_URL } from 'constants/urls'
import { ReactComponent as Icon } from './resources/icon.svg'
import styles from './Success.module.css'

export const Success: React.FunctionComponent = (): JSX.Element => (
  <div className={styles.root}>
    <Icon className={styles.icon} />
    <Heading align='center'>Course added</Heading>
    <Paragraph align='center' marginBottom='xl'>
      Course has just been added to the site. Please check that it looks exactly as you planned
    </Paragraph>
    <Paragraph align='center'>
      <Link to={ADMIN_HOMEPAGE_URL}>To Dashboard</Link>
    </Paragraph>
  </div>
)