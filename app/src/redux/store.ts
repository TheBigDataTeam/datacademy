import { createStore, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'
import { rootReducer } from './rootReducer'
import { composeWithDevTools } from 'redux-devtools-extension'

export const store = createStore(rootReducer, composeWithDevTools(applyMiddleware(thunk)))


/* TODO to be transfered to database!! */
export const listOfCourseCards = [
    {
        id: 1,
        title: 'Big Data',
        image: '../logo_datalearn.png',
        description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
        tech_stack: 'SQL',
        author: 'Dima Anoshin'
    },
    {
        id: 2,
        title: 'Small Data',
        image: '../logo_datalearn.png',
        description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
        tech_stack: 'SQL',
        author: 'Dima Anoshin'
    },
    {
        id: 3,
        title: 'Very Big Data',
        image: 'components/pages/catalog/logo_datalearn.png',
        description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
        tech_stack: 'SQL',
        author: 'Dima Anoshin'
    },
    {
        id: 4,
        title: 'So So Data',
        image: '../logo_datalearn.png',
        description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
        tech_stack: 'SQL',
        author: 'Dima Anoshin'
    }
]