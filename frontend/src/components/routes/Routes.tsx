import React from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import { routes } from './listOfRoutes'

export const RoutesComponent: React.FunctionComponent = (): JSX.Element => {
	
	return (
		<Router>
			<Routes>
				{routes.map(({path, element}, key) => (
					<Route path={path} element={element} key={key} />
				))}
			</Routes>
		</Router>
	)
}