import {useEffect, useState} from "react";
import {NetworkData} from "@/types";
import {fetchGasPrice} from "@/api/fetchGasPrice";
import {noop} from "@/utils";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://0.0.0.0:8000";


interface UseFetchGasPriceDataPollingProps {
    initialData?: NetworkData[]
    interval?: number
}

export function useFetchGasPriceDataPolling({interval, initialData}: UseFetchGasPriceDataPollingProps): {
    networkData: NetworkData[];
    loading: boolean;
    error: Error | null;
    intervalId?: NodeJS.Timeout
} {
    const [networkData, setNetworksData] = useState<NetworkData[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<Error | null>(null);
    let intervalId;

    if (initialData && initialData.length > 0) {
        setNetworksData(initialData)
    }

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data: NetworkData[] = await fetchGasPrice();
                setNetworksData(data);
            } catch (e: any) {
                setError(e);
            } finally {
                setLoading(false);
            }

        };

        fetchData().then(noop);
        intervalId = setInterval(fetchData, interval || 5000);
        return
    }, []);

    return {networkData, loading, error, intervalId};
}

export function useFetchGasPriceData(): { networkData: NetworkData[]; loading: boolean; error: Error | null } {
    const [networkData, setNetworksData] = useState<NetworkData[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<Error | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data: NetworkData[] = await fetchGasPrice();
                setNetworksData(data);
            } catch (e: any) {
                setError(e);
            } finally {
                setLoading(false);
            }

        };
        fetchData();
    }, []);


    return {networkData, loading, error};
}