import React from 'react'
import { Link } from 'react-router-dom'
import { Logo } from 'components/common'
import { Menu, Login, Logout } from './components'
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
		<div className={styles.root}>
			<div className={styles.left}>
				<Logo inverted={inverted} />
			</div>
			{ inverted ? <Menu inverted /> : <Menu /> }
			<div className={styles.right}>
				{ isLoaded ? <span className={styles.name}>{user.name}</span> : <span className={styles.name}>Guest</span>}
				{ isLoaded ?
				<Link to="/auth/logout" className={styles.link}>
					{ inverted ? <Logout inverted /> : <Logout /> }
				</Link>
				: 
				<Link to="/auth/login">
					{ inverted ? <Login inverted /> : <Login /> }
				</Link>
				}
			</div>
		</div>
	)
}
