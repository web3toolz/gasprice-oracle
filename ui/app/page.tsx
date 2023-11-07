"use client";
import {fetchGasPriceData, GasPriceData, NetworkData} from "@/app/api/fetch-gasprise";
import {ReactElement, useEffect, useState} from "react";
import {noop, timeDiffInSeconds} from "@/utils";
import Card from "@/components/Card/Card";


export default function Home() {
    let lastUpdateTime: Date = new Date();
    const [chosenNetwork, setChosenNetwork] = useState('');
    const [updateTimeDiff, setUpdateTimeDiff] = useState<number>(0);
    const [networkData, setNetworkData] = useState<NetworkData[]>([]);
    const ethMainnetData: NetworkData = networkData.filter((item: NetworkData) => item.title === 'ethereum-mainnet')[0] || {};

    useEffect(() => {
        const func = async () => {
            return fetchGasPriceData()
                .then((data: NetworkData[]) => {
                    setNetworkData(data);
                    lastUpdateTime = new Date();
                })
                .catch(noop);
        }

        func().then(noop);

        const apiTimeoutId = setTimeout(() => {
            func().then(noop);
        }, 5 * 1000);

        const timerTimeoutId = setInterval(() => {
            setUpdateTimeDiff(timeDiffInSeconds(lastUpdateTime));
        }, 1000)

        return () => {
            console.log("Clearing timeout");
            clearTimeout(apiTimeoutId);
            clearTimeout(timerTimeoutId);
        };
    }, [])


    let cards: ReactElement[] = [];

    if (ethMainnetData.data) {
        cards = (
            ethMainnetData.data.map((item: GasPriceData) => {
                return <Card title={item.title} gasPriceValue={item.value}/>
            })
        )
    } else {
        cards = [];
    }

    return (
        <section className="bg-black">
            <div className="max-w-6xl px-4 py-8 mx-auto sm:py-24 sm:px-6 lg:px-8">
                <h1 className="text-4xl font-extrabold text-white sm:text-center sm:text-6xl">
                    Cross-Chain Gas tracker
                </h1>
                <p className="max-w-2xl m-auto mt-5 text-xl text-zinc-200 sm:text-center sm:text-2xl">
                    Fetch real-time gas prices for multiple blockchain networks effortlessly, ensuring your transactions
                    are cost-effective.
                </p>
                <div
                    className="mt-12 space-y-4 sm:mt-16 sm:space-y-0 sm:grid sm:grid-cols-2 sm:gap-6 lg:max-w-4xl lg:mx-auto xl:max-w-none xl:mx-0 xl:grid-cols-3">
                    {cards}
                </div>
                <div className="mt-6 text-xs text-zinc-400 text-right">
                    Last update: {updateTimeDiff} seconds ago
                </div>

            </div>
        </section>
    )
}

