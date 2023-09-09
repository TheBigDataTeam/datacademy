import React from 'react'
import classNames from 'classnames'
import { Container, Grid, Paragraph } from 'components/ui'
import { getYears } from './utils'
import styles from './Footer.module.css'

interface Props {
  withPadding?: boolean
}

export const Footer: React.FunctionComponent<Props> = ({ withPadding }): JSX.Element => {
	const className = classNames(styles.root, {
		[styles['root_with-padding']]: withPadding,
	})

	return (
		<div className={className}>
			<Container>
				<Grid.Row>
					<Grid.Col align='center'>
						<div>
							<Paragraph align='center'>© {getYears()} Datalearn</Paragraph>
						</div>
					</Grid.Col>
				</Grid.Row>
			</Container>
		</div>
	)
}
