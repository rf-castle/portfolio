import {NextPage} from 'next';
import styles from '../styles/Home.module.css';
import Head from 'next/head';

const About: NextPage = () => {
    return (
        <div className={styles.container}>
            <Head>
                <title>About Page</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className={styles.main}>
                About
            </main>

        </div>
    )
}
export default About