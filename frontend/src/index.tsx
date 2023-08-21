import React from 'react'
import ReactDOM from 'react-dom/client'
import { App } from './App'
import reportWebVitals from './reportWebVitals'
import { Provider } from 'react-redux'
import { store } from 'redux/store'

const container = document.getElementById('root') as HTMLElement
const rootContainer = ReactDOM.createRoot(container)

rootContainer.render(
		<Provider store={store}>
			<App />
		</Provider>
)

reportWebVitals();
