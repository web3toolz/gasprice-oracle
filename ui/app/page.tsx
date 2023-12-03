"use client";
import React, {useCallback, useEffect, useMemo} from "react";
import CardGrid from "@/components/CardsGrid/CardGrid";
import NetworkSelector from "@/components/NetworkSelector/NetworkSelector";
import Counter from "@/components/Counter/Counter";
import {useFetchGasPriceDataPolling} from "@/hooks/useFetchGasPriceData";
import {useFilterGasPriceData} from "@/hooks/useFilterGasPriceData";
import {clearInterval} from "timers";
import {NetworkData} from "@/types";


interface HomeProps {
    initialData: NetworkData[]
}

function HomeTitle(): React.ReactElement {
    return (
        <h1 className="text-4xl font-extrabold text-white sm:text-center sm:text-6xl">
            Cross-Chain Gas tracker
        </h1>
    )
}

function HomeDescription(): React.ReactElement {
    return (
        <p className="max-w-2xl m-auto mt-5 text-xl text-zinc-200 sm:text-center sm:text-2xl">
            Fetch real-time gas prices for multiple blockchain networks effortlessly, ensuring your transactions
            are cost-effective.
        </p>
    )
}


export default function Home() {
    const {
        networksData,
        error, loading,
        intervalId,
    } = useFetchGasPriceDataPolling({});
    const {
        networkData,
        chosenNetwork,
        setChosenNetwork,
    } = useFilterGasPriceData(networksData);

    useEffect(() => {
        return intervalId && clearInterval(intervalId);
    }, []);

    const availableNetworks = useMemo(
        () => networksData.map(item => item.title),
        [networksData]
    );

    const onSelectorChange = useCallback((value: string) => {
        setChosenNetwork(value);
    }, [setChosenNetwork]);

    const hasNetworkData = networksData.length > 0;

    return (
        <section className="bg-black">
            <div className="max-w-6xl px-4 py-8 mx-auto sm:py-24 sm:px-6 lg:px-8">
                <HomeTitle/>
                <HomeDescription/>
                {
                    hasNetworkData && (
                        <>
                            <NetworkSelector
                                value={chosenNetwork}
                                options={availableNetworks}
                                onChange={onSelectorChange}
                            />
                            <CardGrid networkData={networkData}/>
                            <div className="mt-6 text-xs text-zinc-400 text-right">
                                <Counter updateTime={networkData?.updatedAt}/>
                            </div>
                        </>
                    )
                }
                {error && <div>Error fetching data: {error.message}</div>}
            </div>
        </section>
    )
}

