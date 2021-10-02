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

	return (
		<div className={styles.root}>
			<div className={styles.left}>
				<Logo inverted={inverted} />
			</div>
			{ inverted ? <Menu inverted /> : <Menu /> }
			{ isLoaded ?
			<Link to="/auth/login" className={styles.right}>
				{ inverted ? <Logout inverted /> : <Logout /> }
			</Link>
			: <Link to="/auth/logout" className={styles.right}>
			{ inverted ? <Login inverted /> : <Login /> }
			</Link>
			}
		</div>
	)
}
