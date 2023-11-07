import type {Metadata} from 'next'
import {Inter} from 'next/font/google'
import '@/styles/globals.css'
import {ColorSchemeScript, MantineProvider} from '@mantine/core';

const inter = Inter({subsets: ['latin']})

export const metadata: Metadata = {
    title: 'Gas Price Oracle',
    description: 'A simple gas price oracle for Ethereum and EVM chains.',
}

export default function RootLayout({children}: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
        <head>
            <ColorSchemeScript defaultColorScheme="auto"/>
        </head>
        <body className="bg-black loading">
        <MantineProvider>{children}</MantineProvider>
        </body>
        </html>
    )
}
