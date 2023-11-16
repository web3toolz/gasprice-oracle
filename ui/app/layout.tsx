import type {Metadata} from 'next'
import {Inter} from 'next/font/google'
import {MantineProvider} from '@mantine/core';
import {theme} from '@/theme';
import Script from "next/script";

import '@/styles/globals.css'
import '@mantine/core/styles.css';
import {ReactElement} from "react";

const inter = Inter({subsets: ['latin']})

const gtagId: string | undefined = process.env.NEXT_PUBLIC_GTAG_ID;

export const metadata: Metadata = {
    title: 'Gas Price Oracle | Web3toolz.com',
    description: 'Real-time gas price data for Ethereum and EVM-compatible blockchains.',
    keywords: 'EVM, Blockchain, Gas Price, Web3, Ethereum, BNB, Polygon, Optimism, Avalanche'
}

export default function RootLayout({children}: {
    children: React.ReactNode
}) {
    const gTagScript: ReactElement = (
        <>
            <Script id="gtag-source" async src={`https://www.googletagmanager.com/gtag/js?id=${gtagId}`}></Script>
            <Script id="gtag-content">
                {`
                    window.dataLayer = window.dataLayer || [];
                    function gtag(){dataLayer.push(arguments);}
                    gtag('js', new Date());
                    gtag('config', '${gtagId}');
            `}
            </Script>
        </>
    )

    return (
        <html lang="en">
        <body className="bg-black">
        <MantineProvider theme={theme}>{children}</MantineProvider>
        {gtagId && gTagScript}
        </body>
        </html>
    )
}
