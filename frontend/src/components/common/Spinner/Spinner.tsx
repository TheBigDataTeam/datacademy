import React from 'react'
import spinner from './assets/spinner.gif'

interface Props {
    witdth: number
    alt?: string
}

export const Spinner: React.FunctionComponent<Props> = ({witdth, alt}: Props): JSX.Element => {
    return (
        <img
            src={spinner}
            alt={alt ? alt : 'Loading...'}
            style={{width: `${witdth}px`, margin: 'auto', display: 'block'}} 
        />
    )
}
