import React from 'react'
import { Link } from 'react-router-dom'
import { Heading, Paragraph } from 'components/ui'
import { ADMIN_HOMEPAGE_URL } from 'constants/common'
import { ReactComponent as Icon } from './resources/icon.svg'
import { useSelector } from 'react-redux'
import styles from './Success.module.css'
import { AppStateType } from 'redux/rootReducer'

export const Success: React.FunctionComponent = (): JSX.Element => {

  const module = useSelector((state: AppStateType) => state.module.module)

  return (
    <div className={styles.root}>
      <Icon className={styles.icon} />
      <Heading align='center'>Module added</Heading>
      <Paragraph align='center' marginBottom='xl'>
        Module {module.title} has just been added to the site. Please check that it looks exactly as you planned
      </Paragraph>
      <Paragraph align='center'>
        <Link to={ADMIN_HOMEPAGE_URL}>To Dashboard</Link>
      </Paragraph>
  </div>
  ) 
}