import type { NextPage } from 'next'
import Head from 'next/head'
import { useAuthNextSSG } from '@m3o/auth'
import { Layout } from '@/components/Layout'

const PrivateClient: NextPage = () => {
  const { user, isAuthenticating } = useAuthNextSSG({ redirectTo: '/' })

  return (
    <div>
      <Head>
        <title>Test M3O Auth</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Layout>
        {isAuthenticating ? (
          <p>Authenticating...</p>
        ) : user ? (
          'Authenticated'
        ) : null}
      </Layout>
    </div>
  )
}

export default PrivateClient