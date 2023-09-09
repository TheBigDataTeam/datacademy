import React from 'react'
import { Link } from 'react-router-dom'
import { Logo } from 'components/common'
import { Menu, Login, Logout } from './components'
import { LOGIN_URL, LOGOUT_URL } from 'constants/urls'
import styles from './Header.module.css'
import { useSelector } from 'react-redux'
import { AppStateType } from 'redux/rootReducer'

interface Props {
	inverted?: boolean
}

export const Header: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {

	const isLoaded = useSelector((state: AppStateType) => state.userAuth.isLoaded)
	const user = useSelector((state: AppStateType) => state.userAuth.user)

	return (
		<nav className={styles.root}>
			<div className={styles.left}>
				<Logo inverted={inverted} />
			</div>
			<div className={styles.center}>
				{ inverted ? <Menu inverted /> : <Menu /> }
			</div>
			<div className={styles.right}>
				{ inverted ? <span className={styles.name_inverted}>{isLoaded ? user.name : `Guest`}</span> : 
					<span className={styles.name}>{isLoaded ? user.name : `Guest`}</span> }
				{ isLoaded ?
				<Link to={LOGOUT_URL} className={styles.link}>
					{ inverted ? <Logout inverted /> : <Logout /> }
				</Link>
				: 
				<Link to={LOGIN_URL}>
					{ inverted ? <Login inverted /> : <Login /> }
				</Link>
				}
			</div>
		</nav>
	)
}
