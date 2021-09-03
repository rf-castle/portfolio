import {NextPage} from 'next';
import MenuIcon from '@material-ui/icons/Menu';
import {IconButton} from '@material-ui/core';
import React from 'react';
import styles from '../styles/Home.module.css';

const Test: NextPage = () => {
    return (
        <div>
            <main className={styles.main}>
                <IconButton edge="start">
                    <MenuIcon/>
                </IconButton>
            </main>
        </div>

    )
}
export default Test