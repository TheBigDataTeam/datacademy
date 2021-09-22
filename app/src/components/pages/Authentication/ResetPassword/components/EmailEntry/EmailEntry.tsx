import React, { useCallback, /* useState  */} from 'react';
import { Input, Button, Grid, Paragraph } from '../../../../../ui';

/* interface FormData {
    email: string;
}

const initialFormData: FormData = {
    email: '',
} */

interface Props {
    handlePageChangeState: () => void
}

export const EmailEntry: React.FunctionComponent<Props> = ({ handlePageChangeState }): JSX.Element => {
    //const [formData, setFormData] = useState<FormData>(initialFormData); // TODO

    const handleChange = () => {
        console.log('TODO')
    }

    const handleSubmit = useCallback(
        (event) => {
            event.preventDefault();
            handlePageChangeState();
        },
        [handlePageChangeState]
    );

    return (
        <>
            <Paragraph>
                Put your email down here and we will send you a link to reset your passowrd
            </Paragraph>
            <form onSubmit={handleSubmit}>
                <Grid.Row marginBottom="l" marginTop="l">
                    <Input
                        name='email'
                        /* value={formData.email} */
                        placeholder='Email'
                        onChange={handleChange}
                        autoFocus
                    />
                </Grid.Row>
                <Grid.Row>
                    <Button type='submit' fullWidth design='primary' rounded>
                        Reset Password
                    </Button>
                </Grid.Row>
            </form>
        </>
    );
};

