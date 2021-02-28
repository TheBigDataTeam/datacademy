import React, { useCallback, useState } from 'react';
import { Link } from 'react-router-dom';
import styles from './Confirmation.module.css';
import { Button, Grid, Paragraph } from '../../../../../ui';
import { HOMEPAGE_URL } from '../../../../../../constants/common';

export const Confirmation: React.FunctionComponent = (): JSX.Element => {
    const email = 'test@test.com';

    return (
        <>
            <Paragraph marginBottom='xl'>
                Check email box: <b>{email}</b>. There is a link to reset your password
            </Paragraph>
            <Grid.Row>
                <Link to={HOMEPAGE_URL} className={styles.link}>
                    <Button type='button' fullWidth design='primary' rounded>
                        Ok
                    </Button>
                </Link>
            </Grid.Row>
        </>
    );
};
