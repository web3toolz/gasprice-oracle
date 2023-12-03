import axios from "axios";

import {GasPriceData, GasPriceStrategy, NetworkData} from "@/types";


const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://0.0.0.0:8000/";


export async function fetchGasPrice(): Promise<NetworkData[]> {
    return (
        axios
            .get(API_URL)
            .then((response) => response.data)
            .then((data) => {
                const networkData: NetworkData[] = [];
                const networks: string[] = Object.keys(data);
                networks.map((network) => {
                    const title: string = network;
                    const updatedAt: Date = new Date(data[network].updatedAt * 1000);
                    const gasPriceData: GasPriceData[] = [
                        {title: GasPriceStrategy.SLOW, value: data[network][GasPriceStrategy.SLOW]},
                        {title: GasPriceStrategy.NORMAL, value: data[network][GasPriceStrategy.NORMAL]},
                        {title: GasPriceStrategy.FAST, value: data[network][GasPriceStrategy.FAST]},
                        // {title: GasPriceStrategy.FASTEST, value: data[network][GasPriceStrategy.FASTEST]},
                    ];
                    networkData.push({title, updatedAt: updatedAt, data: gasPriceData});
                });
                return networkData;
            })
            .catch((error) => {
                console.error("Failed to fetch data", error);
                throw error;
            })
    );
}