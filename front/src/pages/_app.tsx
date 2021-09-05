import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Layout from '../templates/templates';
import { Provider } from "next-auth/client"

function MyApp({ Component, pageProps }: AppProps) {
  return (
      <Provider session={pageProps.session}>
          <Layout>
              <Component {...pageProps} />
          </Layout>
      </Provider>
  )
}
export default MyApp
