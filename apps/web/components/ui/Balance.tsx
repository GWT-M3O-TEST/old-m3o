import type { FC } from 'react'
import Link from 'next/link'
import { useGetCurrentBalance } from '@/hooks'
import { Routes } from '@/lib/constants'
import { Spinner } from '@/components/ui'

export const Balance: FC = () => {
  const { data = 0, isLoading } = useGetCurrentBalance()

  return (
    <div className="bg-zinc-900 dark:bg-zinc-800 text-white">
      <h5 className="font-bold text-xl ">Balance</h5>
      <p className="mt-2 font-light mb-4 text-zinc-400">
        Your current account balance
      </p>
      {isLoading ? (
        <Spinner />
      ) : (
        <p className="text-3xl dark:text-white">
          {data.toFixed(6)}
        </p>
      )}
      <Link href={Routes.UserBilling}>
        <a className="block mt-4 text-xs bg-indigo-600 p-4 rounded-lg font-bold text-white align-middle text-center dark:text-white">
          Add credit
        </a>
      </Link>
    </div>
  )
}
