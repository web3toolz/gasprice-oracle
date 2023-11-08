import type {Metadata} from 'next'
import {Inter} from 'next/font/google'
import { MantineProvider, ColorSchemeScript } from '@mantine/core';
import {theme} from '@/theme';

import '@/styles/globals.css'
import '@mantine/core/styles.css';

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
        <body className="bg-black loading">
        <MantineProvider theme={theme}>{children}</MantineProvider>
        </body>
        </html>
    )
}
