import React, { useState } from 'react'
import { Grid, Button } from 'components/ui'
import Modal from 'react-modal'
import styles from './SubscribeSection.module.css'

export const SubscribeSection: React.FunctionComponent = (): JSX.Element => {

    const [isModalOpen, setIsModalOpen] = useState<boolean>(false)

    const customStyles = {
        content : {
          top: '50%',
          left: '50%',
          right: 'auto',
          bottom: 'auto',
          marginRight: '-50%',
          transform: 'translate(-50%, -50%)'
        }
      }

    return (
        <Grid.Row>
            <Grid.Col>
                <Button design="primary" onClick={() => setIsModalOpen(true)}>Subscribe</Button>
                <Modal isOpen={isModalOpen} onRequestClose={() => setIsModalOpen(false)} ariaHideApp={false} style={customStyles}>
                    <div className={styles.modal_content}>
                        <h2>This function will soon be available!</h2>
                        <Button design="caution" onClick={() => setIsModalOpen(false)}>Ok</Button>
                    </div>
                </Modal>
            </Grid.Col>
        </Grid.Row>
    )
}
