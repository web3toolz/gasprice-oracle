"use client";
import {fetchGasPriceData, NetworkData} from "@/app/api/fetch-gasprise";
import {useCallback, useEffect, useMemo, useState} from "react";
import CardGrid from "@/components/CardsGrid/CardGrid";
import NetworkSelector from "@/components/NetworkSelector/NetworkSelector";
import Counter from "@/components/Counter/Counter";
import {noop} from "@/utils";

const defaultChosenNetwork: string = "ethereum-mainnet";

function useGasPriceData(defaultNetwork: string) {
    const [lastUpdateTime, setLastUpdateTime] = useState<Date>(new Date());
    const [networkData, setNetworkData] = useState<NetworkData | undefined>(undefined);
    const [networksData, setNetworksData] = useState<NetworkData[]>([]);
    const [chosenNetwork, setChosenNetwork] = useState<string>(defaultNetwork);
    const [error, setError] = useState<Error | null>(null);

    const fetchGasPriceDataWrapper = useCallback(async () => {
        try {
            const data: NetworkData[] = await fetchGasPriceData();
            setNetworksData(data);
            setNetworkData(data.find(i => i.title === chosenNetwork));
            setLastUpdateTime(new Date());
        } catch (e: any) {
            setError(e);
        }
    }, [chosenNetwork]);

    useEffect(() => {
        fetchGasPriceDataWrapper().then(noop);
        const intervalId = setInterval(fetchGasPriceDataWrapper, 5000);
        return () => clearInterval(intervalId);
    }, [fetchGasPriceDataWrapper]);
    return {lastUpdateTime, networkData, networksData, chosenNetwork, setChosenNetwork, error};
}


export default function Home() {
    const {
        lastUpdateTime,
        networkData,
        networksData,
        chosenNetwork,
        setChosenNetwork,
        error
    } = useGasPriceData(defaultChosenNetwork);

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
                <h1 className="text-4xl font-extrabold text-white sm:text-center sm:text-6xl">
                    Cross-Chain Gas tracker
                </h1>
                <p className="max-w-2xl m-auto mt-5 text-xl text-zinc-200 sm:text-center sm:text-2xl">
                    Fetch real-time gas prices for multiple blockchain networks effortlessly, ensuring your transactions
                    are cost-effective.
                </p>
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
                                <Counter updateTime={lastUpdateTime}/>
                            </div>
                        </>
                    )
                }
                {error && <div>Error fetching data: {error.message}</div>}
            </div>
        </section>
    )
}

