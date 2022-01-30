import React, { useEffect } from 'react'
import './App.module.css'
import { useDispatch } from 'react-redux'
import { fetchUserLogin } from 'redux/user_auth/userAuthActions'
import { Routes } from 'components/routes'

export const App: React.FunctionComponent = (): JSX.Element => {

	const dispatch = useDispatch()

	useEffect(() => {
		try {
			dispatch(fetchUserLogin())
		} catch (error) {
			console.log(error) /* TODO handle errors properly */
		}
	}, [dispatch])
	
	return (
		<Routes />
	)
}
