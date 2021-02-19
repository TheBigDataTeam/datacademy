import React from 'react';
import { Link } from 'react-router-dom';
import { Button, Text } from '../../../ui';
import styles from './Buttons.module.css';

export const Buttons:React.FunctionComponent = () => {
    return (
        <div className={styles.root}>
            <Link to="/signup" className={styles.root_link}>
                <Button design="default">
                    <Text size="l" color="inverted">Sign Up</Text>
                </Button>
            </Link>
            <Link to="/login" className={styles.root_link}>
                <Button design="default">
                    <Text size="l" color="inverted">Sign In</Text>
                </Button>
            </Link>
        </div>
    )
}
