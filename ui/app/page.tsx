"use client";
import {fetchGasPriceData, NetworkData} from "@/app/api/fetch-gasprise";
import {useEffect, useState} from "react";
import {useInterval} from '@mantine/hooks';
import {noop} from "@/utils";
import CardGrid from "@/components/CardsGrid/CardGrid";
import NetworkSelector from "@/components/NetworkSelector/NetworkSelector";
import Counter from "@/components/Counter/Counter";

const defaultChosenNetwork: string = "ethereum-mainnet";


export default function Home() {
    const [lastUpdateTime, setLastUpdateTime] = useState<Date>(new Date());
    const [networkData, setNetworkData] = useState<NetworkData>();
    const [networksData, setNetworksData] = useState<NetworkData[]>([]);
    const [chosenNetwork, setChosenNetwork] = useState<string>(defaultChosenNetwork);
    const availableNetworks: string[] = networksData.map((item: NetworkData) => item.title);

    const onSelectorChange = (value: string) => {
        setChosenNetwork(value);
        setNetworkData(networksData.filter(i => i.title === value)[0]);
    }

    const fetchGasPriceDataWrapper = async () => {
        return fetchGasPriceData()
            .then((data: NetworkData[]) => {
                console.log("update")
                setNetworksData(data);
                setNetworkData(data.filter(i => i.title === chosenNetwork)[0]);
                setLastUpdateTime(new Date());
            })
            .catch(noop);
    }

    // fetch data every 5 seconds
    const fetchDataInterval = useInterval(() => fetchGasPriceDataWrapper().then(noop), 5 * 1000);
    useEffect(() => {
        fetchGasPriceDataWrapper().then(noop);
        fetchDataInterval.start();
        return fetchDataInterval.stop;
    }, [chosenNetwork]);


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
                    networksData.length > 0 &&
                    <div className="mt-12 space-y-4 flex justify-center items-center">
                        <div className="sm:w-1/3">
                            <NetworkSelector
                                value={chosenNetwork}
                                options={availableNetworks}
                                onChange={onSelectorChange}
                            />
                        </div>
                    </div>
                }
                {
                    networksData.length > 0 &&
                    <CardGrid networkData={networkData}/>
                }
                {
                    networksData.length > 0 &&
                    <div className="mt-6 text-xs text-zinc-400 text-right">
                        <Counter updateTime={lastUpdateTime}/>
                    </div>
                }
            </div>
        </section>
    )
}

