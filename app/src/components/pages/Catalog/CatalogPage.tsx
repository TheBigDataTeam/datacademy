import React from 'react';
import { CourseCard} from "./CourseCard";
import { PageLayout } from 'components/layouts';
import { Header, Footer } from 'components/common';
import styles from './CatalogPage.module.css';

export const CatalogPage: React.FunctionComponent = (): JSX.Element => {

	const ListOfCourseCards = [
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
			title: 'Big Data',
			image: '../logo_datalearn.png',
			description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
			tech_stack: 'SQL',
			author: 'Dima Anoshin'
		},
		{
			id: 3,
			title: 'Big Data',
			image: 'components/pages/catalog/logo_datalearn.png',
			description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
			tech_stack: 'SQL',
			author: 'Dima Anoshin'
		},
		{
			id: 4,
			title: 'Big Data',
			image: '../logo_datalearn.png',
			description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt, error!',
			tech_stack: 'SQL',
			author: 'Dima Anoshin'
		}
	]

	return (
		<PageLayout header={<Header />} footer={<Footer />} topOffset>
			<div className={styles.root}>
				{ ListOfCourseCards.map(item => (
					<CourseCard
						key={item.id}
						title={item.title}
						image={item.image}
						description={item.description}
						tech_stack={item.tech_stack}
						author={item.author}
					/>))}
			</div>
		</PageLayout>
	);
};
